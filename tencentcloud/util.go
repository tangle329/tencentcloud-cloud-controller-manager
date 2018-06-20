package tencentcloud

import (
	"fmt"
	"net"
)

func (cloud *Cloud) getNodeHostIP(name string) (string, error) {
	addrs, err := net.LookupHost(name)
	if err != nil {
		return "", err
	}

	// FIXME: Should have a more reasonable way to choose ip
	if !cloud.isPrivateIP(addrs[0]) {
		return "", fmt.Errorf("Resovled an invalid IP(%s)", addrs[0])
	}

	return addrs[0], nil
}

func (cloud *Cloud) isPrivateIP(ip string) bool {
	private := false
	IP := net.ParseIP(ip)
	if IP == nil {
		return private
	} else {
		_, private24BitBlock, _ := net.ParseCIDR("10.0.0.0/8")
		_, private20BitBlock, _ := net.ParseCIDR("172.16.0.0/12")
		_, private16BitBlock, _ := net.ParseCIDR("192.168.0.0/16")
		private = private24BitBlock.Contains(IP) || private20BitBlock.Contains(IP) || private16BitBlock.Contains(IP)
	}
	return private
}
