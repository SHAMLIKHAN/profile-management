package cmd

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"os"
	"profile/consts"

	// go package for postgres
	_ "github.com/lib/pq"
)

const (
	HOST     = "HOST"
	PORT     = "PORT"
	USER     = "USER"
	PASSWORD = "PASSWORD"
	DATABASE = "DATABASE"
	PREFIX   = "PREFIX"
	SERVER   = "SERVER"

	DBHOST       = "HOST"
	DBPORT       = "PORT"
	DBUSER       = "USER"
	DBPASSWORD   = "PASSWORD"
	DBNAME       = "DATABASE"
	SERVERPREFIX = "SERVERPREFIX"
	SERVERPORT   = "SERVERPORT"
)

func prepareDatabase() (*sql.DB, error) {
	url, err := getDatabaseURL()
	if err != nil {
		return nil, err
	}
	db, err := sql.Open("postgres", url)
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		return nil, err
	}
	log.Println(consts.App + " : " + "Database connected successfully!")
	return db, nil
}

func getDatabaseURL() (string, error) {
	env, err := getEnv()
	if err != nil {
		return consts.EmptyString, err
	}
	psql := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s database=%s sslmode=disable",
		env[DBHOST],
		env[DBPORT],
		env[DBUSER],
		env[DBPASSWORD],
		env[DBNAME],
	)
	return psql, nil
}

func getEnv() (map[string]string, error) {
	env := make(map[string]string)
	host, ok := os.LookupEnv(HOST)
	if !ok {
		return nil, errors.New("HOST environment variable required but not set")
	}
	port, ok := os.LookupEnv(PORT)
	if !ok {
		return nil, errors.New("PORT environment variable required but not set")
	}
	user, ok := os.LookupEnv(USER)
	if !ok {
		return nil, errors.New("USER environment variable required but not set")
	}
	password, ok := os.LookupEnv(PASSWORD)
	if !ok {
		return nil, errors.New("PASSWORD environment variable required but not set")
	}
	database, ok := os.LookupEnv(DATABASE)
	if !ok {
		return nil, errors.New("DATABASE environment variable required but not set")
	}
	env[DBHOST] = host
	env[DBPORT] = port
	env[DBUSER] = user
	env[DBPASSWORD] = password
	env[DBNAME] = database
	return env, nil
}

func getServerData() (map[string]string, error) {
	env := make(map[string]string)
	prefix, ok := os.LookupEnv(PREFIX)
	if !ok {
		return nil, errors.New("PREFIX environment variable required but not set")
	}
	port, ok := os.LookupEnv(SERVER)
	if !ok {
		return nil, errors.New("SERVER environment variable required but not set")
	}
	addr := ":" + port
	env[SERVERPREFIX] = prefix
	env[SERVERPORT] = addr
	return env, nil
}
