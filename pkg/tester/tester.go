package tester

import (
	"fmt"
	"net"

	"github.com/rimaulana/netest/pkg/logger"
)

// Tester desc
type Tester struct {
	Address string
	Port    int
	Proto   string
	Logger  logger.Logger
}

// Connect desc
func (t *Tester) Connect() error {
	t.Logger.Infof("Initiating connection to %s://%s:%d", t.Proto, t.Address, t.Port)
	conn, err := net.Dial(t.Proto, fmt.Sprintf("%s:%d", t.Address, t.Port))
	if err != nil {
		t.Logger.Errorf("Unable to connect to %s:%d, error: %s", t.Address, t.Port, err.Error())
		return err
	}
	t.Logger.Info("Connected")
	t.Logger.Info("Closing connection")
	defer conn.Close()
	return nil
}
