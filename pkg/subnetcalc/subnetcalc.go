package subnetcalc

import "fmt"

// IPv4Address represents an IPv4 address along with related information.
type IPv4Address struct {
	Octets           [4]int    // The four octets of the IPv4 address.
	cidr             uint8     // The CIDR notation for the subnet mask.
	subnetMask       [4]int    // The subnet mask derived from the CIDR notation.
	networkAddress   [4]int    // The network address calculated from the IPv4 address and subnet mask.
	broadcastAddress [4]int    // The broadcast address calculated from the IPv4 address and subnet mask.
	usableHostRange  [2][4]int // The range of usable host addresses within the network.
	totalHosts       int       // The total number of possible host addresses in the network.
	usablesHosts     int       // The number of usable host addresses in the network.
}

func (ip *IPv4Address) SetCIDR(cidr uint8) {
	ip.cidr = cidr
}

func (ip *IPv4Address) SetOctets(octets [4]int) {
	ip.Octets = octets
}

func (ip *IPv4Address) GetIPv4Address() string {
	return fmt.Sprintf("%d.%d.%d.%d", ip.Octets[0], ip.Octets[1], ip.Octets[2], ip.Octets[3])
}

func (ip *IPv4Address) GetCIDR() string {
	return fmt.Sprintf("/%d", ip.cidr)
}

func (ip *IPv4Address) GetSubnetMask() string {
	return fmt.Sprintf("%d.%d.%d.%d", ip.subnetMask[0], ip.subnetMask[1], ip.subnetMask[2], ip.subnetMask[3])
}

func (ip *IPv4Address) GetNetworkAddress() string {
	return fmt.Sprintf("%d.%d.%d.%d", ip.networkAddress[0], ip.networkAddress[1], ip.networkAddress[2], ip.networkAddress[3])
}

func (ip *IPv4Address) GetBroadcastAddress() string {
	return fmt.Sprintf("%d.%d.%d.%d", ip.broadcastAddress[0], ip.broadcastAddress[1], ip.broadcastAddress[2], ip.broadcastAddress[3])
}

func (ip *IPv4Address) GetUsableHostRange() string {
	return fmt.Sprintf("%d.%d.%d.%d - %d.%d.%d.%d", ip.usableHostRange[0][0], ip.usableHostRange[0][1], ip.usableHostRange[0][2], ip.usableHostRange[0][3], ip.usableHostRange[1][0], ip.usableHostRange[1][1], ip.usableHostRange[1][2], ip.usableHostRange[1][3])
}

func (ip *IPv4Address) GetTotalHosts() string {
	return fmt.Sprintf("%d", ip.totalHosts)
}

func (ip *IPv4Address) GetUsableHosts() string {
	return fmt.Sprintf("%d", ip.usablesHosts)
}

func (ip *IPv4Address) Calculate() {
	ip.calculateSubnetMask()
	ip.calculateNetworkAddress()
	ip.calculateBroadcastAddress()
	ip.calculateUsableHostRange()
	ip.calculateTotalHosts()
	ip.calculateUsableHosts()
}

func (ip *IPv4Address) calculateSubnetMask() {
	ip.subnetMask = [4]int{0, 0, 0, 0}
	for i := 0; i < int(ip.cidr); i++ {
		ip.subnetMask[i/8] |= 1 << uint(7-i%8)
	}
}

func (ip *IPv4Address) calculateNetworkAddress() {
	for i := 0; i < 4; i++ {
		ip.networkAddress[i] = ip.Octets[i] & ip.subnetMask[i]
	}
}

func (ip *IPv4Address) calculateBroadcastAddress() {
	for i := 0; i < 4; i++ {
		ip.broadcastAddress[i] = ip.networkAddress[i] | (^ip.subnetMask[i] & 0xff)
	}
}

func (ip *IPv4Address) calculateUsableHostRange() {
	ip.usableHostRange = [2][4]int{{0, 0, 0, 0}, {0, 0, 0, 0}}
	ip.usableHostRange[0][0] = ip.networkAddress[0]
	ip.usableHostRange[1][0] = ip.broadcastAddress[0]
	ip.usableHostRange[0][1] = ip.networkAddress[1]
	ip.usableHostRange[1][1] = ip.broadcastAddress[1]
	ip.usableHostRange[0][2] = ip.networkAddress[2]
	ip.usableHostRange[1][2] = ip.broadcastAddress[2]
	ip.usableHostRange[0][3] = ip.networkAddress[3] + 1
	ip.usableHostRange[1][3] = ip.broadcastAddress[3] - 1
}

func (ip *IPv4Address) calculateTotalHosts() {
	ip.totalHosts = 1
	for i := 0; i < 32-int(ip.cidr); i++ {
		ip.totalHosts *= 2
	}
}

func (ip *IPv4Address) calculateUsableHosts() {
	ip.usablesHosts = ip.totalHosts - 2
}

func NewIPv4Address(octets [4]int, cidr uint8) *IPv4Address {
	return &IPv4Address{
		Octets: octets,
		cidr:   cidr,
	}
}
