package modules

import "fmt"

const (
    Red    = "\033[31m"
    Green  = "\033[32m"
    Reset  = "\033[0m"
)

func CrackDone(user string, passw string, crackSec float64) {
	fmt.Println("<=============================>")
	fmt.Println(Green+"Brute force done!"+Reset)
	fmt.Println("<=============================>")
	fmt.Println("User:", user)
	fmt.Println("Password:", passw)
	fmt.Println("<=============================>")
	fmt.Printf("%vFTP client cracked in %v%f%v seconds%v\n",Green, Reset, crackSec, Green, Reset)
	fmt.Println("<=============================>")
}