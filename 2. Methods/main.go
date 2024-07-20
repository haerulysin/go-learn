package main

import (
	"fmt"
)

type MyNumber struct {
	X []int
}

func (v *MyNumber) getEven() {

	for _, num := range v.X {

		if num%2 == 0 {
			fmt.Println(num)
		}
	}
}
func RemoveIndex(s *[]int, index int) {
	slc := *s
	*s = append(slc[:index], slc[index+1:]...)
}
func (v *MyNumber) setEven() {
	for i, num := range v.X {
		if num%2 != 0 {
			RemoveIndex(&v.X, i)
		}
		continue
	}
}

func main() {
	////Methods & pointer indirection to get even
	// nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	// MN := &MyNumber{nums}
	// MN.setEven()
	// fmt.Println(MN.X)
	// MN.getEven()

	var helloText interface{} = "Hello"
	_, isString := helloText.(string)
	fmt.Println("Is String?", isString)
	if _, isFloat := helloText.(float64); isFloat {
		fmt.Println(helloText, " is Float")
	} else {
		fmt.Println(helloText, "is Not Float")
	}

	// switch dataType := helloText.(type) {
	// case int:
	// 	fmt.Println("[Switch Case]", dataType, "is Integer")
	// case string:
	// 	fmt.Println("[Switch Case]", dataType, "is String")
	// default:
	// 	fmt.Println("[Switch Case]", dataType, "is Unknown Type")
	// }
	// ExerciseStringers()

}

type IPAddr [4]byte

// TODO: Add a "String() string" method to IPAddr.
// Make the IPAddr type implement fmt.Stringer to print the address as a dotted quad.

// For instance, IPAddr{1, 2, 3, 4} should print as "1.2.3.4".
func (addr IPAddr) String() string {
	return fmt.Sprintf("%v.%d.%d.%d", addr[0], addr[1], addr[2], addr[3])
}
func ExerciseStringers() {

	hosts := map[string]IPAddr{
		"loopback":  {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}
	// fmt.Printf("%b\n", 2) // 10
	// fmt.Printf("%c\n", 1400) // n
	// fmt.Printf("%e \n", 172.111) //1.721110e+02
	// fmt.Printf("%.2f\n", 172.123333) //172.12

	for name, ip := range hosts {
		fmt.Printf("%v: %v\n", name, ip)
	}
}
