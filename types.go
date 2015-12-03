package main

var ELEMENT_IDS = map[byte]string{
	0:   "SSID",
	1:   "SUPPORTED_RATES",
	50:  "EXTENDED_SUPPORTED_RATES",
	3:   "DS_PARAMETER_SET",
	45:  "HT_CAPABILITIES",
	127: "EXTENDED_CAPABILITIES",
	221: "VENDOR_SPECIFIC",
	//107: "??",
	//191: "??",
}

type Wireless80211Frame struct {
	Length           uint16
	TSFT             uint64
	FlagsRadio       uint8
	DBMAntennaSignal int8
	Type             string
	Flags80211       uint8
	Proto            uint8
	DurationID       uint16
	Address1         string
	Address2         string
	Address3         string
	Address4         string
	SequenceNumber   uint16
	FragmentNumber   uint16
	Checksum         uint32
	Elements         map[string][]byte
}