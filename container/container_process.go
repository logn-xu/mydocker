package container

import (
	"os"
	"os/exec"
	"syscall"

	log "github.com/sirupsen/logrus"
)

/*
1. 调用/proc/self/exe 对进程初始化
2. 调用initCommand 初始化进程和环境
3. clone fork一个新进程并使用namespace隔离新创建的进程和外部环境
4. 如果用户指定了 -ti 参数将当前进程的输入输出导入到当前终端的输入输出
*/
func NewParentProcess(tty bool) (*exec.Cmd, *os.File) {
	//使用管道在进程间传递参数
	// 创建匿名管道用于传递参数，将readPipe作为子进程的ExtraFiles，子进程从readPipe中读取参数
	// 父进程中则通过writePipe将参数写入管道
	readPipe, writePipe, err := os.Pipe()
	if err != nil {
		log.Errorf("New pipe error %v", err)
		return nil, nil
	}

	cmd := exec.Command("/proc/self/exe", "init")
	cmd.SysProcAttr = &syscall.SysProcAttr{
		Cloneflags: syscall.CLONE_NEWUTS | syscall.CLONE_NEWPID | syscall.CLONE_NEWNS | syscall.CLONE_NEWNET | syscall.CLONE_NEWIPC,
	}

	if tty {
		cmd.Stdin = os.Stdin
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
	}

	//建readPipe 作为ExtraFiles cmd执行会携带文件句柄去创建子进程
	cmd.ExtraFiles = []*os.File{readPipe}

	return cmd, writePipe
}
