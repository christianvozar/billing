package cmd

import (
	"bufio"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/christianvozar/billing/pkg/paasio"
	"github.com/christianvozar/billing/pkg/rpc"

	"github.com/spf13/cobra"
)

// clietnCmd represents the post command
var clietnCmd = &cobra.Command{
	Use:   "client",
	Short: "start emitter for billing metrics.",
	Long: `start a long-running service to emit disk
and network IO metrics for the billing service
consume.`,
	Run: func(cmd *cobra.Command, args []string) {
		client := rpc.NewBillingProtobufClient("http://localhost:8080", &http.Client{})

		ticker := time.NewTicker(5 * time.Second)
		done := make(chan bool)

		rm := paasio.NewReaderWithMetrics()

		go func() {
			for {
				select {
				case <-done:
					return
				case _ = <-ticker.C:
					f, err := os.Open("/Users/christianvozar/Code/github.com/christianvozar/billing/data/ginsberg-howl.txt")
					openFileReader := bufio.NewReader(f)
					i, _ := openFileReader.Peek(50)
					rm.Read(i)

					dio := &rpc.IOMetricsReq{}
					dio.ReadBytes = rm.Bytes()
					dio.Reads = rm.Reads()

					postReadIO, err := client.AddIOMetrics(context.Background(), dio)
					if err != nil {
						fmt.Printf("oh no: %v", err)
						os.Exit(1)
					}
					dioJSON, err := json.Marshal(dio)
					if err != nil {
						log.Fatal(err)
					}
					fmt.Printf("%s\n", dioJSON)
					fmt.Printf("Server Response: %s\n", postReadIO.String())
				}
			}
		}()

		time.Sleep(30 * time.Second)
		ticker.Stop()
		done <- true

	},
}

func init() {
	rootCmd.AddCommand(clietnCmd)
}
