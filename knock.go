package main

import (
    "net"
    "os"
    "time"
    "fmt"
    "strconv"
    "github.com/docopt/docopt-go"
)

func main() {

    usage := `Go Knock

Usage:
  knock [-d=<ms>] [-u] <host> <port>...
  knock -h | --help
  knock -v | --version

Options:
  -h, --help         Show this screen.
  -v, --version      Show version.
  -d, --delay=<ms>   Delay between ports in milliseconds [default: 300].
  -u, --udp          Send UDP packets instead of TCP [default: false].`

    arguments, _ := docopt.Parse(usage, nil, true, "Go Knock 1.0", false)
    // fmt.Println(arguments)

    var host, dest, delay, packetType string
    var ports, portargs []string

    if val, ok := arguments["--delay"].(string); ok {
        delay = val
    } else {
        delay = "300"
    }

    // Check if host set
    if val, ok := arguments["<host>"].(string); ok {
        host = val
    } else {
        fmt.Println("No host specified")
        os.Exit(1)
    }

    // Check if ports set
    if val, ok := arguments["<port>"].([]string); ok {
        portargs = val
    } else {
        fmt.Println("No ports specified")
        os.Exit(1)
    }

    // Check if UDP flag set
    if val, ok := arguments["--udp"].(bool); ok {
        if val {
            packetType = "udp"
        } else {
            packetType = "tcp"
        }
    }
    // fmt.Println(packetType)


    // Check if port arguments are valid ports
    for i := 0; i < len(portargs); i++ {
        // fmt.Println(portargs[i])
        if p, err := strconv.Atoi(portargs[i]); err == nil {
            // fmt.Printf("%q is a number.\n", portargs[i])
            // Check if port is in valid range
            if p < 1 || p > 65535 {
                fmt.Printf("%q is larger than 65535.\n", portargs[i])
                os.Exit(1)
            } else {
                ports = append(ports, portargs[i])
            }
        }
    }

    // Check if more than one port
    if len(ports) > 1 {
        for i := 0; i < len(ports); i++ {
            port := ports[i]
            dest = fmt.Sprintf("%s:%s", host, port)
            // Timeout is irrelevant, just need a knock
            conn, err := net.DialTimeout(packetType, dest, 1*time.Millisecond)
            if err != nil {
                // Expected error
                // fmt.Println(err)
            } else {
                defer conn.Close()
            }
            // Timeout sleep
            if ms, err := strconv.Atoi(delay); err == nil {
                time.Sleep(time.Duration(ms)*time.Millisecond)
            }
        }
    }
}
