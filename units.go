package units

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

type Bytes uint64

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
