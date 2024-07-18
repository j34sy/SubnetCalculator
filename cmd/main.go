package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/j34sy/SubnetCalculator/pkg/subnetcalc"
)

func main() {
	fmt.Println("Subnet Calculator")
	fmt.Println("-----------------")
	fmt.Println("Enter an IP address and CIDR to calculate the subnet mask, network address, broadcast address, usable host range, total hosts, and usable hosts.")
	fmt.Println("Example: 127.0.0.1/24")
	fmt.Println("Enter 'exit' to quit.")
	fmt.Println()

	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("Subnet Calculator> ")
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println(err)
			return
		}
		input = strings.Trim(input, "\n")

		if strings.TrimSpace(input) == "exit" {
			break
		}

		test := parseInput(strings.TrimSpace(input))
		if test == nil {
			continue
		}
		test.Calculate()
		fmt.Printf("IPv4 address: %s \n", test.GetIPv4Address())
		fmt.Printf("CIDR: %s \n", test.GetCIDR())
		fmt.Printf("Subnet Mask: %s \n", test.GetSubnetMask())
		fmt.Printf("Network Address: %s \n", test.GetNetworkAddress())
		fmt.Printf("Broadcast Address: %s \n", test.GetBroadcastAddress())
		fmt.Printf("Usable Host Range: %s \n", test.GetUsableHostRange())
		fmt.Printf("Total Hosts: %s \n", test.GetTotalHosts())
		fmt.Printf("Usable Hosts: %s \n", test.GetUsableHosts())
		fmt.Println()

	}
}

func parseInput(input string) *subnetcalc.IPv4Address {
	var ip [4]int
	var cidr uint8
	slashSplit := strings.Split(input, "/")
	if len(slashSplit) != 2 {
		fmt.Println("Invalid input. Please enter an IP address and CIDR.")
		return nil
	}
	ipSplit := strings.Split(slashSplit[0], ".")
	if len(ipSplit) != 4 {
		fmt.Println("Invalid input. Please enter a valid IP address.")
		return nil
	}
	cidrInt, err := strconv.Atoi(slashSplit[1])
	if err != nil {
		fmt.Println("Invalid input. Please enter a valid CIDR.")
		return nil
	}
	if cidrInt > 0 || cidrInt < 33 {
		cidr = uint8(cidrInt)
	} else {
		fmt.Println("Invalid input. Please enter a valid CIDR.")
		return nil
	}

	for i, octet := range ipSplit {
		octetInt, err := strconv.Atoi(octet)
		if err != nil {
			fmt.Println("Invalid input. Please enter a valid IP address.")
			return nil
		}
		if octetInt > 0 || octetInt < 256 {
			ip[i] = octetInt
		} else {
			fmt.Println("Invalid input. Please enter a valid IP address.")
			return nil
		}
	}

	return subnetcalc.NewIPv4Address(ip, cidr)
}
