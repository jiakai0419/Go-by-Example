package main

import "syscall"
import "os"
import "os/exec"
import "fmt"

func main() {
	binary, lookErr := exec.LookPath("ls")
	fmt.Println(binary)
	if lookErr != nil {
		panic(lookErr)
	}
	args := []string{"ls", "-a", "-l", "-h"}
	env := os.Environ()
	execErr := syscall.Exec(binary, args, env)
	if execErr != nil {
		panic(execErr)
	}
}
