package lcjune

import (
	"testing"
)

func TestIpAddrIpv4(t *testing.T) {
	var r string
	var addr string

	var ripv4 string = "IPv4"

	var ipv4 []string = []string{
		"172.16.254.1",
		"1.1.1.1",
		"12.12.12.12",
		"1.12.123.2",
		"192.0.0.1",
	}

	for _, addr = range ipv4 {
		r = validIPAddress(addr)
		if r != ripv4 {
			t.Error(r, "should be", ripv4)
		}
	}

}

func TestIpAddrIpv6(t *testing.T) {
	var r string
	var addr string

	var ripv6 string = "IPv6"

	var ipv6 []string = []string{
		"2001:0db8:85a3:0000:0000:8a2e:0370:7334",
		"2001:db8:85a3:0:0:8A2E:0370:7334",
	}

	for _, addr = range ipv6 {
		r = validIPAddress(addr)
		if r != ripv6 {
			t.Error(r, "should be", ripv6)
		}
	}
}

func TestIpAddrNeither(t *testing.T) {
	var r string
	var addr string

	var invalid []string = []string{
		"172.16.254.01",
		"072.16.254.1",
		"172.016.54.1",
		"172.16.054.1",
		"02001:0db8:85a3:0000:0000:8a2e:0370:7334",
		"172.168.1.256",
	}

	for _, addr = range invalid {
		r = validIPAddress(addr)
		if r != "Neither" {
			t.Error(r, "should be", "Neither")
		}
	}
}
