package main

import (
	"github.com/codegangsta/cli"
	"net"
	"os"
	"time"
)

const (
	NAGIOS_OK       int = 0
	NAGIOS_WARNING  int = 1
	NAGIOS_CRITICAL int = 2
	NAGIOS_UNKNOWN  int = 3
)

var ch = make(chan int)
var exitCode int
var quiet bool

// TODO
// var ok, warning, critical, unknown int

func main() {

	app := cli.NewApp()
	app.Name = "check_portsnoop"
	app.Version = "1.0.0"
	app.Author = "Paul Swanson"
	app.Usage = "Test one or more TCP ports"
	app.Flags = []cli.Flag{
		cli.BoolFlag{"quiet, q", "Be quiet"},
		cli.IntFlag{"timeout, t", 100, "Timeout for connections in milliseconds, default 100ms"},
		// TODO
		// cli.IntFlag{"ok, o", 1, "Nagios: OK threshold"},
		// cli.IntFlag{"warning, w", 1, "Nagios: Warning threshold"},
		// cli.IntFlag{"critical, c", 0, "Nagios: Critical threshold"},
		// cli.IntFlag{"unknown", u", -1, "Nagios: Unknown threshold"},
	}

	cli.AppHelpTemplate = `NAME:
{{.Name}} - {{.Usage}}

USAGE:
{{.Name}} [global options] [arguments...]

For example: {{.Name}} -quiet 8.8.8.8:53 www.google.com:80

VERSION:
{{.Version}}

GLOBAL OPTIONS:
{{range .Flags}}{{.}}
{{end}}
`
	app.Action = func(c *cli.Context) {

		var portsOK, portCount int
		ports := c.Args()
		portCount = len(ports)

		quiet = c.Bool("quiet")

		if portCount > 0 {
			// Test ports concurrently
			for _, p := range ports {
				go portSnoop(p, time.Millisecond*time.Duration(c.Int("timeout")))
			}
			// Wait for all tests to complete
			for i := 0; i < portCount; i++ {
				portsOK += <-ch
			}
			if !quiet {
				println(portsOK, "successful connection(s)")
			}
			exitCode = nagiosExitCode(portCount, portsOK)
		} else {
			// If no ports supplied, show help
			cli.ShowAppHelp(c)
			exitCode = NAGIOS_UNKNOWN
		}
	}

	app.Run(os.Args)

	os.Exit(exitCode)
}

func nagiosExitCode(portCount, portsOK int) int {

	var c int

	// Default Nagios behaviour:
	// 	All down is CRITICAL
	//	All up is OK
	//	For 2 or more ports, only one up is WARNING

	if portsOK == 0 {
		c = NAGIOS_CRITICAL
	} else {
		if portCount == portsOK {
			c = NAGIOS_OK
		} else {
			c = NAGIOS_WARNING
		}
	}
	return c
}

func portSnoop(address string, timeout time.Duration) {
	_, err := net.DialTimeout("tcp", address, timeout)
	if err != nil {
		if !quiet {
			println(address, "FAIL.")
		}
		ch <- 0
	} else {
		if !quiet {
			println(address, "OK.")
		}
		ch <- 1
	}
}
