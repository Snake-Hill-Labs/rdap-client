package protocol

import "time"

// IPAddresses describes the ipAddresses field as it is in RFC 7483, section
// 5.2
type IPAddresses struct {
	V4 []string `json:"v4,omitempty"`
	V6 []string `json:"v6,omitempty"`
}

// Nameserver describes the Nameserver Object Class as it is in RFC 7483,
// section 5.2
type Nameserver struct {
	ObjectClassName string       `json:"objectClassName"`
	LDHName         string       `json:"ldhName,omitempty"`
	UnicodeName     string       `json:"unicodeName,omitempty"`
	IPAddresses     *IPAddresses `json:"ipAddresses,omitempty"`
	HostStatus      string       `json:"nicbr_status,omitempty"`
	LastCheckAt     time.Time    `json:"nicbr_lastCheck,omitempty"`
	LastOKAt        time.Time    `json:"nicbr_lastOK,omitempty"`
	Remarks         []Remark     `json:"remarks,omitempty"`
}
