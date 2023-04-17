package units

import (
	"fmt"
)

const (
	B Bytes = 1 << (10 * iota)
	KiB
	MiB
	GiB
	TiB
)

const (
	KB = 1000 * B
	MB = 1000 * KB
	GB = 1000 * MB
	TB = 1000 * GB
)

var (
	// key indicates whether '#' is represented in format flags
	magnitudes = map[bool][]Bytes{
		true:  decimalMagnitudes,
		false: binaryMagnitudes,
	}
	binaryMagnitudes  = []Bytes{B, KiB, MiB, GiB, TiB}
	decimalMagnitudes = []Bytes{B, KB, MB, GB, TB}
	unitNames         = map[Bytes][]byte{
		B:   []byte("B"),
		KiB: []byte("kiB"),
		MiB: []byte("MiB"),
		GiB: []byte("GiB"),
		TiB: []byte("TiB"),
		KB:  []byte("kB"),
		MB:  []byte("MB"),
		GB:  []byte("GB"),
		TB:  []byte("TB"),
	}
)

type Bytes uint64

func (b Bytes) Format(f fmt.State, verb rune) {
	width := 0
	if w, ok := f.Width(); ok {
		width = w
	}
	precision := 1
	if p, ok := f.Precision(); ok {
		precision = p
	}
	mag := B
	switch verb {
	case 's', 'v', 'f':
		if f.Flag('#') {
			mag = b.decimalMagnitude()
		} else {
			mag = b.magnitude()
		}
	case 'k':
		mag = magnitudes[f.Flag('#')][1]
	case 'm':
		mag = magnitudes[f.Flag('#')][2]
	case 'g':
		mag = magnitudes[f.Flag('#')][3]
	case 't':
		mag = magnitudes[f.Flag('#')][4]
	case 'b':
	case 'd':
		if width > 1 {
			fmt.Fprintf(f, "%*d", width, uint64(b))
		} else {
			fmt.Fprintf(f, "%d", uint64(b))
		}
		return
	}

	if _, ok := f.Width(); ok && !f.Flag(' ') {
		width -= len(unitNames[mag])
	}

	if verb == 'b' {
		if width > 1 {
			fmt.Fprintf(f, "%*d", width, uint64(b))
		} else {
			fmt.Fprintf(f, "%d", uint64(b))
		}
	} else if verb == 'f' {
		b.formatFloat(f, mag, width, precision)
	} else {
		b.formatInteger(f, mag, width)
	}
	if !f.Flag(' ') {
		f.Write(unitNames[mag])
	}
}

func (b Bytes) formatFloat(f fmt.State, mag Bytes, width, precision int) {
	if width > 1 {
		fmt.Fprintf(f, "%*.*f", width, precision, float64(b)/float64(mag))
	} else {
		fmt.Fprintf(f, "%.*f", precision, float64(b)/float64(mag))
	}
}

func (b Bytes) formatInteger(f fmt.State, mag Bytes, width int) {
	if width > 1 {
		fmt.Fprintf(f, "%*d", width, b/mag)
	} else {
		fmt.Fprintf(f, "%d", b/mag)
	}
}

func (b Bytes) magnitude() Bytes {
	switch {
	case b < KiB:
		return B
	case b < MiB:
		return KiB
	case b < GiB:
		return MiB
	case b < TiB:
		return GiB
	default:
		return TiB
	}
}

func (b Bytes) decimalMagnitude() Bytes {
	switch {
	case b < KB:
		return B
	case b < MB:
		return KB
	case b < GB:
		return MB
	case b < TB:
		return GB
	default:
		return TB
	}
}

func (b Bytes) Ceil() Bytes {
	mag := b.magnitude()
	return (b + mag - 1) & ^(mag - 1)
}

func (b Bytes) DecimalCeil() Bytes {
	mag := b.decimalMagnitude()
	switch mod := b % mag; mod {
	case 0:
		return b
	default:
		return b + mag - mod
	}
}

func (b Bytes) Floor() Bytes {
	return b & ^(b.magnitude() - 1)
}

func (b Bytes) DecimalFloor() Bytes {
	return b - b%b.decimalMagnitude()
}

func (b Bytes) Truncate(mag Bytes) Bytes {
	return b - b%mag
}

func (b Bytes) RoundBy(mag Bytes) Bytes {
	switch mag {
	case B:
		return b
	case KiB, MiB, GiB, TiB:
		return (b + mag>>1) & ^(mag - 1)
	default:
		return (b + mag/2).Truncate(mag)
	}
}

func (b Bytes) Round() Bytes {
	return b.RoundBy(b.magnitude())
}

func (b Bytes) DecimalRound() Bytes {
	return b.RoundBy(b.decimalMagnitude())
}
