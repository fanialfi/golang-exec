package lib

import (
	"fmt"
	"os/exec"
	"runtime"
)

func Rekomendasi() {
	if runtime.GOOS == "windows" {
		output, err := exec.Command("cmd", "/C", "git config user.name").Output()
		fmt.Println("runtime is windows")

		if err != nil {
			panic(err.Error())
		} else {
			fmt.Println(output)
		}
	} else {
		output, err := exec.Command("bash", "-c", "notify-send \"hello world\" \"hello bro\"").Output()
		fmt.Println("runtime is", runtime.GOOS)

		if err != nil {
			fmt.Println(output)
			fmt.Println(err.Error())
		} else {
			fmt.Println(string(output))
		}
	}
}
