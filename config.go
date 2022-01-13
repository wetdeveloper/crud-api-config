package config
func AppConfig() map[string]string {
	config := make(map[string]string)
	config["database-address"] = "users.db"
	config["database-name"] = "users"
	config["database-username"] = "None"
	config["database-password"] = "None"
	return config
}
