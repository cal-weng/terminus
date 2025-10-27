package terminus

import (
	"context"
	"fmt"
	"path"
	"time"

	"github.com/beclab/Olares/cli/pkg/core/logger"
	"github.com/beclab/Olares/cli/pkg/storage"

	"github.com/beclab/Olares/cli/pkg/common"
	"github.com/beclab/Olares/cli/pkg/core/connector"
	"github.com/beclab/Olares/cli/pkg/core/task"
	"github.com/beclab/Olares/cli/pkg/utils"
	ctrl "sigs.k8s.io/controller-runtime"
)

type ClearBFLValues struct {
	common.KubeAction
}

func (c *ClearBFLValues) Execute(runtime connector.Runtime) error {
	_, _ = runtime.GetRunner().SudoCmd(fmt.Sprintf("cat /dev/null > %s/wizard/config/launcher/values.yaml", runtime.GetInstallerDir()), false, false)

	return nil
}

type InstallBFL struct {
	common.KubeAction
}

func (t *InstallBFL) Execute(runtime connector.Runtime) error {
	var ns = fmt.Sprintf("user-space-%s", t.KubeConf.Arg.User.UserName)

	config, err := ctrl.GetConfig()
	if err != nil {
		return err
	}
	actionConfig, settings, err := utils.InitConfig(config, ns)
	if err != nil {
		return err
	}

	var ctx, cancel = context.WithTimeout(context.Background(), 3*time.Minute)
	defer cancel()

	var r = utils.Random()
	var key = fmt.Sprintf("bytetrade_bfl_%d", r)
	var secret, _ = utils.GeneratePassword(16)

	var launchName = fmt.Sprintf("launcher-%s", t.KubeConf.Arg.User.UserName)
	var launchPath = path.Join(runtime.GetInstallerDir(), "wizard/config/launcher")
	vals := make(map[string]interface{})
	vals["bfl"] = map[string]interface{}{
		"nodeport":               30883,
		"nodeport_ingress_http":  30083,
		"nodeport_ingress_https": 30082,
		"username":               t.KubeConf.Arg.User.UserName,
		"admin_user":             true,
		"appKey":                 key,
		"appSecret":              secret,
	}
	vals[common.HelmValuesKeyOlaresRootFSPath] = storage.OlaresRootDir

	if err := utils.UpgradeCharts(ctx, actionConfig, settings, launchName, launchPath, "", ns, vals, false); err != nil {
		return err
	}

	return nil
}

type InstallLauncherModule struct {
	common.KubeModule
}

func (m *InstallLauncherModule) Init() {
	logger.InfoInstallationProgress("Installing launcher ...")
	m.Name = "InstallLauncher"

	installBFL := &task.LocalTask{
		Name:   "InstallLauncher",
		Desc:   "InstallLauncher",
		Action: new(InstallBFL),
		Retry:  3,
		Delay:  30 * time.Second,
	}

	checkBFLRunning := &task.LocalTask{
		Name: "CheckLauncherStatus",
		Action: &CheckPodsRunning{
			labels: map[string][]string{
				fmt.Sprintf("user-space-%s", m.KubeConf.Arg.User.UserName): {"tier=bfl"},
			},
		},
		Retry: 20,
		Delay: 10 * time.Second,
	}

	m.Tasks = []task.Interface{
		installBFL,
		checkBFLRunning,
	}
}
