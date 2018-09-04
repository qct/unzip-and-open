package main

import (
	"flag"
	"log"
	"os/exec"
	"strings"
)

const (
	TARGET_DIR = "~/Downloads/"
)

func main() {
	flag.Parse()
	zip := flag.Arg(0)

	cmd := exec.Command("unzip", "-a", zip, "-d", TARGET_DIR)
	err := cmd.Run()
	if err != nil {
		log.Println("non zero exist code but we proceed anyway..", err)
	}

	exec.Command("idea", TARGET_DIR+strings.Split(zip, ".")[0]+"/pom.xml").Run()
}
