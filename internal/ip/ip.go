package ip

import (
	"encoding/binary"
	"errors"
	"fmt"
	"math/big"
	"net"
	"strconv"
	"strings"
)

// CIDR parses a CIDR and returns a human-readable report.
func CIDR(in string) (string, error) {
	ip, ipNet, err := net.ParseCIDR(in)
	if err != nil {
		return "", err
	}
	ones, bits := ipNet.Mask.Size()
	hostBits := bits - ones

	count := new(big.Int).Lsh(big.NewInt(1), uint(hostBits))

	var b strings.Builder
	fmt.Fprintf(&b, "Network address: %s\n", ipNet.IP.String())
	fmt.Fprintf(&b, "Netmask: %s\n", maskString(ipNet.Mask))
	fmt.Fprintf(&b, "Prefix length: %d\n", ones)
	fmt.Fprintf(&b, "Number of addresses: %s\n", count.String())

	if v4 := ipNet.IP.To4(); v4 != nil {
		network := binary.BigEndian.Uint32(v4)
		mask := binary.BigEndian.Uint32(net.IP(ipNet.Mask).To4())
		broadcast := network | ^mask
		first := network
		last := broadcast
		if hostBits >= 2 {
			first = network + 1
			last = broadcast - 1
		}
		fmt.Fprintf(&b, "First host: %s\n", uint32ToIP(first))
		fmt.Fprintf(&b, "Last host: %s\n", uint32ToIP(last))
		fmt.Fprintf(&b, "Broadcast: %s\n", uint32ToIP(broadcast))
	}
	_ = ip
	return strings.TrimRight(b.String(), "\n"), nil
}

func maskString(m net.IPMask) string {
	if len(m) == 4 {
		return fmt.Sprintf("%d.%d.%d.%d", m[0], m[1], m[2], m[3])
	}
	return m.String()
}

func uint32ToIP(v uint32) string {
	b := make([]byte, 4)
	binary.BigEndian.PutUint32(b, v)
	return net.IP(b).String()
}

// ToInt parses an IPv4 address and returns its uint32 value as a decimal string.
func ToInt(in string) (string, error) {
	ip := net.ParseIP(strings.TrimSpace(in))
	if ip == nil {
		return "", errors.New("invalid IP address")
	}
	v4 := ip.To4()
	if v4 == nil {
		return "", errors.New("not an IPv4 address")
	}
	return strconv.FormatUint(uint64(binary.BigEndian.Uint32(v4)), 10), nil
}

// FromInt parses a decimal uint32 and returns the dotted IPv4 address.
func FromInt(in string) (string, error) {
	n, err := strconv.ParseUint(strings.TrimSpace(in), 10, 32)
	if err != nil {
		return "", err
	}
	return uint32ToIP(uint32(n)), nil
}
