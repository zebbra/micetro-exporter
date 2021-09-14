package micetro

import (
	"github.com/prometheus/client_golang/prometheus"
)

type MicetroCollector struct {
	Client *Client
}

func (mc MicetroCollector) Describe(ch chan<- *prometheus.Desc) {
	prometheus.DescribeByCollect(mc, ch)
}

func (mc MicetroCollector) Collect(ch chan<- prometheus.Metric) {
	servers, err := mc.Client.DHCPServers()

	if err != nil {
		return
	}

	serverMap := make(map[string]DHCPServer)

	for _, server := range servers {
		serverMap[server.Ref] = server

		state := 1

		if server.State != "OK" {
			state = 0
		}

		ch <- prometheus.MustNewConstMetric(
			prometheus.NewDesc(
				"dhcp_server_status",
				"Status of DHCP servers",
				[]string{"ref", "host"},
				nil,
			),
			prometheus.GaugeValue,
			float64(state),
			server.Ref,
			server.Name,
		)
	}

	ranges, err := mc.Client.Ranges()

	if err != nil {
		return
	}

	rangeMap := make(map[string]Range)

	for _, r := range ranges {
		rangeMap[r.Ref] = r
	}

	scopes, err := mc.Client.DHCPScopes()

	if err != nil {
		return
	}

	for _, scope := range scopes {
		ch <- prometheus.MustNewConstMetric(
			prometheus.NewDesc(
				"dhcp_scope_available",
				"Number of addresses available in scope",
				[]string{"ref", "host", "name", "range"},
				nil,
			),
			prometheus.GaugeValue,
			float64(scope.Available),
			scope.Ref,
			serverMap[scope.DhcpServerRef].Name,
			scope.Name,
			rangeMap[scope.RangeRef].Name,
		)
	}
}
