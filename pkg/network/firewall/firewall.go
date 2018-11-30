package firewall

import (
	"os"
	"os/exec"
	"time"

	"github.com/golang/glog"
	"k8s.io/apimachinery/pkg/util/wait"
)

func SetupEbtables(quit <-chan struct{}) {
	ebtableFile := "/etc/sysconfig/galaxy-ebtable-filter"
	go wait.Until(func() {
		glog.Infof("starting to ensure ebtable rules")
		ebtablesRestore, err := exec.LookPath("ebtables-restore")
		if err != nil {
			glog.Warning("ebtables unavailable - unable to locate ebtables-restore")
			return
		}
		fi, err := os.Open(ebtableFile)
		if err != nil {
			glog.Infof("%s not exists, ignore ebtables restore", ebtableFile)
			return
		}
		cmd := &exec.Cmd{
			Path:  ebtablesRestore,
			Stdin: fi,
		}
		ret, err := cmd.CombinedOutput()
		if err != nil {
			glog.Warningf("Error executing ebtables restore %v, %s", err, string(ret))
			return
		}
		glog.Infof("executed ebtables restore %s", string(ret))
	}, 5*time.Minute, quit)
}
