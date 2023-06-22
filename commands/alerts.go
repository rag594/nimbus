package commands

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/fatih/color"
	"github.com/rodaine/table"
	"github.com/urfave/cli/v2"
	"net/http"
	"nimbus/httpClient"
	"nimbus/models"
)

const (
	alertPath = "/vmalert/api/v1/alerts"
)

type AlertCommand struct {
	Command *cli.Command
}

func NewAlertCommand(host string, client *httpClient.Client) *AlertCommand {
	uri := fmt.Sprintf("%s%s", host, alertPath)
	var team string
	var alertGroup string
	command := &cli.Command{
		Name:  "alerts",
		Usage: "enter your team name to get info on alerts",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:        "team",
				Usage:       "team wise alerts in firing or pending state",
				Destination: &team,
			},
			&cli.StringFlag{
				Name:        "alertGroup",
				Usage:       "group wise alerts in firing or pending state",
				Destination: &alertGroup,
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
				if vmAlert.Labels.Team == team || vmAlert.Labels.Alertgroup == alertGroup {
					tbl.AddRow(vmAlert.Labels.Alertname, vmAlert.Annotations.Summary, vmAlert.Annotations.Description, vmAlert.State)
				}

			}

			tbl.Print()

			return nil
		},
	}

	return &AlertCommand{Command: command}
}
