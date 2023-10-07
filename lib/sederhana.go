package lib

import (
	"fmt"
	"os/exec"
)

func Sederhana() {
	var output1, err = exec.Command("ls").Output()
	if err != nil {
		panic(err.Error())
	} else {
		fmt.Println(string(output1))
	}

	output3, err3 := exec.Command("notify-send", "-i", "fish", "-t", "5000", "-A", "ok=OK", "-A", "wait=WAIT", "Pomodoro Time", "Time To Sleep").Output()
	if err == nil {
		fmt.Println(string(output3))
	} else {
		panic(err3.Error())
	}

	output2, err2 := exec.Command("git", "config", "user.name").Output()
	if err != nil {
		panic(err2.Error())
	} else {
		fmt.Println(string(output2))
	}
}
