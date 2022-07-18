/**
* @Author: scjtqs
* @Date: 2022/7/18 11:22
* @Email: scjtqs@qq.com
 */
package main

import (
	"flag"
	"fmt"
	"github.com/scjtqs2/fqsign_go/app"
	log "github.com/sirupsen/logrus"
	"os"
)

var (
	c string
	h bool

	Version string
	Build   string
)

func init() {
	flag.StringVar(&c, "c", "config.yaml", "configuration filename")
	flag.BoolVar(&h, "h", false, "this help")
	flag.Parse()
}

func help() {
	fmt.Printf(`scjtqs fqsign onetime script
version: %s
built-on: %s

Usage:

fsc [OPTIONS]

Options:
`, Version, Build)

	flag.PrintDefaults()
	os.Exit(0)
}

func main() {
	if h {
		help()
	}
	ct, err := bootstrap()
	if err != nil {
		log.Fatalf("faild init bootstrap err=%v", err)
	}
	app.RunClient(ct)
}
