package fwdnet

import (
	"net"
	"os/exec"

	"github.com/txn2/kubefwd/pkg/fwdIp"
)

// ReadyInterface prepares a local IP address on
// the loopback interface.
func ReadyInterface(opts fwdIp.ForwardIPOpts) (net.IP, error) {
	return fwdIp.GetIp(opts)
}

// RemoveInterfaceAlias can remove the Interface alias after port forwarding.
// if -alias command get err, just print the error and continue.
func RemoveInterfaceAlias(ip net.IP) {
	cmd := "ifconfig"
	args := []string{"lo0", "-alias", ip.String()}
	if err := exec.Command(cmd, args...).Run(); err != nil {
		// suppress for now
		// @todo research alternative to ifconfig
		// @todo suggest ifconfig or alternative
		// @todo research libs for interface management
		//fmt.Println("Cannot ifconfig lo0 -alias " + ip.String() + "\r\n" + err.Error())
	}
}
