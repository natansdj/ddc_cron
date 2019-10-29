package main

import (
	"fmt"
	"os/exec"
	"time"
)

func Output(str ...interface{}) {
	var output string
	output += time.Now().Format("01-01-2006 15:04:05.00") + " : "
	output += fmt.Sprintln(str...)

	fmt.Print(output)
}

func CurlGET(url string) {
	Output(url)

	curl := exec.Command("curl", url)
	out, err := curl.Output()

	if err != nil {
		fmt.Println(err)
		return
	} else {
		fmt.Println(string(out))
	}
}
