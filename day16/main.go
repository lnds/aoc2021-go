package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"strconv"

	"github.com/lnds/aoc2021-go/shared"
)

func main() {
	lines := shared.ReadLines(os.Args[1])
	for _, line := range lines {
		fmt.Printf("packet hex = %s\n", line)
		packet := newPacket(hexToBin(line))
		fmt.Println("sum versions = ", sumVersions(packet))
		fmt.Println("value = ", packet.Eval())
		fmt.Println()
	}
}

var hexMap = map[rune]string{
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

func hexToBin(hex string) string {
	bin := ""
	for _, c := range hex {
		bin += hexMap[c]
	}
	return bin
}

func sumVersions(packet *Packet) int {
	if packet == nil {
		return 0
	}
	//fmt.Printf("packet header = %#v\n", packet.Header)
	sum := packet.version
	if packet.operator != nil {
		for _, p := range packet.operator.subPackets {
			sum += sumVersions(p)
		}
	}
	return sum
}

type Header struct {
	version int
	typeId  int
}

type ValuePacket struct {
	value int
}

type OperatorPacket struct {
	lengthTypeId int
	length       int
	subPackets   []*Packet
}

type Packet struct {
	Header
	operator *OperatorPacket
	value    *ValuePacket
}

const (
	SumType     = 0
	ProductType = 1
	MinimunType = 2
	MaximunType = 3
	ValueType   = 4
	GreaterType = 5
	LessType    = 6
	EqualType   = 7
)

func (p *Packet) Eval() int {
	if p == nil {
		log.Fatal("p nil on Eval")
	}
	if p.value != nil {
		return p.value.Eval()
	} else if p.operator != nil {
		return p.operator.Eval(p.Header)
	}
	log.Fatal("can't be value or operator")
	return 0
}

func (v *ValuePacket) Eval() int {
	return v.value
}

func (o *OperatorPacket) Eval(header Header) int {
	switch header.typeId {
	case SumType:
		return o.sum()
	case ProductType:
		return o.product()
	case MinimunType:
		return o.min()
	case MaximunType:
		return o.max()
	case GreaterType:
		return o.greater()
	case LessType:
		return o.less()
	case EqualType:
		return o.equal()
	default:
		return 0
	}
}

func (o *OperatorPacket) sum() int {
	sum := 0
	for _, p := range o.subPackets {
		sum += p.Eval()
	}
	return sum
}

func (o *OperatorPacket) product() int {
	res := 1
	for _, p := range o.subPackets {
		if p != nil {
			res *= p.Eval()
		}
	}
	fmt.Printf("product result = %d\n", res)
	return res
}

func (o *OperatorPacket) min() int {
	m := math.MaxInt
	for _, p := range o.subPackets {
		v := p.Eval()
		if v < m {
			m = v
		}
	}
	return m
}

func (o *OperatorPacket) max() int {
	m := 0
	for _, p := range o.subPackets {
		v := p.Eval()
		if v > m {
			m = v
		}
	}
	return m
}

func (o *OperatorPacket) greater() int {
	if o.subPackets[0].Eval() > o.subPackets[1].Eval() {
		return 1
	} else {
		return 0
	}
}

func (o *OperatorPacket) less() int {
	if o.subPackets[0].Eval() < o.subPackets[1].Eval() {
		return 1
	} else {
		return 0
	}
}

func (o *OperatorPacket) equal() int {
	if o.subPackets[0].Eval() == o.subPackets[1].Eval() {
		return 1
	} else {
		return 0
	}
}

func newPacket(binary string) *Packet {
	p, _ := parsePacket(binary)
	return p
}

func parsePacket(binary string) (*Packet, string) {
	//fmt.Printf("parse packet [%s]\n", binary)
	if len(binary) < 6 {
		return nil, ""
	}
	ver := parseBinary(binary[0:3])
	typ := parseBinary(binary[3:6])
	hdr := Header{ver, typ}

	if typ == ValueType {
		p, s := newValuePacket(binary[6:])
		packet := Packet{
			Header:   hdr,
			operator: nil,
			value:    p,
		}
		return &packet, s
	} else {
		p, s := newOperatorPacket(binary[6:])
		packet := Packet{
			Header:   hdr,
			operator: p,
			value:    nil,
		}
		return &packet, s
	}
}

func newValuePacket(binary string) (*ValuePacket, string) {
	binValue := ""
	var pos = 0
	exit := false
	for pos < len(binary) && !exit {
		exit = binary[pos] == '0'
		binValue += binary[pos+1 : pos+5]
		pos += 5
	}
	packet := ValuePacket{
		value: parseBinary(binValue),
	}
	return &packet, binary[pos:]
}

func newOperatorPacket(binary string) (*OperatorPacket, string) {
	if len(binary) == 0 {
		return nil, ""
	}
	if binary[0] == '0' && len(binary) < 16 {
		return nil, ""
	}
	if binary[0] == '1' && len(binary) < 12 {
		return nil, ""
	}
	i := binary[0]
	var l int
	var pos int
	var lengthTypeId int
	packets := []*Packet{}
	if i == '0' {
		lengthTypeId = 0
		l = parseBinary(binary[1:16])
		pos = 16
		//fmt.Println("OPERATOR")
		//fmt.Printf("binary = [%s]\n", binary)
		bin := binary[pos : pos+l]
		rest := binary[pos+l:]
		//fmt.Printf("l = %d, pos+l = %d, len(binary) = %d\n", l, pos+l, len(binary))
		//fmt.Printf("bin = [%s] len = %d\n", bin, len(bin))
		//fmt.Printf("rest = [%s] len = %d\n", rest, len(rest))
		for len(bin) > 0 {
			p, s := parsePacket(bin)
			if p != nil {
				packets = append(packets, p)
			}
			bin = s
		}
		binary = rest
	} else {
		lengthTypeId = 1
		l = parseBinary(binary[1:12])
		pos = 12
		bin := binary[pos:]
		for i := 0; i < l; i++ {
			p, s := parsePacket(bin)
			if p != nil {
				packets = append(packets, p)
			}
			bin = s
		}
		binary = bin
	}
	packet := OperatorPacket{
		lengthTypeId: lengthTypeId,
		length:       l,
		subPackets:   packets,
	}
	return &packet, binary
}

func parseBinary(binary string) int {
	v, err := strconv.ParseInt(binary, 2, 64)
	if err != nil {
		log.Fatal(err)
	}
	return int(v)
}
