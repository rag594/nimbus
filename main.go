package main

import (
	"github.com/urfave/cli/v2"
	"log"
	commands "nimbus/commands"
	"nimbus/httpClient"
	"os"
)

func main() {
	vmHost := os.Getenv("VM_HOST")
	if vmHost == "" {
		log.Fatal("Please set VM_HOST in your env variable")
		os.Exit(0)
	}
	client := httpClient.Init()
	alertCommand := commands.NewAlertCommand(vmHost, client).Command
	groupCommand := commands.NewGroupCommand(vmHost, client).Command
	cliCommands := []*cli.Command{alertCommand, groupCommand}
	app := &cli.App{
		Commands: cliCommands,
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
