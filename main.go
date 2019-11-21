package main

import (
	"log"
	"profile/cmd"
	"profile/consts"
)

func main() {
	log.Println("[Service]" + " " + consts.App + " : " + "Begin")
	cmd.Begin()
}
