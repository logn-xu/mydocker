package main

import (
	"mydocker/container"

	log "github.com/sirupsen/logrus"
)

/*
clone 一个namespace隔离的进程 在子进程调用 /proc/self/exe  也就是调用自己 调用自身的init方法初始化容器的资源
*/
func Run(tty bool, cmd string) error {
	parent := container.NewParentProcess(tty, cmd)
	if err := parent.Start(); err != nil {
		log.Error(err)
	}

	err := parent.Wait()
	if err != nil {
		return err
	}

	return nil
}
