package patch

import (
	"github.com/beclab/Olares/cli/pkg/binaries"
	"github.com/beclab/Olares/cli/pkg/common"
	"github.com/beclab/Olares/cli/pkg/core/logger"
	"github.com/beclab/Olares/cli/pkg/core/prepare"
	"github.com/beclab/Olares/cli/pkg/core/task"
	"github.com/beclab/Olares/cli/pkg/manifest"
)

type InstallDepsModule struct {
	common.KubeModule
	manifest.ManifestModule
}

func (m *InstallDepsModule) Init() {
	logger.InfoInstallationProgress("installing and configuring OS dependencies ...")
	m.Name = "InstallDeps"

	patchAppArmor := &task.RemoteTask{
		Name:  "PatchAppArmor",
		Hosts: m.Runtime.GetAllHosts(),
		Prepare: &prepare.PrepareCollection{
			&binaries.Ubuntu24AppArmorCheck{},
		},
		Action: &binaries.InstallAppArmorTask{
			ManifestAction: manifest.ManifestAction{
				BaseDir:  m.BaseDir,
				Manifest: m.Manifest},
		},
		Parallel: false,
		Retry:    0,
	}

	raspbianCheck := &task.RemoteTask{
		Name:     "RaspbianCheck",
		Hosts:    m.Runtime.GetAllHosts(),
		Action:   new(RaspbianCheckTask),
		Parallel: false,
		Retry:    0,
	}

	correctHostname := &task.RemoteTask{
		Name:     "CorrectHostname",
		Hosts:    m.Runtime.GetAllHosts(),
		Action:   new(CorrectHostname),
		Parallel: false,
		Retry:    0,
	}

	disableDNS := &task.RemoteTask{
		Name:     "DisableLocalDNS",
		Hosts:    m.Runtime.GetAllHosts(),
		Action:   new(DisableLocalDNSTask),
		Parallel: false,
		Retry:    0,
	}

	patchOs := &task.RemoteTask{
		Name:   "PatchOs",
		Hosts:  m.Runtime.GetAllHosts(),
		Action: new(PatchTask),
		Retry:  0,
	}

	enableSSHTask := &task.RemoteTask{
		Name:     "EnableSSH",
		Hosts:    m.Runtime.GetAllHosts(),
		Action:   new(EnableSSHTask),
		Parallel: false,
		Retry:    1,
	}

	m.Tasks = []task.Interface{
		patchAppArmor,
		raspbianCheck,
		correctHostname,
		disableDNS,
		patchOs,
		enableSSHTask,
	}
}
