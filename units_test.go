package units

import (
	"fmt"
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBytes_Ceil(t *testing.T) {
	tests := []struct {
		name string
		b    Bytes
		want Bytes
	}{
		{
			name: "1",
			b:    1,
			want: 1 * B,
		},
		{
			name: "1kiB",
			b:    1 * KiB,
			want: 1 * KiB,
		},
		{
			name: "1kiB+1B",
			b:    1*KiB + 1,
			want: 2 * KiB,
		},
		{
			name: "2kiB-1B",
			b:    2*KiB - 1,
			want: 2 * KiB,
		},
		{
			name: "2kiB",
			b:    2 * KiB,
			want: 2 * KiB,
		},
		{
			name: "2kiB+1B",
			b:    2*KiB + 1*B,
			want: 3 * KiB,
		},
		{
			name: "1MiB",
			b:    1 * MiB,
			want: 1 * MiB,
		},
		{
			name: "1MiB+1B",
			b:    1*MiB + 1,
			want: 2 * MiB,
		},
		{
			name: "2MiB-1B",
			b:    2*MiB - 1,
			want: 2 * MiB,
		},
		{
			name: "2MiB",
			b:    2 * MiB,
			want: 2 * MiB,
		},
		{
			name: "2MiB+1B",
			b:    2*MiB + 1*B,
			want: 3 * MiB,
		},
		{
			name: "1GiB",
			b:    1 * GiB,
			want: 1 * GiB,
		},
		{
			name: "1GiB+1B",
			b:    1*GiB + 1,
			want: 2 * GiB,
		},
		{
			name: "2GiB-1B",
			b:    2*GiB - 1,
			want: 2 * GiB,
		},
		{
			name: "2GiB",
			b:    2 * GiB,
			want: 2 * GiB,
		},
		{
			name: "2GiB+1B",
			b:    2*GiB + 1*B,
			want: 3 * GiB,
		},
		{
			name: "1TiB",
			b:    1 * TiB,
			want: 1 * TiB,
		},
		{
			name: "1TiB+1B",
			b:    1*TiB + 1,
			want: 2 * TiB,
		},
		{
			name: "2TiB-1B",
			b:    2*TiB - 1,
			want: 2 * TiB,
		},
		{
			name: "2TiB",
			b:    2 * TiB,
			want: 2 * TiB,
		},
		{
			name: "2TiB+1B",
			b:    2*TiB + 1*B,
			want: 3 * TiB,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.b.Ceil(); got != tt.want {
				t.Errorf("Ceil() = %v, want %v", got, tt.want)
			}
		})
	}
}

func FuzzBytes_Ceil(f *testing.F) {
	f.Fuzz(func(t *testing.T, i uint64) {
		var mag float64 = 1
		switch {
		case i >= 1<<40:
			mag = 1 << 40
		case i >= 1<<30:
			mag = 1 << 30
		case i >= 1<<20:
			mag = 1 << 20
		case i >= 1<<10:
			mag = 1 << 10
		}
		expect := uint64(math.Ceil(float64(i)/mag) * mag)
		got := Bytes(i).Ceil()
		if got != Bytes(expect) {
			t.Errorf("ceil of %d expect %d, got %d", i, expect, got)
		}
	})
}

