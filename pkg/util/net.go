package util

import "net"

// GetIntranetIP gets the intranet ip of this machine.
func GetIntranetIP() string {
	addresses, err := net.InterfaceAddrs()
	if err != nil {
		return "127.0.0.1"
	}

	for _, address := range addresses {
		if ipNet, ok := address.(*net.IPNet); ok && !ipNet.IP.IsLoopback() {
			if ipNet.IP.To4() != nil {
				return ipNet.IP.String()
			}
		}
	}
	return "127.0.0.1"
}
