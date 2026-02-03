package model

type Config struct{
	PORT string
	DB_HOST string
	DB_PORT string
	DB_USER string
	DB_PASSWORD string
	DB_NAME string

	JWT_SECRET string
	N8N_BASE_URL string
}