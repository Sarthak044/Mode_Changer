package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
)

var input string

func main() {
	fmt.Println("Welcome to the WIRELESS ADAPTER MODE CHANGER")
	fmt.Println("Make sure your WIRELESS ADAPTER SUPPORTS MODES")
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("[+] Enter the interface")
	input, _ := reader.ReadString('\n')

	if input == "eth0" && input == "lo" {

		fmt.Println("[+]Not a Wireless adapter")
		fmt.Println("[+]BYE BYE !")
	} else {
		fmt.Println("Select the Mode\n1.Monitor Mode\n2.Managed Mode")
		Mode_number, _ := reader.ReadString('\n')
		if Mode_number == "1" {
			process1()
		} else if Mode_number == "2" {
			process2()
		} else {
			fmt.Println("[-] Enter a valid Option", "[+]Exiting")
		}
	}
}

func process1() {
	fmt.Println("[+]Changing the mode to Monitor")
	exec.Command("ifconfig" + input + "down")
	exec.Command("iwconfig" + input + "mode" + "monitor")
	exec.Command("ifconfig" + input + "up")
	fmt.Println("[+] Mode changed Successfully")
}

func process2() {
	fmt.Println("[+]Changing the Mode to Managed")
	exec.Command("ifconfig" + input + "down")
	exec.Command("iwconfig" + input + "mode" + "managed")
	exec.Command("ifconfig" + input + "up")
	fmt.Println("[+] Mode changed Successfully")
}
