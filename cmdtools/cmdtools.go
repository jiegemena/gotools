package cmdtools

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os/exec"
	"strings"
	"time"

	"github.com/jiegemena/gotools/stringtools"
)

func CmdTest() {
	//删除C:\Users\Administrator\Desktop目录下的index.html文件
	c := exec.Command("cmd", "/C", "python", "D:\\codehome\\pyhome\\ptest\\index.py")
	if err := c.Run(); err != nil {
		fmt.Println("Error: ", err)
	}
}

func RunCmd(cmdStr string, runpath string) string {
	list := strings.Split(cmdStr, " ")
	cmd := exec.Command(list[0], list[1:]...)
	var out bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &out
	cmd.Dir = runpath
	cmd.Stderr = &stderr

	go func() {
		for {
			time.Sleep(time.Second)
			fmt.Println(stderr.String())
		}

	}()

	err := cmd.Run()
	if err != nil {
		return stderr.String()
	} else {
		return out.String()
	}
}

func RunCommand(cmdStr string, runpath string) error {
	list := strings.Split(cmdStr, " ")
	cmd := exec.Command(list[0], list[1:]...)
	cmd.Dir = runpath

	// 命令的错误输出和标准输出都连接到同一个管道
	stdout, err := cmd.StdoutPipe()
	cmd.Stderr = cmd.Stdout

	if err != nil {
		return err
	}

	if err = cmd.Start(); err != nil {
		return err
	}
	// 从管道中实时获取输出并打印到终端
	for {
		tmp := make([]byte, 1024)
		_, err := stdout.Read(tmp)
		fmt.Print(stringtools.ConvertByte2String(tmp, stringtools.GB18030))
		if err != nil {
			break
		}
	}

	if err = cmd.Wait(); err != nil {
		return err
	}
	return nil
}

func ExecCommand(cmdStr string, runpath string) bool {
	list := strings.Split(cmdStr, " ")
	cmd := exec.Command(list[0], list[1:]...)
	// cmd := exec.Command(commandName, params...)
	cmd.Dir = runpath
	//显示运行的命令
	fmt.Println(cmd.Args)

	stdout, err := cmd.StdoutPipe()

	if err != nil {
		fmt.Println(err)
		return false
	}

	cmd.Start()

	reader := bufio.NewReader(stdout)

	//实时循环读取输出流中的一行内容
	for {
		line, err2 := reader.ReadString('\n')
		if err2 != nil || io.EOF == err2 {
			break
		}
		var data1 []byte = []byte(line)
		fmt.Println(stringtools.ConvertByte2String(data1, stringtools.GB18030))
	}

	cmd.Wait()
	return true
}
