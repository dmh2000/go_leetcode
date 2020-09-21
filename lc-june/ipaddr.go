package lcjune

import (
	"regexp"
	"strings"
)

func validIPv4(ip string) bool {
	var p1 string = `^(\d{1,3})\.(\d{1,3})\.(\d{1,3})\.(\d{1,3})$`
	var a1 [][]string
	var s1 []string
	var r1 *regexp.Regexp

	r1 = regexp.MustCompile(p1)
	a1 = r1.FindAllStringSubmatch(ip, 4)

	// must match
	if a1 == nil {
		return false
	}

	// get first set of matches
	s1 = a1[0]

	// must have 5 groups (original + 4 digit groups)
	if len(s1) < 5 {
		return false
	}

	// test groups for invalid leading 0's
	for i := 1; i < len(s1); i++ {
		if s1[i][0] == '0' && len(s1[i]) > 1 {
			return false
		}
	}

	// check for '255
	for i := 1; i < len(s1); i++ {
		if s1[i] == "256" {
			return false
		}
	}

	return true
}

func validIPv6(ip string) bool {
	var p1 string = `^([a-f0-9]{1,4})\:` +
		`([a-f0-9]{1,4})\:` +
		`([a-f0-9]{1,4})\:` +
		`([a-f0-9]{1,4})\:` +
		`([a-f0-9]{1,4})\:` +
		`([a-f0-9]{1,4})\:` +
		`([a-f0-9]{1,4})\:` +
		`([a-f0-9]{1,4})$`

	var a1 [][]string
	var s1 []string
	var r1 *regexp.Regexp

	r1 = regexp.MustCompile(p1)
	a1 = r1.FindAllStringSubmatch(ip, 4)

	// must match
	if a1 == nil {
		return false
	}

	// get first match set
	s1 = a1[0]

	// must have 9 matches
	if len(s1) < 9 {
		return false
	}

	return true
}

func validIPAddress(IP string) string {

	IP = strings.ToLower(IP)

	// test ipv4
	if validIPv4(IP) {
		return "IPv4"
	}

	if validIPv6(IP) {
		return "IPv6"
	}
	return "Neither"
}
