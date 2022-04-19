package models

type BaseDBResponse struct {
	Success bool   `mapstructure:"success"`
	Message string `mapstructure:"message"`
}
