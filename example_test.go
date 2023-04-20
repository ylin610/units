package units_test

import (
	"fmt"

	"github.com/ylin610/units"
)

func ExampleBytes_Ceil() {
	b := units.KiB + units.B
	fmt.Printf("%.3f\n", b)
	fmt.Printf("%.3f\n", b.Ceil())

	b = units.MiB + units.B
	fmt.Printf("%.3f\n", b)
	fmt.Printf("%.3f\n", b.Ceil())

	// difference with 'DecimalCeil'
	fmt.Printf("%s\n", units.Bytes(2001).Ceil())
	fmt.Printf("%#s\n", units.Bytes(2001).DecimalCeil())

	// output:
	// 1.001kiB
	// 2.000kiB
	// 1.000MiB
	// 2.000MiB
	// 2kiB
	// 3kB
}

func ExampleBytes_DecimalCeil() {
	b := units.KB + units.B
	fmt.Printf("%#.3f\n", b)
	fmt.Printf("%#.3f\n", b.DecimalCeil())

	b = units.MB + units.B
	fmt.Printf("%#.3f\n", b)
	fmt.Printf("%#.3f\n", b.DecimalCeil())

	// difference with 'Ceil'
	fmt.Printf("%s\n", units.Bytes(2001).Ceil())
	fmt.Printf("%#s\n", units.Bytes(2001).DecimalCeil())

	// output:
	// 1.001kB
	// 2.000kB
	// 1.000MB
	// 2.000MB
	// 2kiB
	// 3kB
}

func ExampleBytes_Floor() {
	b := units.KiB + units.B
	fmt.Printf("%.3f\n", b)
	fmt.Printf("%.3f\n", b.Floor())

	// difference with 'DecimalFloor'
	fmt.Printf("%s\n", units.Bytes(2001).Floor())
	fmt.Printf("%#s\n", units.Bytes(2001).DecimalFloor())

	// output:
	// 1.001kiB
	// 1.000kiB
	// 1kiB
	// 2kB
}

func ExampleBytes_DecimalFloor() {
	b := units.KB + units.B
	fmt.Printf("%#.3f\n", b)
	fmt.Printf("%#.3f\n", b.DecimalFloor())

	// difference with 'DecimalFloor'
	fmt.Printf("%s\n", units.Bytes(2001).Floor())
	fmt.Printf("%#s\n", units.Bytes(2001).DecimalFloor())

	// output:
	// 1.001kB
	// 1.000kB
	// 1kiB
	// 2kB
}

func ExampleBytes_Truncate() {
	b := units.MiB + 512*units.KiB + 512*units.B
	fmt.Printf("%.4f\n", b)
	fmt.Printf("%.4f\n", b.Truncate(units.KiB))
	fmt.Printf("%.4f\n", b.Truncate(units.MiB))

	// output:
	// 1.5005MiB
	// 1.5000MiB
	// 1.0000MiB
}

func ExampleBytes_RoundBy() {
	b := units.MB + 500*units.KB + 500*units.B
	fmt.Printf("%#.4f\n", b)
	fmt.Printf("%#.4f\n", b.RoundBy(units.KB))
	fmt.Printf("%#.4f\n", b.RoundBy(units.MB))

	// output:
	// 1.5005MB
	// 1.5010MB
	// 2.0000MB
}

func ExampleBytes_Round() {
	b := units.KiB + 512*units.B
	fmt.Printf("%f\n", b)
	fmt.Printf("%f\n", b.Round())

	b -= units.B
	fmt.Printf("%f\n", b.Round())

	// output:
	// 1.5kiB
	// 2.0kiB
	// 1.0kiB
}

func ExampleBytes_DecimalRound() {
	b := units.KB + 500*units.B
	fmt.Printf("%#f\n", b)
	fmt.Printf("%#f\n", b.DecimalRound())

	b -= units.B
	fmt.Printf("%#f\n", b.DecimalRound())

	// output:
	// 1.5kB
	// 2.0kB
	// 1.0kB
}
