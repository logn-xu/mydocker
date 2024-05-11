package main

import (
	"fmt"
	"mydocker/container"

	"github.com/urfave/cli"

	log "github.com/sirupsen/logrus"
)

var runCommand = cli.Command{
	Name:  "run",
	Usage: `Create a container with namespace and cgroups limit mydocker run -it [command]`,
	Flags: []cli.Flag{
		cli.BoolFlag{
			Name:  "it", // 合并参数 -i -t
			Usage: "enable tty",
		},
	},

	/*
		run指定函数
		1. 判断参数是否包含command
		2. 获取用户指定的command
		3. 调用Run function 去准备启动容器
	*/
	Action: func(context *cli.Context) error {
		if len(context.Args()) < 1 {
			return fmt.Errorf("missing container command")
		}

		cmd := context.Args()
		tty := context.Bool("it")
		err := Run(tty, cmd)
		if err != nil {
			return err
		}
		return nil
	},
}

var initCommand = cli.Command{
	Name:  "init",
	Usage: "Init container process run user's process in container. Do not call it outside",
	/*
		1. 获取传递过来的 command参数
		2. 执行容器初始化操作
	*/

	Action: func(context *cli.Context) error {
		log.Infof("init come on")
		cmd := context.Args().Get(0)
		log.Infof("command: %s", cmd)
		err := container.RunContainerInitProcess()
		if err != nil {
			log.Error(err)
			return err
		}
		return nil
	},
}
