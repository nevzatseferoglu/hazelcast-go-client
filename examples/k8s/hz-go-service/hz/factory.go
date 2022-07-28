package hz

import (
	"context"
	"fmt"
	"os"

	"github.com/hazelcast/hazelcast-go-client"
	"github.com/hazelcast/hazelcast-go-client/logger"
)

const withoutK8s = "HZ_GO_SERVICE_WITHOUT_K8S"

// ClientInfo contains info about client
type ClientInfo struct {
	ClientName    string
	ClientRunning bool
	MapSize       int
}

// NewHzClient returns Hazelcast client instance with default config.
func NewHzClient(ctx context.Context) (*hazelcast.Client, error) {
	config := hazelcast.Config{
		ClientName: "hz-go-service-client",
	}
	cc := &config.Cluster
	_, locally := os.LookupEnv(withoutK8s)
	if locally {
		cc.Network.SetAddresses(fmt.Sprintf("%s:%s", "localhost", "5701"))
	} else {
		cc.Network.SetAddresses(fmt.Sprintf("%s:%s", "hazelcast-sample.default.svc", "5701"))
	}
	// Unisocket network configuration is not a mandatory setting.
	cc.Unisocket = true
	config.Logger.Level = logger.InfoLevel
	client, err := hazelcast.StartNewClientWithConfig(ctx, config)
	if err != nil {
		return nil, err
	}
	return client, nil
}
