package httpClient

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/rag594/nimbus/alerts"
	"github.com/rag594/nimbus/rules"
)

const (
	alertsURI = "/api/v1/alerts"
	rulesURI  = "/api/v1/rules"
)

type VMClient struct {
	client *Client
	host   string
}

func NewVmClient(host string, client *Client) *VMClient {
	return &VMClient{client: client, host: host}
}

func (v *VMClient) GetVMAlerts() (*alerts.VMAlertsResponse, error) {
	res, err := v.client.Get(context.Background(), alertsURI, nil)
	if err != nil {
		fmt.Println("Error in requesting vmalerts data", err)
	}
	defer res.Body.Close()

	if res.StatusCode == http.StatusOK {
		vmAlerts := &alerts.VMAlertsResponse{}
		if err := json.NewDecoder(res.Body).Decode(vmAlerts); err != nil {
			fmt.Println("Err in decoding ", err)
			return nil, err
		}
		return vmAlerts, nil
	}

	return nil, errors.New("failure in getting alerts repsonse from vm")
}

func (v *VMClient) GetVMGroups() (*rules.VMGroupResponse, error) {
	res, err := v.client.Get(context.Background(), rulesURI, nil)
	if err != nil {
		fmt.Println("Error in requesting vmalerts data", err)
	}
	defer res.Body.Close()

	if res.StatusCode == http.StatusOK {
		vmRules := &rules.VMGroupResponse{}
		if err := json.NewDecoder(res.Body).Decode(vmRules); err != nil {
			fmt.Println("Err in decoding ", err)
			return nil, err
		}
		return vmRules, nil
	}

	return nil, errors.New("failure in getting rules repsonse from vm")
}
