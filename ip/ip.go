package ip

import "fmt"

func IpHello(name string) string {
	message := fmt.Sprintf("Hi. %v !", name)
	return message
}
