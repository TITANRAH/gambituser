package models

type SecretRDSJson struct {
	// donde venga este campo lo grabara como username
	Username string `json:"username"`
	Password string `json:"password"`
	Engine string `json:"engine"`
	Host string `json:"host"`
	Port int `json:"port"`
	DbClusterIdentifier string `json:"dbClusteridentifier"`
}

type SignUp struct {
	UserEmail string `json:"UserEmail"`
	UserUUID string `json:"UserUUID"`

}