package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strings"
)

func main() {

	require_packages()
	loader()
}

func loader() {
	menu("| 1  - Install informant \n| 2  - Unistall \n| 0  - Exit")
	input := bufio.NewReader(os.Stdin)
	fmt.Print("Enter your choice: ")
	choice, _ := input.ReadString('\n')
	choice = strings.TrimSpace(choice)
	switch choice {
	case "1":
		install()
	case "2":
		fmt.Println("unistall....")
	case "0":
		fmt.Printf("Exiting program...")
	default:
		fmt.Println("Invalid choice. Please enter a valid option.")
	}
}

func require_packages() {
	installPackage("jq")
}

func install() {
	// // utils.hello()
	// utils.PrintWelcomeMessage()
}

func clearTerminal() {
	switch runtime.GOOS {
	case "linux", "darwin":
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
	case "windows":
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
	default:
		fmt.Println("Unsupported platform")
	}
}

func installPackage(packageName string) error {
	cmd := exec.Command("sudo", "apt", "install", "-y", packageName)
	output, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("failed to install package: %s, output: %s", err, string(output))
	}
	fmt.Printf("Package installed successfully: %s\n", string(output))
	return nil
}

func menu(items string) {
	clearTerminal()
	green := "\033[32m"
	red := "\033[31m"
	yellow := "\033[33m"
	reset := "\033[0m"

	serverCountry := "Germany"
	serverIP := "192.168.1.1"
	serverISP := "Deutsche Telekom"
	version := "0.1"

	fmt.Println("+----------------------------------------------------------------------------------------------------------------+")
	fmt.Println("| ##  ##   ##   ##  ######       ######   ##   ##  #######    ###    ######   ##   ##    ##     ##   ##  ######  |")
	fmt.Println("| ##  ##   ##   ##    ##           ##     ###  ##  ##        ## ##   ##   ##  ##   ##    ##     ###  ##    ##    |")
	fmt.Println("|  ####    ##   ##    ##           ##     ###  ##  ##       ##   ##  ##   ##  ### ###   ####    ###  ##    ##    |")
	fmt.Println("|   ##     ##   ##    ##           ##     ## # ##  #####    ##   ##  ######   ## # ##   ## #    ## # ##    ##    |")
	fmt.Println("|  ####    ##   ##    ##           ##     ## # ##  ##       ##   ##  ## ##    ## # ##  ######   ## # ##    ##    |")
	fmt.Println("| ##  ##   ##   ##    ##           ##     ##  ###  ##        ## ##   ##  ##   ##   ##  ##   #   ##  ###    ##    |")
	fmt.Println("| ##  ##    #####   ######       ######   ##   ##  ##         ###    ##   ##  ##   ## ###   ##  ##   ##    ##    |")
	fmt.Println("+----------------------------------------------------------------------------------------------------------------+")
	fmt.Printf("|  Telegram Channel : %s@DVHOST_CLOUD%s	    |   YouTube : %syoutube.com/@dvhost_cloud%s\n", green, reset, red, reset)
	fmt.Println("+----------------------------------------------------------------------------------------------------------------+")
	fmt.Printf("|%sServer Country    |%s %s\n", green, reset, serverCountry)
	fmt.Printf("|%sServer IP         |%s %s\n", green, reset, serverIP)
	fmt.Printf("|%sServer ISP        |%s %s\n", green, reset, serverISP)
	fmt.Printf("|%sVersion           |%s %s\n", green, reset, version)
	fmt.Println("+----------------------------------------------------------------------------------------------------------------+")
	fmt.Printf("|%sPlease choose an option:%s\n", yellow, reset)
	fmt.Println("+----------------------------------------------------------------------------------------------------------------+")
	fmt.Println(items)
	fmt.Println("+----------------------------------------------------------------------------------------------------------------+")
	fmt.Print("\033[0m")

}
