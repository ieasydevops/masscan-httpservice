package main

import (
	"flag"
	"fmt"
	"github.com/meixinyun/masscan-httpservice/core"
	"github.com/meixinyun/masscan-httpservice/g"
	myhttp "github.com/meixinyun/masscan-httpservice/http"
	"net/http"
	"os"
)

func configMasscanRoutes() {
	http.HandleFunc("/masscan", func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		cmd := r.Form.Get("cmd")
		core.ExecutionCmd(cmd)
		result := core.GetResult()
		response := fmt.Sprintf("%s\r\n",result)
		w.Write([]byte(response))
	})
}

func main() {
	cfg := flag.String("c", "cfg.json", "configuration file")
	version := flag.Bool("v", false, "show version")

	flag.Parse()

	if *version {
		fmt.Println(g.VERSION)
		os.Exit(0)
	}


	g.ParseConfig(*cfg)
	configMasscanRoutes()
	go myhttp.Start()
	select {}

}
