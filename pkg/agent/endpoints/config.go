package endpoints

import (
	"net"

	"github.com/sirupsen/logrus"
	"github.com/spiffe/spire/pkg/agent/catalog"
	"github.com/spiffe/spire/pkg/agent/manager"
	"github.com/spiffe/spire/pkg/common/telemetry"

	"google.golang.org/grpc"
)

type Config struct {
	BindAddr *net.UnixAddr

	GRPCHook func(*grpc.Server) error

	Catalog catalog.Catalog
	Manager manager.Manager

	Log logrus.FieldLogger
	Tel telemetry.Sink
}

func New(c *Config) *endpoints {
	return &endpoints{
		c: c,
	}
}
