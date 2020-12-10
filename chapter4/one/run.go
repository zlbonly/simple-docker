package main

import (
	"example/chapter4/one/cgroups"
	"example/chapter4/one/cgroups/subsystems"
	"example/chapter4/one/container"
	log "github.com/sirupsen/logrus"
	"os"
	"strings"
)

func Run(tty bool, cmdArray []string, res *subsystems.ResourceConfig) {
	parent, writePipe := container.NewParentProcess(tty)
	if parent == nil {
		log.Errorf("New parent process error")
		return
	}
	if err := parent.Start(); err != nil {
		log.Error(err)
	}

	// use mydocker-cgroup as cgroup name
	cgroupmanager := cgroups.NewCgroupManager("mydocker-cgroup")
	defer cgroupmanager.Destory()
	cgroupmanager.Set(res)
	cgroupmanager.Apply(parent.Process.Pid)

	sendInitCommand(cmdArray, writePipe)
	parent.Wait()
	os.Exit(0)
}

func sendInitCommand(cmdArray []string, writePipe *os.File) {
	command := strings.Join(cmdArray, " ")
	log.Infof("command all is %s", command)
	writePipe.WriteString(command)
	writePipe.Close()
}
