package main

import (
	"flag"
	"log"
	"os"
	"os/exec"
	"strings"
)

var (
	TARGET_DIR = os.Getenv("HOME") + "/IdeaProjects/"
)

func main() {
	flag.Parse()
	zip := flag.Arg(0)

	cmd := exec.Command("unzip", "-a", zip, "-d", TARGET_DIR)
	err := cmd.Run()
	if err != nil {
		log.Fatalln("non zero exist code while extract..", err)
	}

	exec.Command("idea", TARGET_DIR+strings.Split(zip, ".")[0]+"/pom.xml").Run()
}
