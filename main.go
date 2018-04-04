//
// process.go : Contains main package drivers and stuff
// Written By : @codingo
//		@ice3man
//
// Distributed Under MIT License
// Copyrights (C) 2018 Ice3man
//

package main

import (
	"fmt"
	"flag"
	"os"

	"subfinder/libsubfinder/helper"
	"subfinder/libsubfinder/engines/passive"
)


var banner = `
             __     ___ __          __            
.-----.--.--|  |--.'  _|__.-----.--|  .-----.----.
|__ --|  |  |  _  |   _|  |     |  _  |  -__|   _|
|_____|_____|_____|__| |__|__|__|_____|_____|__|  

`

// Parses command line arguments into a setting structure
func ParseCmdLine() (state *helper.State, err error) {

	// Initialize current state and read Config file
	s, err := helper.InitState()
	if err != nil {
		return &s, err
	}

	flag.BoolVar(&s.Verbose, "v", false, "Verbose output")
	flag.BoolVar(&s.Color, "c", true, "Use colour in outpout")
	flag.IntVar(&s.Threads, "t", 10, "Number of concurrent threads")
	flag.StringVar(&s.Domain, "d", "", "Domain to find subdomains for")
	flag.BoolVar(&s.Recursive, "r", true, "Use recursion to find subdomains")

	flag.Parse()

	return &s, nil
}


func main() {

	fmt.Println(banner)
	fmt.Printf("\nSubFinder v0.1.0 			Made with ❤ by @Ice3man")
	fmt.Printf("\n===============================================================")

	state, err := ParseCmdLine()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// Improve Usage guide here
	if state.Domain == "" {
		fmt.Printf("\n\nsubfinder: Missing domain argument\nTry './subfinder -h' for more information\n")
		os.Exit(1)
	}

	passive.PassiveDiscovery(state)
}
