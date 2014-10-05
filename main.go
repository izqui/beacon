package beacon

import (
	"encoding/hex"
	"errors"
	"strings"

	"github.com/paypal/gatt"
)

type Beacon struct {
	Identifier []byte
	Major      uint16
	Minor      uint16
	Power      byte

	GattServer gatt.Server
}

func NewBeacon(uuid string, major, minor int) (*Beacon, error) {

	id, err := ParseUUID(uuid)
	if err != nil {
		return nil, err
	}

	return &Beacon{Identifier: id, Major: uint16(major), Minor: uint16(minor), Power: 0xC5}, nil
}

func (b *Beacon) StartAdvertising() error {

	server := &gatt.Server{Name: "iBeacon"}

	advertisingPacket := []byte{}
	advertisingPacket = append(advertisingPacket, 0x02)                                                  // Number of bytes that follow in first advertising structure
	advertisingPacket = append(advertisingPacket, 0x01)                                                  // Number of flags
	advertisingPacket = append(advertisingPacket, 0x1A)                                                  // Flag -> 0x1A = 000011010
	advertisingPacket = append(advertisingPacket, 0x1A)                                                  // Number of bytes that follow in second advertising structure
	advertisingPacket = append(advertisingPacket, 0xFF)                                                  // Manufacturer specific data advertising type
	advertisingPacket = append(advertisingPacket, []byte{0x4C, 0x00}...)                                 // Apple company identifier
	advertisingPacket = append(advertisingPacket, []byte{0x02, 0x15}...)                                 // iBeacon identifier
	advertisingPacket = append(advertisingPacket, b.Identifier...)                                       // iBeacon UUID
	advertisingPacket = append(advertisingPacket, []byte{uint8(b.Major >> 8), uint8(b.Major & 0xff)}...) // iBeacon major
	advertisingPacket = append(advertisingPacket, []byte{uint8(b.Minor >> 8), uint8(b.Minor & 0xff)}...) // iBeacon minor
	advertisingPacket = append(advertisingPacket, b.Power)

	server.AdvertisingPacket = advertisingPacket
	return server.AdvertiseAndServe()
}

//From https://github.com/paypal/gatt/blob/master/uuid.go
func ParseUUID(s string) ([]byte, error) {
	s = strings.Replace(s, "-", "", -1)
	b, err := hex.DecodeString(s)
	if err != nil {
		return nil, err
	}
	if len(b) != 16 {
		return nil, errors.New("UUID length must be 16 bytes")
	}
	return b, nil
}