func TestBytes_DecimalCeil(t *testing.T) {
	tests := []struct {
		name string
		b    Bytes
		want Bytes
	}{
		{
			name: "1",
			b:    1,
			want: 1 * B,
		},
		{
			name: "1kB",
			b:    1 * KB,
			want: 1 * KB,
		},
		{
			name: "1kB+1B",
			b:    1*KB + 1,
			want: 2 * KB,
		},
		{
			name: "2kB-1B",
			b:    2*KB - 1,
			want: 2 * KB,
		},
		{
			name: "2kB",
			b:    2 * KB,
			want: 2 * KB,
		},
		{
			name: "2kB+1B",
			b:    2*KB + 1*B,
			want: 3 * KB,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.b.DecimalCeil(); got != tt.want {
				t.Errorf("DecimalCeil() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBytes_Floor(t *testing.T) {
	tests := []struct {
		name string
		b    Bytes
		want Bytes
	}{
		{
			name: "1",
			b:    1,
			want: 1 * B,
		},
		{
			name: "1kiB",
			b:    1 * KiB,
			want: 1 * KiB,
		},
		{
			name: "1kiB+1B",
			b:    1*KiB + 1,
			want: 1 * KiB,
		},
		{
			name: "2kiB+1B",
			b:    2*KiB + 1,
			want: 2 * KiB,
		},
		{
			name: "2kiB",
			b:    2 * KiB,
			want: 2 * KiB,
		},
		{
			name: "2kiB-1B",
			b:    2*KiB - 1*B,
			want: 1 * KiB,
		},
		{
			name: "1MiB",
			b:    1 * MiB,
			want: 1 * MiB,
		},
		{
			name: "1MiB+1B",
			b:    1*MiB + 1,
			want: 1 * MiB,
		},
		{
			name: "2MiB+1B",
			b:    2*MiB + 1,
			want: 2 * MiB,
		},
		{
			name: "2MiB",
			b:    2 * MiB,
			want: 2 * MiB,
		},
		{
			name: "2MiB-1B",
			b:    2*MiB - 1*B,
			want: 1 * MiB,
		},
		{
			name: "1GiB",
			b:    1 * GiB,
			want: 1 * GiB,
		},
		{
			name: "1GiB+1B",
			b:    1*GiB + 1,
			want: 1 * GiB,
		},
		{
			name: "2GiB+1B",
			b:    2*GiB + 1,
			want: 2 * GiB,
		},
		{
			name: "2GiB",
			b:    2 * GiB,
			want: 2 * GiB,
		},
		{
			name: "2GiB-1B",
			b:    2*GiB - 1*B,
			want: 1 * GiB,
		},
		{
			name: "1TiB",
			b:    1 * TiB,
			want: 1 * TiB,
		},
		{
			name: "1TiB+1B",
			b:    1*TiB + 1,
			want: 1 * TiB,
		},
		{
			name: "2TiB+1B",
			b:    2*TiB + 1,
			want: 2 * TiB,
		},
		{
			name: "2TiB",
			b:    2 * TiB,
			want: 2 * TiB,
		},
		{
			name: "2TiB-1B",
			b:    2*TiB - 1*B,
			want: 1 * TiB,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.b.Floor(); got != tt.want {
				t.Errorf("Floor() = %v, want %v", got, tt.want)
			}
		})
	}
}

func FuzzBytes_Floor(f *testing.F) {
	f.Fuzz(func(t *testing.T, i uint64) {
		var mag float64 = 1
		switch {
		case i >= 1<<40:
			mag = 1 << 40
		case i >= 1<<30:
			mag = 1 << 30
		case i >= 1<<20:
			mag = 1 << 20
		case i >= 1<<10:
			mag = 1 << 10
		}
		expect := uint64(math.Floor(float64(i)/mag) * mag)
		got := Bytes(i).Floor()
		if got != Bytes(expect) {
			t.Errorf("floor of %d expect %d, got %d", i, expect, got)
		}
	})
}

func TestBytes_DecimalFloor(t *testing.T) {
	tests := []struct {
		name string
		b    Bytes
		want Bytes
	}{
		{
			name: "1",
			b:    1,
			want: 1 * B,
		},
		{
			name: "1kB",
			b:    1 * KB,
			want: 1 * KB,
		},
		{
			name: "1kB+1B",
			b:    1*KB + 1,
			want: 1 * KB,
		},
		{
			name: "2kB-1B",
			b:    2*KB - 1,
			want: 1 * KB,
		},
		{
			name: "2kB",
			b:    2 * KB,
			want: 2 * KB,
		},
		{
			name: "2kB+1B",
			b:    2*KB + 1*B,
			want: 2 * KB,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.b.DecimalFloor(); got != tt.want {
				t.Errorf("DecimalFloor() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBytes_Truncate(t *testing.T) {
	tests := []struct {
		name string
		b    Bytes
		mag  Bytes
		want Bytes
	}{
		{
			name: "zero truncate by B",
			b:    0,
			mag:  B,
			want: 0,
		},
		{
			name: "zero truncate by kiB",
			b:    0,
			mag:  KiB,
			want: 0,
		},
		{
			name: "zero truncate by kB",
			b:    0,
			mag:  KB,
			want: 0,
		},
		{
			name: "kiB truncate by kiB",
			b:    1*KiB + 1*B,
			mag:  KiB,
			want: 1 * KiB,
		},
		{
			name: "MiB truncate by kiB",
			b:    1*MiB + 1*KiB + 1*B,
			mag:  KiB,
			want: 1*MiB + 1*KiB,
		},
		{
			name: "kiB truncate by MiB",
			b:    1*KiB + 1*B,
			mag:  MiB,
			want: 0,
		},
		{
			name: "kB truncate by kB",
			b:    1*KB + 1*B,
			mag:  KB,
			want: 1 * KB,
		},
		{
			name: "MB truncate by kB",
			b:    1*MB + 1*KB + 1*B,
			mag:  KB,
			want: 1*MB + 1*KB,
		},
		{
			name: "kB truncate by MB",
			b:    1*KB + 1*B,
			mag:  MB,
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.b.Truncate(tt.mag); got != tt.want {
				t.Errorf("Truncate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBytes_RoundBy(t *testing.T) {
	tests := []struct {
		name string
		b    Bytes
		mag  Bytes
		want Bytes
	}{
		{
			name: "round up by kiB",
			b:    512*TiB + 512*GiB + 512*MiB + 512*KiB + 512*B,
			mag:  KiB,
			want: 512*TiB + 512*GiB + 512*MiB + 512*KiB + 1*KiB,
		},
		{
			name: "round up by MiB",
			b:    512*TiB + 512*GiB + 512*MiB + 512*KiB + 512*B,
			mag:  MiB,
			want: 512*TiB + 512*GiB + 512*MiB + 1*MiB,
		},
		{
			name: "round up by GiB",
			b:    512*TiB + 512*GiB + 512*MiB + 512*KiB + 512*B,
			mag:  GiB,
			want: 512*TiB + 512*GiB + 1*GiB,
		},
		{
			name: "round up by TiB",
			b:    512*TiB + 512*GiB + 512*MiB + 512*KiB + 512*B,
			mag:  TiB,
			want: 512*TiB + 1*TiB,
		},
		{
			name: "round down by kiB",
			b:    511*TiB + 511*GiB + 511*MiB + 511*KiB + 511*B,
			mag:  KiB,
			want: 511*TiB + 511*GiB + 511*MiB + 511*KiB,
		},
		{
			name: "round down by MiB",
			b:    511*TiB + 511*GiB + 511*MiB + 511*KiB + 511*B,
			mag:  MiB,
			want: 511*TiB + 511*GiB + 511*MiB,
		},
		{
			name: "round down by GiB",
			b:    511*TiB + 511*GiB + 511*MiB + 511*KiB + 511*B,
			mag:  GiB,
			want: 511*TiB + 511*GiB,
		},
		{
			name: "round down by TiB",
			b:    511*TiB + 511*GiB + 511*MiB + 511*KiB + 511*B,
			mag:  TiB,
			want: 511 * TiB,
		},
		{
			name: "round up by kB",
			b:    500*TB + 500*GB + 500*MB + 500*KB + 500*B,
			mag:  KB,
			want: 500*TB + 500*GB + 500*MB + 500*KB + 1*KB,
		},
		{
			name: "round up by MB",
			b:    500*TB + 500*GB + 500*MB + 500*KB + 500*B,
			mag:  MB,
			want: 500*TB + 500*GB + 500*MB + 1*MB,
		},
		{
			name: "round up by GB",
			b:    500*TB + 500*GB + 500*MB + 500*KB + 500*B,
			mag:  GB,
			want: 500*TB + 500*GB + 1*GB,
		},
		{
			name: "round up by TB",
			b:    500*TB + 500*GB + 500*MB + 500*KB + 500*B,
			mag:  TB,
			want: 500*TB + 1*TB,
		},
		{
			name: "round down by kB",
			b:    499*TB + 499*GB + 499*MB + 499*KB + 499*B,
			mag:  KB,
			want: 499*TB + 499*GB + 499*MB + 499*KB,
		},
		{
			name: "round down by MB",
			b:    499*TB + 499*GB + 499*MB + 499*KB + 499*B,
			mag:  MB,
			want: 499*TB + 499*GB + 499*MB,
		},
		{
			name: "round down by GB",
			b:    499*TB + 499*GB + 499*MB + 499*KB + 499*B,
			mag:  GB,
			want: 499*TB + 499*GB,
		},
		{
			name: "round down by TB",
			b:    499*TB + 499*GB + 499*MB + 499*KB + 499*B,
			mag:  TB,
			want: 499 * TB,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.b.RoundBy(tt.mag); got != tt.want {
				t.Errorf("RoundBy() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBytes_Format(t *testing.T) {
	formatToBinaryMagnitude := map[string]Bytes{"b": B, "k": KiB, "m": MiB, "g": GiB, "t": TiB}
	formatToDecimalMagnitude := map[string]Bytes{"b": B, "k": KB, "m": MB, "g": GB, "t": TB}

	bytes := map[string]Bytes{
		"zero": 0,
	}
	for i := 1; i < len(binaryMagnitudes); i++ {
		unit := binaryMagnitudes[i]
		lowerUnit := binaryMagnitudes[i-1]
		bytes["1"+string(unitNames[unit])] = 1 * unit
		bytes["2"+string(unitNames[unit])] = 2 * unit
		bytes["2"+string(unitNames[unit])+"-1"+string(unitNames[lowerUnit])] = 2*unit - 1*lowerUnit
		bytes["2"+string(unitNames[unit])+"-512"+string(unitNames[lowerUnit])] = 2*unit - 512*lowerUnit
		bytes["2"+string(unitNames[unit])+"+1"+string(unitNames[lowerUnit])] = 2*unit + 1*lowerUnit
		bytes["2"+string(unitNames[unit])+"+512"+string(unitNames[lowerUnit])] = 2*unit + 512*lowerUnit
	}
	for i := 1; i < len(decimalMagnitudes); i++ {
		unit := decimalMagnitudes[i]
		lowerUnit := decimalMagnitudes[i-1]
		bytes["1"+string(unitNames[unit])] = 1 * unit
		bytes["2"+string(unitNames[unit])] = 2 * unit
		bytes["2"+string(unitNames[unit])+"-1"+string(unitNames[lowerUnit])] = 2*unit - 1*lowerUnit
		bytes["2"+string(unitNames[unit])+"-500"+string(unitNames[lowerUnit])] = 2*unit - 500*lowerUnit
		bytes["2"+string(unitNames[unit])+"+1"+string(unitNames[lowerUnit])] = 2*unit + 1*lowerUnit
		bytes["2"+string(unitNames[unit])+"+500"+string(unitNames[lowerUnit])] = 2*unit + 500*lowerUnit
	}
	verbs := []string{"d", "b", "k", "m", "g", "t", "s", "v", "f"}
	for name, b := range bytes {
		for _, verb := range verbs {
			testName := name + " with " + verb
			logTpl := testName + ":  %-8s => %s"
			t.Run(testName, func(t *testing.T) {
				var format, result string
				doFormat := func() { result = fmt.Sprintf(format, b) }
				doLog := func() { t.Logf(logTpl, format, result) }
				if verb == "d" {
					format = "%d"
					doFormat()
					doLog()
					assert.Equal(t, fmt.Sprintf("%d", uint64(b)), result)

					format = "%6d"
					doFormat()
					doLog()
					assert.Equal(t, fmt.Sprintf("%6d", uint64(b)), result)
					return
				}

				var binMag, decMag Bytes
				if verb == "f" || verb == "s" || verb == "v" {
					binMag = b.magnitude()
					decMag = b.decimalMagnitude()
				} else {
					binMag = formatToBinaryMagnitude[verb]
					decMag = formatToDecimalMagnitude[verb]
				}

				if verb == "f" {
					binVal := float64(b) / float64(binMag)
					decVal := float64(b) / float64(decMag)

					for _, testcase := range []struct {
						format string
						expect string
					}{
						{
							format: "%f",
							expect: fmt.Sprintf("%.1f"+string(unitNames[binMag]), binVal),
						},
						{
							format: "%.2f",
							expect: fmt.Sprintf("%.2f"+string(unitNames[binMag]), binVal),
						},
						{
							format: "%10.2f",
							expect: fmt.Sprintf("%*.2f"+string(unitNames[binMag]), 10-len(string(unitNames[binMag])), binVal),
						},
						{
							format: "% f",
							expect: fmt.Sprintf("%.1f", binVal),
						},
						{
							format: "% .2f",
							expect: fmt.Sprintf("%.2f", binVal),
						},
						{
							format: "% 8.2f",
							expect: fmt.Sprintf("%8.2f", binVal),
						},
						{
							format: "%#f",
							expect: fmt.Sprintf("%.1f"+string(unitNames[decMag]), decVal),
						},
						{
							format: "%#.2f",
							expect: fmt.Sprintf("%.2f"+string(unitNames[decMag]), decVal),
						},
						{
							format: "%#10.2f",
							expect: fmt.Sprintf("%*.2f"+string(unitNames[decMag]), 10-len(string(unitNames[decMag])), decVal),
						},
						{
							format: "% #f",
							expect: fmt.Sprintf("%.1f", decVal),
						},
						{
							format: "% #.2f",
							expect: fmt.Sprintf("%.2f", decVal),
						},
						{
							format: "% #8.2f",
							expect: fmt.Sprintf("%8.2f", decVal),
						},
					} {
						format = testcase.format
						doFormat()
						doLog()
						assert.Equal(t, testcase.expect, result)
					}
					return
				}
				binVal := uint64(b / binMag)
				decVal := uint64(b / decMag)

				for _, testcase := range []struct {
					format string
					expect string
				}{
					{
						format: "%" + verb,
						expect: fmt.Sprintf("%d"+string(unitNames[binMag]), binVal),
					},
					{
						format: "%6" + verb,
						expect: fmt.Sprintf("%*d"+string(unitNames[binMag]), 6-len(string(unitNames[binMag])), binVal),
					},
					{
						format: "% " + verb,
						expect: fmt.Sprintf("%d", binVal),
					},
					{
						format: "% 6" + verb,
						expect: fmt.Sprintf("%6d", binVal),
					},
					{
						format: "%#" + verb,
						expect: fmt.Sprintf("%d"+string(unitNames[decMag]), decVal),
					},
					{
						format: "%#6" + verb,
						expect: fmt.Sprintf("%*d"+string(unitNames[decMag]), 6-len(string(unitNames[decMag])), decVal),
					},
					{
						format: "% #" + verb,
						expect: fmt.Sprintf("%d", decVal),
					},
					{
						format: "% #6" + verb,
						expect: fmt.Sprintf("%6d", decVal),
					},
				} {
					format = testcase.format
					doFormat()
					doLog()
					assert.Equal(t, testcase.expect, result)
				}
			})
		}
	}
}
