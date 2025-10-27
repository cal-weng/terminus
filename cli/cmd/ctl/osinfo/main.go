package osinfo

import (
	"fmt"

	"github.com/beclab/Olares/cli/pkg/common"
	"github.com/beclab/Olares/cli/pkg/core/connector"
	"github.com/spf13/cobra"
)

func NewCmdInfo() *cobra.Command {
	infoCmd := &cobra.Command{
		Use:   "osinfo",
		Short: "Print system information, etc.",
		Long:  "help for printing info",
	}
	infoCmd.AddCommand(showInfoCommand())
	infoCmd.AddCommand(showAllCommand())

	return infoCmd
}

func showInfoCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "show",
		Short: "Print os information",
		Long:  "help for printing os info",
		Run: func(cmd *cobra.Command, args []string) {
			systemInfo := connector.GetSystemInfo()
			host := systemInfo.HostInfo
			fmt.Printf(`OS_TYPE=%s
OS_PLATFORM=%s
OS_ARCH=%s
OS_VERSION=%s
OS_KERNEL=%s
OS_INFO=%s
`, host.OsType, host.OsPlatformFamily, host.OsArch, host.OsVersion, host.OsKernel, host.OsInfo)
		},
	}
	return cmd
}

func showAllCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "showall",
		Short: "Print full os information",
		Long:  "help for printing os info",
		Run: func(cmd *cobra.Command, args []string) {
			arg := common.NewArgument()

			arg.SystemInfo.Print()
		},
	}
	return cmd
}
