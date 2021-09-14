package micetro

import (
	"fmt"

	"github.com/go-resty/resty/v2"
)

type Client struct {
	ApiURL   string
	Username string
	Password string
}

func (c *Client) Request() *resty.Request {
	rc := resty.New()
	rc.DisableWarn = true

	return rc.R().SetBasicAuth(c.Username, c.Password)
}

func (c *Client) DHCPScopes() ([]DHCPScope, error) {
	resp, err := c.Request().
		SetResult(&DHCPScopeList{}).
		Get(c.ApiURL + "/DHCPScopes")

	if err != nil {
		return nil, err
	}

	if resp.IsError() {
		return nil, fmt.Errorf("%s: %s", resp.Status(), resp.String())
	}

	list := resp.Result().(*DHCPScopeList)
	return list.Result.DHCPScopes, nil
}

func (c *Client) DHCPServers() ([]DHCPServer, error) {
	resp, err := c.Request().
		SetResult(&DHCPServerList{}).
		Get(c.ApiURL + "/DHCPServers")

	if err != nil {
		return nil, err
	}

	if resp.IsError() {
		return nil, fmt.Errorf("%s: %s", resp.Status(), resp.String())
	}

	list := resp.Result().(*DHCPServerList)
	return list.Result.DHCPServers, nil
}

func (c *Client) Ranges() ([]Range, error) {
	resp, err := c.Request().
		SetResult(&RangeList{}).
		Get(c.ApiURL + "/Ranges")

	if err != nil {
		return nil, err
	}

	if resp.IsError() {
		return nil, fmt.Errorf("%s: %s", resp.Status(), resp.String())
	}

	list := resp.Result().(*RangeList)
	return list.Result.Ranges, nil
}
