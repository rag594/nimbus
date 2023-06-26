package commands

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/fatih/color"
	"github.com/rag594/nimbus/httpClient"
	"github.com/rag594/nimbus/models"
	"github.com/rodaine/table"
	"github.com/urfave/cli/v2"
)

const (
	alertPath = "/api/v1/alerts"
)

type AlertCommand struct {
	Command *cli.Command
}

func NewAlertCommand(host string, client *httpClient.Client) *AlertCommand {
	uri := fmt.Sprintf("%s%s", host, alertPath)
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
			res, err := client.Get(context.Background(), uri, nil)
			vmAlerts := &models.VMAlerts{}
			if err != nil {
				fmt.Println("Error in requesting vmalerts data", err)
			}
			defer res.Body.Close()

			if res.StatusCode == http.StatusOK {

				if err := json.NewDecoder(res.Body).Decode(vmAlerts); err != nil {
					fmt.Println("Err in decoding ", err)
				}
			}

			headerFmt := color.New(color.FgGreen, color.Underline).SprintfFunc()
			columnFmt := color.New(color.FgYellow).SprintfFunc()

			tbl := table.New("AlertName", "Summary", "Description", "State")
			tbl.WithHeaderFormatter(headerFmt).WithFirstColumnFormatter(columnFmt)
			for _, vmAlert := range vmAlerts.Data.Alerts {
				if name != "" && vmAlert.Name == name {
					tbl.AddRow(vmAlert.Labels["alertname"], vmAlert.Annotations.Summary, vmAlert.Annotations.Description, vmAlert.State)
				}

				if label != "" && value != "" {
					val, ok := vmAlert.Labels[label]
					if ok && val == value {
						tbl.AddRow(vmAlert.Labels["alertname"], vmAlert.Annotations.Summary, vmAlert.Annotations.Description, vmAlert.State)
					}
				}

			}

			tbl.Print()

			return nil
		},
	}

	return &AlertCommand{Command: command}
}
