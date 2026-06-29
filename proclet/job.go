package proclet

import (
	"fmt"
	"errors"
	"os/exec"
)

type Job struct {
	Name string
	PID int
	cmd *exec.Cmd
}

func TestJob(jobName string, arg ...string) (*exec.Cmd, error){
	var err error
	if jobName == "" {
		return nil, errors.New("empty name")
  }
	fmt.Println("Start:", jobName)
	cmd := exec.Command(jobName, arg...)
	err = cmd.Start()
	if err != nil {
		return nil, err
	}
	err = cmd.Wait()
	if err != nil {
		return nil, err
	}
	fmt.Println("End:", jobName)
	return cmd, nil
}

func Start(jobName string, arg ...string) (*exec.Cmd, error){
	var err error
	if jobName == "" {
		return nil, errors.New("empty name")
  }
	fmt.Println("Start:", jobName)
	cmd := exec.Command(jobName, arg...)
	err = cmd.Start()
	if err != nil {
		return nil, err
	}
	fmt.Println("End:", jobName)
	return cmd, nil
}
