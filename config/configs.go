package config

import "log"

// EnvDevelopment environment consts
const (
	EnvDevelopment = "development"
	EnvStaging     = "staging"
	EnvProduction  = "production"
)

// Env - app configs
var (
	Env    = EnvDevelopment
	Port   = "8080"
	DBType = "inmemory"
	DBName = "endorlabs.db"
)

// Initialize initializes all the env variables for this package.
// This function should be called only once during the application.
//
// Example:
//     LoadFromJSON("config.json", "dash.json")
//     AddEnvEntry("HOST", "localhost", &host)
//     AddEnvEntry("PORT", "8080", &port)
//     ...
//     config.Initialize()
func Initialize(files ...string) {
	if initDone {
		panic("config initialization done already")
	}
	LoadFromJSON(files...)

	addNewEnvEntry("ENV", &Env, EnvDevelopment)
	addNewEnvEntry("PORT", &Port, "8080")

	addNewEnvEntry("DB_TYPE", &DBType, DBType)
	addNewEnvEntry("DB_NAME", &DBName, DBName)

	// load all the env variables. Must be called at the end.
	load()
	log.Println("Inited config...👍")
}
