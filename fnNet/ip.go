package fnNet

import (
	"net"
)

func IP() (res []string, err error) {

	var interfaces []net.Interface
	if interfaces, err = net.Interfaces(); err != nil {
		return
	}

	for _, face := range interfaces {
		if face.Flags&net.FlagUp == 0 || face.Flags&net.FlagLoopback != 0 {
			continue
		}

		var addresses []net.Addr
		if addresses, err = face.Addrs(); err != nil {
			continue
		}

		for _, addr := range addresses {
			var ip net.IP
			switch v := addr.(type) {
			case *net.IPNet:
				ip = v.IP
			case *net.IPAddr:
				ip = v.IP
			}

			if ip != nil && ip.To4() != nil {
				res = append(res, ip.String())
			}
		}
	}

	return
}
