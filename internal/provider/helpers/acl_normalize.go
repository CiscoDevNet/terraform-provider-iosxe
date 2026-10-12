// Copyright © 2023 Cisco Systems, Inc. and its affiliates.
// All rights reserved.
//
// Licensed under the Mozilla Public License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.

package helpers

// NormalizePort converts IOS-XE symbolic port names returned by NETCONF
// get-config to their numeric equivalents. The provider writes numeric
// values via edit-config, but the device returns symbolic names in
// responses, causing idempotency drift.
func NormalizePort(value string) string {
	if v, ok := portSymbolicToNumeric[value]; ok {
		return v
	}
	return value
}

// NormalizeDscp converts IOS-XE symbolic DSCP names returned by NETCONF
// get-config to their numeric equivalents.
func NormalizeDscp(value string) string {
	if v, ok := dscpSymbolicToNumeric[value]; ok {
		return v
	}
	return value
}

var portSymbolicToNumeric = map[string]string{
	"echo":           "7",
	"discard":        "9",
	"daytime":        "13",
	"chargen":        "19",
	"ftp-data":       "20",
	"ftp":            "21",
	"telnet":         "23",
	"smtp":           "25",
	"time":           "37",
	"whois":          "43",
	"tacacs":         "49",
	"domain":         "53",
	"bootps":         "67",
	"bootpc":         "68",
	"tftp":           "69",
	"gopher":         "70",
	"finger":         "79",
	"www":            "80",
	"hostname":       "101",
	"pop2":           "109",
	"pop3":           "110",
	"sunrpc":         "111",
	"ident":          "113",
	"nntp":           "119",
	"ntp":            "123",
	"netbios-ns":     "137",
	"netbios-dgm":    "138",
	"netbios-ss":     "139",
	"imap4":          "143",
	"snmp":           "161",
	"snmptrap":       "162",
	"bgp":            "179",
	"irc":            "194",
	"ldap":           "389",
	"https":          "443",
	"pim-auto-rp":    "496",
	"exec":           "512",
	"login":          "513",
	"cmd":            "514",
	"lpd":            "515",
	"talk":           "517",
	"rip":            "520",
	"uucp":           "540",
	"klogin":         "543",
	"kshell":         "544",
	"rtsp":           "554",
	"ldaps":          "636",
	"kerberos":       "750",
	"h323gatestat":   "1719",
	"h323callsigalt": "1720",
	"pptp":           "1723",
	"radius-acct":    "1813",
	"nfs":            "2049",
	"sip":            "5060",
	"sips":           "5061",
	"aol":            "5190",
}

var dscpSymbolicToNumeric = map[string]string{
	"default": "0",
	"cs1":     "8",
	"af11":    "10",
	"af12":    "12",
	"af13":    "14",
	"cs2":     "16",
	"af21":    "18",
	"af22":    "20",
	"af23":    "22",
	"cs3":     "24",
	"af31":    "26",
	"af32":    "28",
	"af33":    "30",
	"cs4":     "32",
	"af41":    "34",
	"af42":    "36",
	"af43":    "38",
	"cs5":     "40",
	"ef":      "46",
	"cs6":     "48",
	"cs7":     "56",
}
