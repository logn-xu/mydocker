package container

import (
	"os"
	"os/exec"
	"syscall"
)

/*
1. 调用/proc/self/exe 对进程初始化
2. 调用initCommand 初始化进程和环境
3. clone fork一个新进程并使用namespace隔离新创建的进程和外部环境
4. 如果用户指定了 -ti 参数将当前进程的输入输出导入到当前终端的输入输出
*/
func NewParentProcess(tty bool, command string) *exec.Cmd {
	args := []string{"init", command}
	cmd := exec.Command("/proc/self/exe", args...)
	cmd.SysProcAttr = &syscall.SysProcAttr{
		Cloneflags: syscall.CLONE_NEWUTS | syscall.CLONE_NEWPID | syscall.CLONE_NEWNS | syscall.CLONE_NEWNET | syscall.CLONE_NEWIPC,
	}

	if tty {
		cmd.Stdin = os.Stdin
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
	}

	return cmd
}
