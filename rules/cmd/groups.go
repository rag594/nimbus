package cmd

import (
	"github.com/rag594/nimbus/httpClient"
	"github.com/rag594/nimbus/rules"

	"github.com/urfave/cli/v2"
)

const (
	rulesPath = "/api/v1/rules"
)

type GroupCommand struct {
	Command *cli.Command
}

func NewGroupCommand(vmClient *httpClient.VMClient) *GroupCommand {
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
			vmRules, err := vmClient.GetVMGroups()
			if err != nil {
				return err
			}
			vmRulesTable := rules.NewRulesTable(vmRules, "Rule", "State", "Health")

			if table := vmRulesTable.ListRulesByGroupName(groupName); table != nil {
				table.Print()
			}

			if table := vmRulesTable.ListRulesByLabels(label, value); table != nil {
				table.Print()
			}

			return nil
		},
	}

	return &GroupCommand{Command: command}
}
