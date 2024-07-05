package main

import (
	"context"
	"errors"
	"fmt"
	"gambituser/awsgo"
	"gambituser/bd"
	"gambituser/models"
	"os"

	"github.com/aws/aws-lambda-go/events"
	lambda "github.com/aws/aws-lambda-go/lambda"
)

func main() {

	lambda.Start(EjecutoLambda)
}

func EjecutoLambda(ctx context.Context, event events.CognitoEventUserPoolsPostConfirmation) (events.CognitoEventUserPoolsPostConfirmation, error) {
	awsgo.InicializoAWS()

	// si no viene el secrete name que es una variable de entorno en aws
	// voy a devolver un error  y el evento qye recibo

	if !ValidoParametros() {
		fmt.Println("Error en los parametros debe enviar secret name")

		err := errors.New("Error en los parametros debe enviar Secret Name ")

		return event, err
	}

	var datos models.SignUp

	for row, att := range event.Request.UserAttributes {
		switch row {
		case "email":
			datos.UserEmail = att
			fmt.Println("Email = " + datos.UserEmail)

		case "sub":
			datos.UserUUID = att
			fmt.Println("Sub = " + datos.UserEmail)

		}

	}
	err := bd.ReadScret()

	if err != nil {
		fmt.Println("Error al leer el secret")
		return event, err
	}



}

func ValidoParametros() bool {
	var traeParametro bool
	_, traeParametro = os.LookupEnv("SecretName")

	return traeParametro
}
