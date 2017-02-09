package core

import (
	"os/exec"
	"github.com/meixinyun/masscan-httpservice/g"
	"fmt"
	"log"
	"os"
	"io/ioutil"
	s "strings"
)

func ExecutionCmd(_execCmd string) {
	split := " "
	cmdWrapper :=s.Join([]string{g.Config().MasscanHome,_execCmd,g.OUT_PUT_FORMAT,g.Config().OutPutFileName},split)

	fmt.Printf("%s", cmdWrapper)
	cmd := exec.Command("/bin/bash", "-c", cmdWrapper)
	out, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println(err)
	}
	log.Println("execution cmd done, result=%s\n", string(out))
}

func GetResult() string {
	f, err := os.OpenFile(g.Config().OutPutFileName, os.O_RDWR, os.ModeType)
	if err != nil {
		log.Println("open result file failed")
		panic(err)
	}
	defer f.Close()
	b, err2 := ioutil.ReadAll(f)
	if err2 != nil {
		panic(err2)
	}
	var ret = string(b)
	fmt.Printf("%v", ret)
	return ret
}