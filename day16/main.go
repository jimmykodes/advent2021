package main

import (
	"fmt"
	"math"
	"strconv"
)

// lookupTable instead of doing strconv.FormatInt(strconv.ParseInt(x, 16, 64), 2)
// for each value to convert from hex to bin
var lookupTable = map[rune]string{
	'0': "0000",
	'1': "0001",
	'2': "0010",
	'3': "0011",
	'4': "0100",
	'5': "0101",
	'6': "0110",
	'7': "0111",
	'8': "1000",
	'9': "1001",
	'A': "1010",
	'B': "1011",
	'C': "1100",
	'D': "1101",
	'E': "1110",
	'F': "1111",
}

func main() {
	in := "6053231004C12DC26D00526BEE728D2C013AC7795ACA756F93B524D8000AAC8FF80B3A7A4016F6802D35C7C94C8AC97AD81D30024C00D1003C80AD050029C00E20240580853401E98C00D50038400D401518C00C7003880376300290023000060D800D09B9D03E7F546930052C016000422234208CC000854778CF0EA7C9C802ACE005FE4EBE1B99EA4C8A2A804D26730E25AA8B23CBDE7C855808057C9C87718DFEED9A008880391520BC280004260C44C8E460086802600087C548430A4401B8C91AE3749CF9CEFF0A8C0041498F180532A9728813A012261367931FF43E9040191F002A539D7A9CEBFCF7B3DE36CA56BC506005EE6393A0ACAA990030B3E29348734BC200D980390960BC723007614C618DC600D4268AD168C0268ED2CB72E09341040181D802B285937A739ACCEFFE9F4B6D30802DC94803D80292B5389DFEB2A440081CE0FCE951005AD800D04BF26B32FC9AFCF8D280592D65B9CE67DCEF20C530E13B7F67F8FB140D200E6673BA45C0086262FBB084F5BF381918017221E402474EF86280333100622FC37844200DC6A8950650005C8273133A300465A7AEC08B00103925392575007E63310592EA747830052801C99C9CB215397F3ACF97CFE41C802DBD004244C67B189E3BC4584E2013C1F91B0BCD60AA1690060360094F6A70B7FC7D34A52CBAE011CB6A17509F8DF61F3B4ED46A683E6BD258100667EA4B1A6211006AD367D600ACBD61FD10CBD61FD129003D9600B4608C931D54700AA6E2932D3CBB45399A49E66E641274AE4040039B8BD2C933137F95A4A76CFBAE122704026E700662200D4358530D4401F8AD0722DCEC3124E92B639CC5AF413300700010D8F30FE1B80021506A33C3F1007A314348DC0002EC4D9CF36280213938F648925BDE134803CB9BD6BF3BFD83C0149E859EA6614A8C"
	out := ""
	for _, c := range in {
		out += lookupTable[c]
	}
	p, _, err := parsePacket(out)
	if err != nil {
		panic(err)
	}
	fmt.Println(evalPacket(p))
}

type Packet struct {
	Version    int64
	Type       int64
	LengthType int64
	LengthVal  int64

	SubPackets []*Packet
	Value      int64
}

func evalPacket(packet *Packet) int64 {
	packetVals := make([]int64, len(packet.SubPackets))
	for i, subPacket := range packet.SubPackets {
		packetVals[i] = evalPacket(subPacket)
	}
	var v int64
	switch packet.Type {
	case 0:
		for _, val := range packetVals {
			v += val
		}
	case 1:
		v = 1
		for _, val := range packetVals {
			v *= val
		}
	case 2:
		v = math.MaxInt64
		for _, val := range packetVals {
			if val < v {
				v = val
			}
		}
	case 3:
		for _, val := range packetVals {
			if val > v {
				v = val
			}
		}
	case 4:
		v = packet.Value
	case 5:
		if packetVals[0] > packetVals[1] {
			v = 1
		}
	case 6:
		if packetVals[0] < packetVals[1] {
			v = 1
		}
	case 7:
		if packetVals[0] == packetVals[1] {
			v = 1
		}
	}
	return v
}

func parsePacket(packet string) (*Packet, int, error) {
	p := &Packet{}
	var (
		cursor int
		err    error
	)
	p.Version, err = strconv.ParseInt(packet[cursor:cursor+3], 2, 64)
	if err != nil {
		return nil, 0, err
	}
	cursor += 3
	p.Type, err = strconv.ParseInt(packet[cursor:cursor+3], 2, 64)
	if err != nil {
		return nil, 0, err
	}
	cursor += 3
	if p.Type == 4 {
		// this is a literal value, parse the bits
		for {
			chunk := packet[cursor : cursor+5]
			cursor += 5

			chunkVal, err := strconv.ParseInt(chunk, 2, 64)
			if err != nil {
				return nil, 0, err
			}
			p.Value = p.Value << 4
			p.Value |= chunkVal & (1<<4 - 1)

			// check msb of chunkVal, of 0, break
			if chunkVal&(1<<4) == 0 {
				break
			}
		}
	} else {
		// operator packet
		p.LengthType, err = strconv.ParseInt(string(packet[cursor]), 2, 64)
		if err != nil {
			return nil, 0, err
		}
		cursor++
		if p.LengthType == 1 {
			// 11 bits indicating number of sub packets
			p.LengthVal, err = strconv.ParseInt(packet[cursor:cursor+11], 2, 64)
			if err != nil {
				return nil, 0, err
			}
			cursor += 11
			for i := 0; i < int(p.LengthVal); i++ {
				subP, bitsRead, err := parsePacket(packet[cursor:])
				if err != nil {
					return nil, 0, err
				}
				cursor += bitsRead
				p.SubPackets = append(p.SubPackets, subP)
			}
		} else {
			// 15 bits indicating number of bits in sub packets
			p.LengthVal, err = strconv.ParseInt(packet[cursor:cursor+15], 2, 64)
			if err != nil {
				return nil, 0, err
			}
			cursor += 15
			var totalBitsRead int64
			for totalBitsRead < p.LengthVal {
				subP, br, err := parsePacket(packet[cursor:])
				if err != nil {
					return nil, 0, err
				}
				totalBitsRead += int64(br)
				cursor += br
				p.SubPackets = append(p.SubPackets, subP)
			}
		}
	}
	return p, cursor, nil
}
