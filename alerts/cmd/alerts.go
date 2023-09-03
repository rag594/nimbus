package cmd

import (
	"github.com/rag594/nimbus/alerts"
	"github.com/rag594/nimbus/httpClient"
	"github.com/urfave/cli/v2"
)

type AlertCommand struct {
	Command *cli.Command
}

func NewAlertCommand(vmClient *httpClient.VMClient) *AlertCommand {
	var (
		label string
		name  string
		value string
	)
	command := &cli.Command{
		Name:  "alerts",
		Usage: "filter your alerts by name or labels",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:        "label",
				Usage:       "filter alerts by labels, provide key and then value. Example --label <key> --value <value>",
				Destination: &label,
				Category:    "Labels",
			},
			&cli.StringFlag{
				Name:        "value",
				Usage:       "filter alerts by value, value should be provided with label. Example --label <key> --value <value>",
				Destination: &value,
				Category:    "Labels",
			},
			&cli.StringFlag{
				Name:        "name",
				Usage:       "filter alerts by name",
				Destination: &name,
			},
		},
		Action: func(cCtx *cli.Context) error {
			vmAlerts, err := vmClient.GetVMAlerts()
			if err != nil {
				return err
			}

			alertsTable := alerts.NewAlertsTable(vmAlerts, "AlertName", "Labels", "Annotaions", "State", "ActiveAt")
			if table := alertsTable.ListAlertsByLabels(label, value); table != nil {
				table.Print()
			}

			if table := alertsTable.ListAlertsByName(name); table != nil {
				table.Print()
			}

			return nil
		},
	}

	return &AlertCommand{Command: command}
}
