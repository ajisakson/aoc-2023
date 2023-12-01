package main

import (
	"os/exec"
	"strconv"
)

func main() {
	for i := 1; i <= 25; i++ {
		app := "mkdir"
		module := "day" + strconv.Itoa(i)
		cmd := exec.Command(app, module)
		err := cmd.Run()
		if err != nil {
			panic(err)
		}

		app = "go"
		cmd = exec.Command(app, "mod", "init", module)
		cmd.Dir = module
		err = cmd.Run()
		if err != nil {
			panic(err)
		}

		app = "touch"
		cmd = exec.Command(app, module+".go")
		cmd.Dir = module
		err = cmd.Run()
		if err != nil {
			panic(err)
		}
		bob := []byte("package " + module + "\n\nfunc main() {}")
		cmd = exec.Command("sh", "-c", "echo \""+string(bob)+"\" > "+module+".go")
		cmd.Dir = module
		err = cmd.Run()
		if err != nil {
			panic(err)
		}
	}
}
