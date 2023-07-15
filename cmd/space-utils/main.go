// go-spacemesh is a golang implementation of the Spacemesh node.
// See - https://spacemesh.io
package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	_ "net/http/pprof"
	"os"

	units "github.com/docker/go-units"
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/mitchellh/mapstructure"
	"github.com/rodaine/table"
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

	tbl := table.New("机器", "状态", "完成量", "目标量", "完成比例")
	for machine, url := range urlMap {
		status, err := getMachineInfo(ctx, machine, url)
		if err != nil {
			tbl.AddRow(machine, err.Error())
			continue
		}
		tbl.AddRow(machine, status.State, status.CompletedSize, status.CommitmentSize, status.Percent)
	}
	tbl.Print()
	return nil
}

func runServe(data map[string]string) error {
	r := gin.Default()
	r.Use(static.Serve("/", static.LocalFile("./dist", false)))
	myCors := cors.DefaultConfig()
	myCors.AllowAllOrigins = true
	r.Use(cors.New(myCors))

	r.GET("/api/data", func(c *gin.Context) {
		c.JSON(http.StatusOK, data)
	})
	r.GET("/api/machine/info/:machine", func(c *gin.Context) {
		machine := c.Param("machine")
		url, ok := data[machine]
		if !ok {
			c.AbortWithError(http.StatusInternalServerError, errors.New("machine not found"))
			return
		}
		status, err := getMachineInfo(c, machine, url)
		if err != nil {
			c.AbortWithError(http.StatusInternalServerError, err)
			return
		}
		c.JSON(http.StatusOK, status)
	})
	return r.Run()
}

func getMachineInfo(ctx context.Context, name, url string) (*Status, error) {
	conn, err := grpc.Dial(url, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	client := pb.NewSmesherServiceClient(conn)
	status, err := client.PostSetupStatus(ctx, &empty.Empty{})
	if err != nil {
		return nil, err
	}

	postCfg, err := client.PostConfig(ctx, &empty.Empty{})
	if err != nil {
		return nil, err
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

	commitmentSize := postCfg.LabelsPerUnit * uint64(postCfg.BitsPerLabel) * (uint64(status.GetStatus().GetOpts().NumUnits)) / 8
	completed := status.GetStatus().NumLabelsWritten * uint64(postCfg.BitsPerLabel) / 8

	percent := fmt.Sprintf("%.2f %%", 100*(float64(completed)/float64(commitmentSize)))
	fmt.Printf("Machine: %s\n", name)
	fmt.Printf("\tStatus %s\n", status.GetStatus().State)
	fmt.Printf("\tProgress %s / %s %s\n", units.BytesSize(float64(completed)), units.BytesSize(float64(commitmentSize)), percent)
	fmt.Println()
	return &Status{
		Machine:        name,
		CompletedSize:  units.BytesSize(float64(completed)),
		CommitmentSize: units.BytesSize(float64(commitmentSize)),
		Percent:        percent,
		State:          status.GetStatus().State.String(),
	}, nil
}

type Status struct {
	State          string `json:"state"`
	Machine        string `json:"machine"`
	CompletedSize  string `json:"completedSize"`
	CommitmentSize string `json:"commitmentSize"`
	Percent        string `json:"percent"`
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
