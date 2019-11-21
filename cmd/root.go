package cmd

import (
	"log"
	"os"
	"profile/api"
	"profile/consts"
	"profile/core"
)

// Begin : The beginning of the App
func Begin() {
	db, err := prepareDatabase()
	if err != nil {
		log.Println(consts.App + " : " + "Database connection failed!")
		os.Exit(1)
	}
	app := api.App{
		Profile: core.GetCapsule(db),
	}
	env, err := getServerData()
	if err != nil {
		log.Println(consts.App + " : " + "Failed to setup server!")
		os.Exit(1)
	}
	app.Serve(env)
}
