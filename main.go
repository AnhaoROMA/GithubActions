package main

import (
	"flag"
	"fmt"
)

var (
	club = flag.String("club", "Atletico de Madrid", "FC")
)

func main() {
	flag.Parse()
	fmt.Printf(`My favourite football club is "%s".`, *club)
}
