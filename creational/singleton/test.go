package singleton

import "fmt"

func TestSingleton() {
	// for i := 0; i < 30; i++ {
	// 	go getInstance()
	// }

	// fmt.Scanln()

	for i := 0; i < 30; i++ {
		go getInstance2()
	}

	fmt.Scanln()
}
