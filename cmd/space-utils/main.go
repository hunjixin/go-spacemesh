// go-spacemesh is a golang implementation of the Spacemesh node.
// See - https://spacemesh.io
package main

import (
	"context"
	"encoding/json"
	"fmt"
	_ "net/http/pprof"
	"os"
	"time"

	units "github.com/docker/go-units"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/mitchellh/mapstructure"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/push"
	"github.com/rodaine/table"
	pb "github.com/spacemeshos/api/release/go/spacemesh/v1"
	"github.com/spacemeshos/go-spacemesh/common/types"
	"github.com/spacemeshos/go-spacemesh/config"
	"github.com/spacemeshos/go-spacemesh/config/presets"
	"github.com/spacemeshos/go-spacemesh/node/mapstructureutil"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/types/known/emptypb"
)

var (
	path string
)

var rootCmd = &cobra.Command{
	Use:   "space-utils",
	Short: "utils for spacemesh",
}

var printCmd = &cobra.Command{
	Use:   "print",
	Short: "print status",
	Run: func(c *cobra.Command, args []string) {
		ctx := context.Background()
		if err := print(ctx, path); err != nil {
			fmt.Println(err.Error())
		}
	},
}
var metricsCmd = &cobra.Command{
	Use:   "metrics",
	Short: "provider metrics for prometheus",
	Run: func(c *cobra.Command, args []string) {
		ctx := context.Background()
		if err := pushDataToPrometheus(ctx, path); err != nil {
			fmt.Println(err.Error())
		}
	},
}

func init() {
	rootCmd.Flags().StringVar(&path, "path", "", "config to read")
}

func main() {
	rootCmd.AddCommand(printCmd, metricsCmd)
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

type MachineInfo struct {
	PublicUrl  string `json:"publicUrl"`
	PrivateUrl string `json:"privateUrl"`
}

func print(ctx context.Context, path string) error {
	urlMap := make(map[string]MachineInfo)
	cfgData, err := os.ReadFile(path)
	if err != nil {
		return err
	}

	err = json.Unmarshal(cfgData, &urlMap)
	if err != nil {
		return err
	}

	tbl := table.New("机器", "奖励地址", "节点", "GenesisID", "同步状态", "状态", "完成情况")
	for machine, url := range urlMap {
		status, err := getMachineInfo(ctx, machine, url)
		if err != nil {
			tbl.AddRow(machine, err.Error())
			continue
		}

		commitStatus := fmt.Sprintf("%s / %s  %s", status.CompletedSize, status.CommitmentSize, status.Percent)
		var syncStatus string
		if status.IsSynced {
			syncStatus = fmt.Sprintf("同步成功 %d/%d", status.CurrentLayerId, status.GenesisEndLayer)
		} else {
			syncStatus = fmt.Sprintf("同步失败 %d/%d", status.CurrentLayerId, status.GenesisEndLayer)
		}
		tbl.AddRow(machine, status.CoinBase, status.NodeId, status.GenesisId, syncStatus, status.State, commitStatus)
	}
	tbl.Print()
	return nil
}

func pushDataToPrometheus(ctx context.Context, path string) error {
	urlMap := make(map[string]MachineInfo)
	cfgData, err := os.ReadFile(path)
	if err != nil {
		return err
	}

	err = json.Unmarshal(cfgData, &urlMap)
	if err != nil {
		return err
	}

	ticker := time.NewTicker(time.Second * 5)
	go func() {
		select {
		case <-ctx.Done():
			return
		case <-ticker.C:
			var info = promauto.NewGaugeVec(prometheus.GaugeOpts{
				Name: "my_table",
				Help: "This is my table",
			}, []string{"机器", "奖励地址", "节点", "GenesisID", "同步状态", "状态", "完成情况"})

			for machine, url := range urlMap {
				status, err := getMachineInfo(ctx, machine, url)
				if err != nil {
					fmt.Println(err)
					info.WithLabelValues(machine, "", "", "", "", "", "")
					continue
				}

				commitStatus := fmt.Sprintf("%s / %s  %s", status.CompletedSize, status.CommitmentSize, status.Percent)
				var syncStatus string
				if status.IsSynced {
					syncStatus = fmt.Sprintf("同步成功 %d/%d", status.CurrentLayerId, status.GenesisEndLayer)
				} else {
					syncStatus = fmt.Sprintf("同步失败 %d/%d", status.CurrentLayerId, status.GenesisEndLayer)
				}

				info.WithLabelValues(machine, status.CoinBase, status.NodeId, status.GenesisId, syncStatus, status.State, commitStatus)
			}

			pusher := push.New("http://localhost:9091", "机器状态")
			pusher.Collector(info)
			if err := pusher.Push(); err != nil {
				fmt.Println("Error pushing to Pushgateway:", err)
			}
		}
	}()
	return nil
}

func getMachineInfo(ctx context.Context, name string, url MachineInfo) (*Status, error) {
	publicConn, err := grpc.Dial(url.PublicUrl, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	privateConn, err := grpc.Dial(url.PrivateUrl, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	smesherClient := pb.NewSmesherServiceClient(privateConn)
	status, err := smesherClient.PostSetupStatus(ctx, &empty.Empty{})
	if err != nil {
		return nil, err
	}

	postCfg, err := smesherClient.PostConfig(ctx, &empty.Empty{})
	if err != nil {
		return nil, err
	}

	smeshId, err := smesherClient.SmesherID(ctx, &emptypb.Empty{})
	if err != nil {
		return nil, err
	}

	coinBase, err := smesherClient.Coinbase(ctx, &empty.Empty{})
	if err != nil {
		return nil, err
	}

	nodeClient := pb.NewNodeServiceClient(publicConn)
	nodeStatus, err := nodeClient.Status(ctx, &pb.StatusRequest{})
	if err != nil {
		return nil, err
	}

	nodeInfo, err := nodeClient.NodeInfo(ctx, &empty.Empty{})
	if err != nil {
		return nil, err
	}

	meshClient := pb.NewMeshServiceClient(publicConn)
	genesisIDResp, err := meshClient.GenesisID(ctx, &pb.GenesisIDRequest{})
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

	genesisId := types.Hash20{}
	genesisId.SetBytes(genesisIDResp.GenesisId)
	return &Status{
		Machine:        name,
		CompletedSize:  units.BytesSize(float64(completed)),
		CommitmentSize: units.BytesSize(float64(commitmentSize)),
		Percent:        percent,
		State:          status.GetStatus().State.String(),

		CoinBase:        coinBase.AccountId.Address,
		NodeId:          types.BytesToNodeID(smeshId.PublicKey).String(), //todo
		GenesisId:       genesisId.Hex(),
		IsSynced:        nodeStatus.GetStatus().IsSynced,
		CurrentLayerId:  int(nodeStatus.GetStatus().SyncedLayer.Number),
		GenesisEndLayer: int(nodeInfo.GetEffectiveGenesis()),
	}, nil
}

type Status struct {
	State          string `json:"state"`
	Machine        string `json:"machine"`
	CompletedSize  string `json:"completedSize"`
	CommitmentSize string `json:"commitmentSize"`
	Percent        string `json:"percent"`

	CoinBase  string `json:"coinbase"`
	NodeId    string `json:"nodeId"`
	GenesisId string `json:"genesisId"`

	IsSynced        bool `json:"percent"`
	CurrentLayerId  int  `json:"currentLayerId"`
	GenesisEndLayer int  `json:"genesisEndLayer"`
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
