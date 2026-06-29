package proclet

import (
	"fmt"
	"time"
	"os/exec"
)



func Start(jobName string, arg ...strings){
	fmt.Println("Hello")
	cmd := exec.Command("sleep", "10")
	cmd.Start()
	fmt.Println("PID du process:", cmd.Process.Pid)
	time.Sleep(20*time.Second)
}
