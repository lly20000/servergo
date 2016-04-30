package main

import (
	"DB"
	"connector"
	"errors"
	"flag"
	"fmt"
	"github.com/larspensjo/config"
	"os"
	//"strconv"
	//"common"
)

var (
	serverfile = flag.String("serverfile", "./config/server.ini", "Server Config file")
)

func main() {

	ipadrr, portstr, err := loadConfig()
	fmt.Println("err = ", err)
	if err != nil {
		fmt.Println(os.Stderr, " server ini read error\n")
		os.Exit(1)
	}
	fmt.Println("port = ", portstr)
	fmt.Println("ipadrr = ", ipadrr)

	//start DB
	err = DB.ConnectDB()
	if err != nil {
		fmt.Println("DB.ConnectDB() error")
	}
	//test code
	// var i int64
	// for i = 0; i < 6000; i++ {
	// 	strbyte := common.Int64ToBytes(i)
	// 	fmt.Println("i =  %d =====%s", i, strbyte)
	// }

	//start server
	connector.StartServer(ipadrr, portstr)
	//test proto buffer
	//msg := &myproto.
}

func loadConfig() (string, string, error) {
	var TOPIC = make(map[string]string)
	//server config
	cfg, err := config.ReadDefault(*serverfile)
	if err != nil {
		fmt.Println("config.ReadDefault error")
	}
	//
	//Initialized topic from the configuration
	if cfg.HasSection("server") {
		section, err := cfg.SectionOptions("server")
		if err == nil {
			for _, v := range section {
				options, err := cfg.String("server", v)
				if err == nil {
					TOPIC[v] = options
				}
			}
		}
	}

	if len(TOPIC) != 2 {
		return "", "", errors.New("par wrong")
	}
	fmt.Println("len(TOPIC) = ", len(TOPIC))
	return TOPIC["ip"], TOPIC["port"], nil
}
