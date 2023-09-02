package cmd

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/rag594/nimbus/httpClient"
	"github.com/rag594/nimbus/rules"

	"github.com/fatih/color"
	"github.com/rodaine/table"
	"github.com/urfave/cli/v2"
)

const (
	rulesPath = "/api/v1/rules"
)

type GroupCommand struct {
	Command *cli.Command
}

func NewGroupCommand(host string, client *httpClient.Client) *GroupCommand {
	uri := fmt.Sprintf("%s%s", host, rulesPath)
	var (
		label     string
		groupName string
		value     string
	)
	command := &cli.Command{
		Name:      "group",
		Usage:     "lists down the rules",
		UsageText: "use name flag to get info alertRules applied and get their state",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:        "label",
				Usage:       "filter rules by labels, provide key and then value. Example --label <key> --value <value>",
				Destination: &label,
				Category:    "Labels",
			},
			&cli.StringFlag{
				Name:        "value",
				Usage:       "filter rules by value, value should be provided with label. Example --label <key> --value <value>",
				Destination: &value,
				Category:    "Labels",
			},
			&cli.StringFlag{
				Name:        "name",
				Usage:       "list rules name wise",
				Destination: &groupName,
			},
		},
		Action: func(cCtx *cli.Context) error {
			res, err := client.Get(context.Background(), uri, nil)
			vmGroups := &rules.VMGroups{}
			if err != nil {
				fmt.Println("Error in requesting groups data", err)
			}
			defer res.Body.Close()

			if res.StatusCode == http.StatusOK {

				if err := json.NewDecoder(res.Body).Decode(vmGroups); err != nil {
					fmt.Println("Err in decoding ", err)
				}
			}

			headerFmt := color.New(color.FgGreen, color.Underline).SprintfFunc()
			columnFmt := color.New(color.FgYellow).SprintfFunc()

			tbl := table.New("Rule", "State")
			tbl.WithHeaderFormatter(headerFmt).WithFirstColumnFormatter(columnFmt)
			for _, vmGroup := range vmGroups.Data.Groups {
				if groupName != "" && vmGroup.Name == groupName {
					for _, rule := range vmGroup.Rules {
						tbl.AddRow(rule.Name, rule.State)
					}
				}

				if label != "" && value != "" {
					for _, rule := range vmGroup.Rules {
						val, ok := rule.Labels[label]
						if ok && val == value {
							tbl.AddRow(rule.Name, rule.State)
						}
					}
				}

			}

			tbl.Print()

			return nil
		},
	}

	return &GroupCommand{Command: command}
}
