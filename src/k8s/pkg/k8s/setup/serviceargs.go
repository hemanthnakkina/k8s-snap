package setup

import (
	"fmt"

	"github.com/canonical/k8s/pkg/utils"
)

// InitServiceArgs handles the setup of services arguments.
//   - For each service, copies the default arguments files from the snap under $SNAP_DATA/args
//   - Note that the `k8sd` service is already configured in the snap install hook and thus not included here
func InitServiceArgs() error {
	for _, service := range []string{"containerd", "k8s-dqlite", "kube-apiserver", "kube-controller-manager", "kube-proxy", "kube-scheduler", "kubelet"} {
		err := utils.CopyFile(utils.SnapPath("k8s/args", service), utils.SnapDataPath("args", service))
		if err != nil {
			return fmt.Errorf("failed to copy %s args: %w", service, err)
		}
	}

	return nil
}