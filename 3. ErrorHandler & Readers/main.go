package main

import (
	"fmt"
	"io"
	"strings"
)

type ErrNegativeSqrt float64

func Sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, ErrNegativeSqrt(x)
	}
	return 0, nil
}

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("Cannot SQRT negative number [%v]", float64(e))
}
func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
	fmt.Println("==========================================")
	Readers()
}

func Readers() {
	r := strings.NewReader("Hello, Reader!")

	b := make([]byte, 8)
	for {
		n, err := r.Read(b)
		// fmt.Printf("n = %v err = %v b = %v\n", n, err, b)
		// fmt.Printf("b[:n] = %q\n", b[:n])

		//read each 8 byte
		fmt.Printf("%q", b[:n])
		if err == io.EOF {
			break
		}
	}

	//readfull
	if _, err := io.ReadFull(r, b); err != nil {
		fmt.Println(err)
	}

	fmt.Println(b)

}
