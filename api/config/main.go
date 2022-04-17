package config

func InitializeApplication() {
	readConfigFile()

	initializeMainLogger()
	initializeDBCerts()
}

func CloseApplication() {
	closeLogger()
}
