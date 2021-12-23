package main

import (
	"fmt"
	"strings"

	"github.com/mbndr/figlet4go"
)

func BootBanner(label string, version string) string {
	fig := figlet4go.NewAsciiRender()
	banner, _ := fig.Render(label)
	return fmt.Sprintf("boot\n%v     %v\n",
		strings.TrimRight(banner, "\r\n"),
		version)
}
