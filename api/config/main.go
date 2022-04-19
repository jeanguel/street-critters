package config

// InitializeApplication
func InitializeApplication() {
	readConfigFile()

	initializeMainLogger()
	initializeDBCerts()
}

// CloseApplication
func CloseApplication() {
	closeLogger()
}
