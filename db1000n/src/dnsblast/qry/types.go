// Package qry [general utility functions for the dnsblast package]
package qry

import (
	"github.com/miekg/dns"
)

// ResponseCode convert numerical response code value to string
func ResponseCode(rc int) string {
	rcodes := map[int]string{
		0: "NOERROR",
		1: "FORMERR",
		2: "SERVFAIL",
		3: "NXDOMAIN",
		4: "NOTIMP",
		5: "REFUSED",
		6: "YXDOMAIN",
		7: "YXRRSET",
		8: "NOTAUTH",
		9: "NOTZONE",
	}
	return rcodes[rc]
}

// Qtype is used to conver string representation of query type into proper dns format
func Qtype(qt string) uint16 {
	switch qt {
	case "None":
		return dns.TypeNone
	case "A":
		return dns.TypeA
	case "NS":
		return dns.TypeNS
	case "MD":
		return dns.TypeMD
	case "MF":
		return dns.TypeMF
	case "CNAME":
		return dns.TypeCNAME
	case "SOA":
		return dns.TypeSOA
	case "MB":
		return dns.TypeMB
	case "MG":
		return dns.TypeMG
	case "MR":
		return dns.TypeMR
	case "NULL":
		return dns.TypeNULL
	case "PTR":
		return dns.TypePTR
	case "HINFO":
		return dns.TypeHINFO
	case "MINFO":
		return dns.TypeMINFO
	case "MX":
		return dns.TypeMX
	case "TXT":
		return dns.TypeTXT
	case "RP":
		return dns.TypeRP
	case "AFSDB":
		return dns.TypeAFSDB
	case "X25":
		return dns.TypeX25
	case "ISDN":
		return dns.TypeISDN
	case "RT":
		return dns.TypeRT
	case "NSAPPTR":
		return dns.TypeNSAPPTR
	case "SIG":
		return dns.TypeSIG
	case "KEY":
		return dns.TypeKEY
	case "PX":
		return dns.TypePX
	case "GPOS":
		return dns.TypeGPOS
	case "AAAA":
		return dns.TypeAAAA
	case "LOC":
		return dns.TypeLOC
	case "NXT":
		return dns.TypeNXT
	case "EID":
		return dns.TypeEID
	case "NIMLOC":
		return dns.TypeNIMLOC
	case "SRV":
		return dns.TypeSRV
	case "ATMA":
		return dns.TypeATMA
	case "NAPTR":
		return dns.TypeNAPTR
	case "KX":
		return dns.TypeKX
	case "CERT":
		return dns.TypeCERT
	case "DNAME":
		return dns.TypeDNAME
	case "OPT":
		return dns.TypeOPT
	case "DS":
		return dns.TypeDS
	case "SSHFP":
		return dns.TypeSSHFP
	case "RRSIG":
		return dns.TypeRRSIG
	case "NSEC":
		return dns.TypeNSEC
	case "DNSKEY":
		return dns.TypeDNSKEY
	case "DHCID":
		return dns.TypeDHCID
	case "NSEC3":
		return dns.TypeNSEC3
	case "NSEC3PARAM":
		return dns.TypeNSEC3PARAM
	case "TLSA":
		return dns.TypeTLSA
	case "SMIMEA":
		return dns.TypeSMIMEA
	case "HIP":
		return dns.TypeHIP
	case "NINFO":
		return dns.TypeNINFO
	case "RKEY":
		return dns.TypeRKEY
	case "TALINK":
		return dns.TypeTALINK
	case "CDS":
		return dns.TypeCDS
	case "CDNSKEY":
		return dns.TypeCDNSKEY
	case "OPENPGPKEY":
		return dns.TypeOPENPGPKEY
	case "CSYNC":
		return dns.TypeCSYNC
	case "SPF":
		return dns.TypeSPF
	case "UINFO":
		return dns.TypeUINFO
	case "UID":
		return dns.TypeUID
	case "GID":
		return dns.TypeGID
	case "UNSPEC":
		return dns.TypeUNSPEC
	case "NID":
		return dns.TypeNID
	case "L32":
		return dns.TypeL32
	case "L64":
		return dns.TypeL64
	case "LP":
		return dns.TypeLP
	case "EUI48":
		return dns.TypeEUI48
	case "EUI64":
		return dns.TypeEUI64
	case "URI":
		return dns.TypeURI
	case "CAA":
		return dns.TypeCAA
	case "AVC":
		return dns.TypeAVC
	default:
		return dns.TypeA
	}
}
