package bd

import (
	"database/sql"
	"fmt"
	"gambituser/models"
	"gambituser/secretm"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

var SecretModel models.SecretRDSJson
var err error
var DB *sql.DB

func ReadScret() error {

	SecretModel, err = secretm.GetSecret(os.Getenv("SecretName"))

	return err
}

func DbConnect() error {
	DB, err = sql.Open("Mysql", ConStr(SecretModel))
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	err = DB.Ping()

	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	fmt.Println("Conexion existosa abd")

	return nil
}

func ConStr(clave models.SecretRDSJson) string {
	var dbUser, authToken, dbEndpoint, dbName string

	dbUser = clave.Username
	authToken = clave.Password
	dbEndpoint = clave.Host
	dbName = "gambit"
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?allowCleartextPasswords=true", dbUser, authToken, dbEndpoint, dbName)
	fmt.Println(dsn)
	return dsn
}
