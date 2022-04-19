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
<<<<<<< HEAD

// 15000000건
// 8KB * 20000000 = about 150GB per day
// 4.5TB about 5TB per month
// 100TB per year -> $2000 (60TB $1400)
// 일평균 트립 50만건 write transaction 50만
=======
>>>>>>> master
