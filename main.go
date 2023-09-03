package main

import (
	"log"
	"os"

	alertsCmd "github.com/rag594/nimbus/alerts/cmd"
	"github.com/rag594/nimbus/httpClient"
	rulesCmd "github.com/rag594/nimbus/rules/cmd"

	"github.com/urfave/cli/v2"
)

func main() {
	vmHost := os.Getenv("VM_HOST")
	if vmHost == "" {
		log.Fatal("Please set VM_HOST in your env variable")
		os.Exit(0)
	}
	client := httpClient.Init()
	vmClient := httpClient.NewVmClient(vmHost, client)
	alertCommand := alertsCmd.NewAlertCommand(vmClient).Command
	groupCommand := rulesCmd.NewGroupCommand(vmClient).Command
	cliCommands := []*cli.Command{alertCommand, groupCommand}
	app := &cli.App{
		Name:     "CLI for VM Alerts",
		Commands: cliCommands,
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
