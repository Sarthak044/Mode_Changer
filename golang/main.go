package main

import (
	"fmt"
	"os/exec"
)

var input string

func main() {

	fmt.Println("Welcome to the WIRELESS ADAPTER MODE CHANGER")
	fmt.Println("Make sure your WIRELESS ADAPTER SUPPORTS MODES")
	fmt.Println("[+] Enter the interface")
	fmt.Scanln(&input)
	fmt.Println("ifconfig " + input + " down")

	if input == "eth0" || input == "lo" {
		fmt.Println("[+]Not a Wireless adapter")
		fmt.Println("[+]BYE BYE !")

	} else {
		fmt.Println("Select the Mode\n1.Monitor Mode\n2.Managed Mode")
		var Mode_number int8
		fmt.Scanln(&Mode_number)
		if Mode_number == 1 {
			process1()
		} else if Mode_number == 2 {
			process2()
		} else {
			fmt.Println("[-] Enter a valid Option", "[+]Exiting")
		}
	}

}

func process1() {

	fmt.Println("[+]Changing the mode to Monitor")
	c1, _ := exec.Command("sudo", "ifconfig", input, "down").Output()
	c2, _ := exec.Command("sudo", "iwconfig", input, "mode", "monitor").Output()
	c3, _ := exec.Command("sudo", "ifconfig", input, "up").Output()
	out1 := string(c1[:])
	fmt.Println(">", out1)
	out2 := string(c2[:])
	fmt.Println(">>", out2)
	out3 := string(c3[:])
	fmt.Println(">>>", out3)
	fmt.Println("[+] Mode changed Successfully")

}

func process2() {

	fmt.Println("[+]Changing the Mode to Managed")
	c1, _ := exec.Command("sudo", "ifconfig", input, "down").Output()
	c2, _ := exec.Command("sudo", "iwconfig", input, "mode", "managed").Output()
	c3, _ := exec.Command("sudo", "ifconfig", input, "up").Output()
	out1 := string(c1[:])
	fmt.Println(">", out1)
	out2 := string(c2[:])
	fmt.Println(">>", out2)
	out3 := string(c3[:])
	fmt.Println(">>>", out3)
	fmt.Println("[+] Mode changed Successfully")
}
