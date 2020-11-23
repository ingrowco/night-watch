package configurator

import (
	"fmt"
	"time"
)

var (
	AppVersion string
	BuildTime  string
	GitCommit  string
)

const niw = `888b    888 d8b 888       888 
8888b   888 Y8P 888   o   888 
88888b  888     888  d8b  888 
888Y88b 888 888 888 d888b 888 
888 Y88b888 888 888d88888b888 
888  Y88888 888 88888P Y88888 
888   Y8888 888 8888P   Y8888 
888    Y888 888 888P     Y888`

func PrintBuildInformation() {
	buildDate, _ := time.Parse(time.RFC3339, BuildTime)

	fmt.Println(niw)
	fmt.Printf(
		"Version:\t\t%s\n"+
			"Build Commit:\t\t%s\n"+
			"Build Date:\t\t%s\n",
		AppVersion,
		GitCommit,
		buildDate.UTC(),
	)
}
