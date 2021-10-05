package cmd

import (
	"fmt"
	"net/http"
	"os"

	"github.com/zebbra/micetro-exporter/internal/lib/micetro"
	"github.com/zebbra/micetro-exporter/internal/lib/version"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/spf13/cobra"
)

const User_Env = "MICETRO_USER"
const Password_Env = "MICETRO_PASSWORD"

var rootCmd = &cobra.Command{
	Use:           "micetro-exporter --api <url-to-micetro-api> ",
	SilenceErrors: true,
	Version:       fmt.Sprintf("%s-%s", version.Version, version.Commit),
	PreRunE: func(cmd *cobra.Command, args []string) error {
		if os.Getenv(User_Env) == "" || os.Getenv(Password_Env) == "" {
			return fmt.Errorf(
				"Please provide Micetro credentials in environment variables %s and %s",
				User_Env,
				Password_Env,
			)
		}

		return nil
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		addr, err := cmd.Flags().GetString("listen")

		if err != nil {
			return err
		}

		apiURL, err := cmd.Flags().GetString("api")

		if err != nil {
			return err
		}

		client := micetro.Client{
			ApiURL:   apiURL,
			Username: os.Getenv(User_Env),
			Password: os.Getenv(Password_Env),
		}

		if _, err := client.DHCPServers(); err != nil {
			return fmt.Errorf("API validation failed, returned: %s", err)
		}

		reg := prometheus.NewPedanticRegistry()

		mc := &micetro.MicetroCollector{
			Client: &client,
		}

		_ = reg.Register(mc)
		http.Handle("/metrics", promhttp.HandlerFor(reg, promhttp.HandlerOpts{}))
		return http.ListenAndServe(addr, nil)
	},
}

// Execute runs root command
func Execute() {
	rootCmd.Flags().StringP("listen", "l", ":9909", "Listen address")
	rootCmd.Flags().StringP("api", "a", "", "Micetro Central API URL")
	_ = rootCmd.MarkFlagRequired("api")

	if err := rootCmd.Execute(); err != nil {
		_, _ = fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
