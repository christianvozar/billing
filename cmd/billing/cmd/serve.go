package cmd

import (
	// Standard Library

	// Internal Dependencies
	"net/http"

	"github.com/christianvozar/billing/internal/billing"
	"github.com/christianvozar/billing/pkg/rpc"

	// External Packages
	"github.com/spf13/cobra"
)

// serveCmd represents the get command
var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "start rpc server for collecting metrics.",
	Long: `start rpc server for collecting metrics for
billing customers.`,
	Run: func(cmd *cobra.Command, args []string) {
		server := &billing.Server{} // implements Haberdasher interface
		twirpHandler := rpc.NewBillingServer(server)

		http.ListenAndServe(":8080", twirpHandler)
	},
}

func init() {
	rootCmd.AddCommand(serveCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// serveCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// serveCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
