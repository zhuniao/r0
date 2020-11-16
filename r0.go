package r0

import "fmt"

func SayHello(lang string) {
	if "en" == lang {
		fmt.Println("Hi,I'm r0!")
	} else {
		fmt.Println("sorry,unsupport language.")
	}
}
