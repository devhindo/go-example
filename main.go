package main

import (
	"os/exec"
	"runtime"
	"fmt"
)

func main() {
    err := OpenDefault("F:\\task.txt")
    if err != nil {
        fmt.Println(err)
    }
}

func OpenDefault(fileOrURL string) error {
	var cmd string
	var args []string

	switch runtime.GOOS {
	case "windows":
        cmd = "rundll32"
        args = []string{"url.dll,FileProtocolHandler"}
	case "darwin":
		cmd = "open"
	default: // "linux", "freebsd", "openbsd", "netbsd"
		cmd = "xdg-open"
	}
	args = append(args, fileOrURL)
	return exec.Command(cmd, args...).Start()
}