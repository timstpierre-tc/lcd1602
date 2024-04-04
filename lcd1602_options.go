/*
Copyright 2024 Tim St. Pierre
Options for lcd1602 character display
*/
package lcd1602

import (
	"errors"
	"time"
)

type Opts struct {
	// The I²C slave address
	I2CAddr uint16
	// How many lines does the display have
	Lines     uint8
	Cols      uint8
	CharDelay time.Duration
}

var DefaultOpts = Opts{
	I2CAddr:   0x27,
	Lines:     2,
	Cols:      16,
	CharDelay: 1 * time.Millisecond,
}

func (o *Opts) i2cAddr() (uint16, error) {
	switch o.I2CAddr {
	case 0:
		// Default address.
		return 0x27, nil
	case 0x20, 0x21, 0x22, 0x23, 0x24, 0x25, 0x26, 0x27:
		return o.I2CAddr, nil
	default:
		return 0, errors.New("given address not supported by device")
	}
}
