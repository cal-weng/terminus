package utils

import (
	"context"
	"fmt"
	"os"
	"os/exec"
	"strings"

	"k8s.io/klog/v2"
)

type Device struct {
	Name        string
	Type        string
	State       string
	Connection  string
	Ipv4Gateway string
	Ipv6Gateway string
	Ipv4DNS     string
	Ipv6DNS     string
	Ipv4Address string
	Ipv4Mask    string
	Ipv6Address string
	Method      string
}

func findCommand(ctx context.Context, cmdName string) (cmdPath string, err error) {
	cmd := exec.CommandContext(ctx, "bash", "-c", fmt.Sprintf("command -v %s", cmdName))
	cmd.Env = os.Environ()
	output, err := cmd.Output()
	if err != nil {
		klog.Error("find nmcli error, ", err)
		return
	}

	cmdPath = strings.TrimSpace(string(output))

	return
}
