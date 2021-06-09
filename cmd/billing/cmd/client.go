package cmd

import (
	"context"
	"fmt"
	"net/http"
	"os"

	"github.com/christianvozar/billing/pkg/rpc"

	"github.com/shirou/gopsutil/v3/disk"
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

		d, _ := disk.IOCounters()
		/*
			type DiskIOMetricsReq struct {
				state         protoimpl.MessageState
				sizeCache     protoimpl.SizeCache
				unknownFields protoimpl.UnknownFields

				Customer   string `protobuf:"bytes,1,opt,name=customer,proto3" json:"customer,omitempty"`                        // UUID for Customer to charge (bill)
				Name       string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`                                // Device name
				Serial     string `protobuf:"bytes,3,opt,name=serial,proto3" json:"serial,omitempty"`                            // Device serial number
				Read       int64  `protobuf:"varint,4,opt,name=read,proto3" json:"read,omitempty"`                               // Counter
				Writes     int64  `protobuf:"varint,5,opt,name=writes,proto3" json:"writes,omitempty"`                           // Counter
				ReadBytes  int64  `protobuf:"varint,6,opt,name=read_bytes,json=readBytes,proto3" json:"read_bytes,omitempty"`    // Counter in Bytes
				WriteBytes int64  `protobuf:"varint,7,opt,name=write_bytes,json=writeBytes,proto3" json:"write_bytes,omitempty"` // Counter in Bytes
			}
		*/

		/*
		   ReadCount        uint64 `json:"readCount"`
		   MergedReadCount  uint64 `json:"mergedReadCount"`
		   WriteCount       uint64 `json:"writeCount"`
		   MergedWriteCount uint64 `json:"mergedWriteCount"`
		   ReadBytes        uint64 `json:"readBytes"`
		   WriteBytes       uint64 `json:"writeBytes"`
		   ReadTime         uint64 `json:"readTime"`
		   WriteTime        uint64 `json:"writeTime"`
		   IopsInProgress   uint64 `json:"iopsInProgress"`
		   IoTime           uint64 `json:"ioTime"`
		   WeightedIO       uint64 `json:"weightedIO"`
		   Name             string `json:"name"`
		   SerialNumber     string `json:"serialNumber"`
		   Label            string `json:"label"`
		*/

dio := &rpc.DiskIOMetricsReq{}
		for _, counters := range d {
			dio.Customer =    "CirrusMD",/*
				//dio.Name =      counters.Name
				dio.ReadBytes = int64(counters.ReadBytes)
				dio.WriteBytes = int64(counters.WriteBytes)
				dio.Reads =     int64(counters.ReadCount)
				dio.Writes =    int64(counters.WriteCount)
*/
			if counters.SerialNumber != "" {
				dio.Serial = counters.SerialNumber
			}}
		}

		subDiskIO, err := client.AddDiskIOMetrics(context.Background(), dio)
		if err != nil {
			fmt.Printf("oh no: %v", err)
			os.Exit(1)
		}
		fmt.Printf("%s", dio.String())
	},
}

func init() {
	rootCmd.AddCommand(clietnCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// clietnCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// clietnCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
