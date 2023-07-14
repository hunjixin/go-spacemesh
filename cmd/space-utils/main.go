// go-spacemesh is a golang implementation of the Spacemesh node.
// See - https://spacemesh.io
package main

import (
	"context"
	"encoding/json"
	"fmt"
	_ "net/http/pprof"
	"os"

	"github.com/golang/protobuf/ptypes/empty"
	"github.com/mitchellh/mapstructure"
	pb "github.com/spacemeshos/api/release/go/spacemesh/v1"
	"github.com/spacemeshos/go-spacemesh/config"
	"github.com/spacemeshos/go-spacemesh/config/presets"
	"github.com/spacemeshos/go-spacemesh/node/mapstructureutil"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	path string
)

func main() { // run the app
	mycmd := &cobra.Command{
		Use:   "query",
		Short: "start node",
		Run: func(c *cobra.Command, args []string) {
			ctx := context.Background()
			if err := run(ctx, path); err != nil {
				fmt.Println(err.Error())
			}
		},
	}

	mycmd.Flags().StringVar(&path, "path", "", "url to connect")

	if err := mycmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func run(ctx context.Context, path string) error {
	urlMap := make(map[string]string)

	cfgData, err := os.ReadFile(path)
	if err != nil {
		return err
	}

	err = json.Unmarshal(cfgData, &urlMap)
	if err != nil {
		return err
	}

	for name, url := range urlMap {
		conn, err := grpc.Dial(url, grpc.WithTransportCredentials(insecure.NewCredentials()))
		if err != nil {
			return err
		}

		client := pb.NewSmesherServiceClient(conn)
		status, err := client.PostSetupStatus(ctx, &empty.Empty{})
		if err != nil {
			return err
		}

		postCfg, err := client.PostConfig(ctx, &empty.Empty{})
		if err != nil {
			return err
		}

		/*
		 const progress =
		      ((numLabelsWritten * smesherConfig.bitsPerLabel) /
		        (BITS * commitmentSize)) *
		      100;


		  {formatBytes(
		                  (numLabelsWritten * smesherConfig.bitsPerLabel) / BITS
		                )}{' '}
		                / {formatBytes(commitmentSize)}, {progress.toFixed(2)}%

		 const commitmentSize = state.config
		        ? (state.config.labelsPerUnit * state.config.bitsPerLabel * numUnits) /
		          BITS
		        : 0;
		*/

		commitmentSize := postCfg.LabelsPerUnit * uint64(postCfg.BitsPerLabel) * (1000000) / 8
		completed := status.GetStatus().NumLabelsWritten * uint64(postCfg.BitsPerLabel) / 8

		fmt.Printf("Name: %s\n", name)
		fmt.Printf("\tStatus %s\n", status.GetStatus().State)
		fmt.Printf("\tProgress %d GiB / %d GiB %f\n", completed, commitmentSize, completed/commitmentSize)
		fmt.Println()
	}

	return nil
}

// LoadConfigFromFile tries to load configuration file if the config parameter was specified.
func LoadConfigFromFile(path string) (*config.Config, error) {
	// read in default config if passed as param using viper
	if err := config.LoadConfig(path, viper.GetViper()); err != nil {
		return nil, err
	}
	conf := config.MainnetConfig()
	if name := viper.GetString("preset"); len(name) > 0 {
		preset, err := presets.Get(name)
		if err != nil {
			return nil, err
		}
		conf = preset
	}

	hook := mapstructure.ComposeDecodeHookFunc(
		mapstructure.StringToTimeDurationHookFunc(),
		mapstructure.StringToSliceHookFunc(","),
		mapstructureutil.BigRatDecodeFunc(),
		mapstructure.TextUnmarshallerHookFunc(),
	)

	// load config if it was loaded to the viper
	if err := viper.Unmarshal(&conf, viper.DecodeHook(hook), withZeroFields()); err != nil {
		return nil, fmt.Errorf("unmarshal viper: %w", err)
	}
	return &conf, nil
}

func withZeroFields() viper.DecoderConfigOption {
	return func(cfg *mapstructure.DecoderConfig) {
		cfg.ZeroFields = true
	}
}
