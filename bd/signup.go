package bd

import (
	"fmt"
	"gambituser/models"
	"gambituser/tools"

	_ "github.com/go-sql-driver/mysql"
)

func SignUp(sig models.SignUp) error {
	fmt.Println("Comienza Registro")

	err := DbConnect()
	if err != nil {
		return err
	}
	defer DB.Close()

	sentencia := "INSER INTO users (User_Email, User_UUID, User_DateAdd) VALUES ('" + sig.UserEmail + "','" + sig.UserUUID + "', '" + tools.FechaMySQL() + "')"
	fmt.Println(sentencia)

	_, err = DB.Exec(sentencia)

	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	fmt.Println("Sign up > ejecución exítosa")

	return nil
}
