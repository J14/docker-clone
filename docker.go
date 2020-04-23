package main

import (
	"flag"
	"fmt"
	"os"

	"./images"
	"./ps"
	"./run"
)

var flagPs = flag.NewFlagSet("flagPs", flag.ContinueOnError)

var all = flagPs.Bool("a", false, "Show all containers (default show just running)")

func main() {
	if len(os.Args[1:]) == 0 {
		help()
		os.Exit(1)
	}
	
	switch os.Args[1] {
	case "run":
		run.Run()
	case "ps":
		if err := flagPs.Parse(os.Args[2:]); err != nil {
			fmt.Fprintf(os.Stderr, "%v\n", err)
		}
		ps.Ps(*all)
	case "images":
		images.Images()
	default:
		help()
	}
}

func usage() {
	fmt.Printf("Usage: docker [OPTIONS] COMMAND\n\n")
}

func help() {
	usage()
	fmt.Printf("A self-sufficient runtime for containers\n\n")
	fmt.Println("Commands:")
	fmt.Printf("images\tList images\n")
	fmt.Printf("ps\tList containers\n")
	fmt.Printf("run\tRun a command in a new container\n")
}
