package bd

import (
	"gambituser/models"
	"gambituser/secretm"
	"os"
)

var SecretModel models.SecretRDSJson
var err error

func ReadScret() error {

	SecretModel, err = secretm.GetSecret(os.Getenv("SecretName"))

	return err
}
