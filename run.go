package main

import (
	"fmt"
	"mydocker/container"
	"os"
	"strings"

	log "github.com/sirupsen/logrus"
)

/*
clone 一个namespace隔离的进程 在子进程调用 /proc/self/exe  也就是调用自己 调用自身的init方法初始化容器的资源
*/
func Run(tty bool, cmdArray []string) error {
	parent, writePipe := container.NewParentProcess(tty)
	if parent == nil {
		log.Errorf("New parent process error")
		return fmt.Errorf("parent not fork")
	}

	if err := parent.Start(); err != nil {
		log.Errorf("Run parent.Start err: %v", err)
	}

	//在子进程创建后通过管道传递参数
	err := sendInitCommand(cmdArray, writePipe)
	if err != nil {
		return err
	}

	err = parent.Wait()
	if err != nil {
		return err
	}

	return nil
}

// 将参数发送给子进程
func sendInitCommand(cmdArray []string, writePipe *os.File) error {
	defer func() {
		writePipe.Close()
	}()

	command := strings.Join(cmdArray, " ")
	log.Infof("command all is %s", command)

	//将命令字符串写入到管道
	_, err := writePipe.WriteString(command)
	if err != nil {
		log.Errorf("write commond err: %s", err)
		return err
	}
	return nil
}
