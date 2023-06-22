# nimbus
CLI based tool to query VMAlerts.VMAlerts Ui is currently cluttered and too slow.Want fast access to your rules/alerts then just nimbus it.

## Install


## Usage

### Alerts

```console
foo@bar:~$ nimbus alerts help
NAME:
   CLI for VM Alerts alerts - get the list of alerts and other info(state, description etc)

USAGE:
   use team or alertGroup flags to get info on team/alertGroup wise alerts

OPTIONS:
   --team value        team wise alerts in firing or pending state
   --alertGroup value  group wise alerts in firing or pending state
   --help, -h          show help
```

### Groups

```console
foo@bar:~$ nimbus group help
NAME:
   CLI for VM Alerts group - lists down the rules

USAGE:
   use name flag to get info alertRules applied and get their state

OPTIONS:
   --name value  list rules name wise
   --help, -h    show help
```
