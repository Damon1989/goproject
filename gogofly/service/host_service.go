package service

import (
	"context"
	"fmt"

	"github.com/apenella/go-ansible/pkg/adhoc"
	"github.com/apenella/go-ansible/pkg/options"
	"github.com/damon/gogofly/service/dto"
	"github.com/spf13/viper"
)

var hostService *HostService

type HostService struct {
	BaseService
}

func NewHostService() *HostService {
	if hostService == nil {
		hostService = &HostService{}
	}
	return hostService
}

func (m *HostService) Shutdown(iShutdownHostDTO dto.ShutdownHostDTO) error {
	var errResult error
	stHostIP := iShutdownHostDTO.HostIP
	fmt.Println("Shutdown host:", stHostIP)

	return errResult

	ansibleConnectionOptions := &options.AnsibleConnectionOptions{
		Connection: "ssh",
		User:       viper.GetString("ansible.user.name"),
	}
	ansibleAdhocOptions := &adhoc.AnsibleAdhocOptions{
		Inventory:  fmt.Sprintf("%s,", stHostIP),
		ModuleName: "command",
		Args:       viper.GetString("ansible.ShutdownHost.Args"),
		ExtraVars: map[string]interface{}{
			"ansible_password": viper.GetString("ansible.user.password"),
		},
	}
	ansibleAdhoc := adhoc.AnsibleAdhocCmd{
		Pattern:           "all",
		Options:           ansibleAdhocOptions,
		ConnectionOptions: ansibleConnectionOptions,
		StdoutCallback:    "oneline",
	}
	errResult = ansibleAdhoc.Run(context.TODO())
	return errResult
}
