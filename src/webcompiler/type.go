package webcompiler

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
)

type ExecInfo struct {
	Command string
	Args    []string
	In      string
}

type ExecResult struct {
	Status string
	Out    string
	Err    string
}

type CompileInfo struct {
	Compiler string
	Args     []string
	Code     string
	In       string
}

func (p *ExecInfo) exec() (result *ExecResult) {
	result = new(ExecResult)

	cmd := exec.Command(p.Command, p.Args...)

	in := new(bytes.Buffer)
	out := new(bytes.Buffer)
	err := new(bytes.Buffer)

	in.Write([]byte(p.In))

	cmd.Stdin = in
	cmd.Stdout = out
	cmd.Stderr = err
	e := cmd.Run()
	if e != nil {
		result.Status = "failed"
		fmt.Println(e.Error())
	} else {
		result.Status = "ok"
	}
	result.Out = out.String()
	result.Err = err.String()
	return
}

func (p *CompileInfo) CompileAndRun() (result *ExecResult) {
	//将代码写入文件并添加编译参数
	ioutil.WriteFile("D:\\test.c", []byte(p.Code), 0666)
	p.Args = append(p.Args, "D:\\test.c", "-oD:\\test.exe")

	//执行编译
	cmd := ExecInfo{Command: Compiler[p.Compiler], Args: p.Args, In: ""}
	result = cmd.exec()
	//fmt.Println(out.String())
	//判断是否编译成功
	if isFileExists("D:\\test.exe") {
		cmd = ExecInfo{Command: "D:\\test.exe", Args: nil, In: p.In}
		result = cmd.exec()
		os.Remove("D:\\test.c")
		os.Remove("D:\\test.exe")
	}
	return
}
