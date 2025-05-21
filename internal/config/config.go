package config

// DBConfig contient les configurations pour les bases de données.
type DBConfig struct {
	MongoURI      string
	PostgresDSN   string
}

// LoadConfig charge la configuration (depuis env vars, fichiers, etc.).
// Pour l'instant, retourne des valeurs hardcodées pour l'exemple.
func LoadConfig() (*DBConfig, error) {
	// Dans une vraie application, chargez cela depuis des variables d'environnement ou un fichier.
	return &DBConfig{
		MongoURI:    "mongodb://root:rootpassword@localhost:27017/", // Correspond à docker-compose
		PostgresDSN: "postgres://spmuser:spmpassword@localhost:5432/spm_database?sslmode=disable", // Correspond à docker-compose
	}, nil
}