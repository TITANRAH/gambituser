package secretm

import (
	"encoding/json"
	"fmt"
	"gambituser/awsgo"
	"gambituser/models"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/secretsmanager"
)

func GetSecret(nombreSecret string) (models.SecretRDSJson, error) {
	var datosSecret models.SecretRDSJson
	// este print queda grabado en cloudwatch
	fmt.Println("-> Pido secreto" + nombreSecret)

	// iniciar
	svc := secretsmanager.NewFromConfig(awsgo.Cfg)
	// clave tienee los valores
	clave, err := svc.GetSecretValue(awsgo.Ctx, &secretsmanager.GetSecretValueInput{
		SecretId: aws.String(nombreSecret),
	})

	if err != nil {
		fmt.Println(err.Error())
		return datosSecret, err
	}

	// json unmarchal pide lo que va a recibir y donde lo tiene que grabar

	json.Unmarshal([]byte(*clave.SecretString), &datosSecret)

	fmt.Println(" Lectura Secret ok")

	return datosSecret, nil

}
