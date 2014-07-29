package main

import (
	"flag"
	"fmt"
	"github.com/ezotrank/logsend/logsend"
	"os"
)

var (
	logDir       = flag.String("log-dir", "./tmp", "log directories")
	config       = flag.String("config", "config.json", "path to config.json file")
	dbhost       = flag.String("db-host", "localhost:8086", "db host")
	dbuser       = flag.String("db-user", "root", "db user")
	dbpassword   = flag.String("db-password", "root", "db-password")
	database     = flag.String("database", "test1", "database")
	udp          = flag.Bool("udp", false, "send series over udp")
	check        = flag.Bool("check", false, "check config.json")
	debug        = flag.Bool("debug", false, "turn on debug messages")
	sendBuffer   = flag.Int("send-buffer", 8, "send buffer")
	stopContinue = flag.Bool("stop-continue", false, "watching folder for new files")
	memprofile   = flag.String("memprofile", "", "memory profiler")
)

func main() {
	flag.Parse()

	logsend.Conf.SendBuffer = *sendBuffer
	logsend.Conf.Debug = *debug
	logsend.Conf.DBHost = *dbhost
	logsend.Conf.DBUser = *dbuser
	logsend.Conf.DBPassword = *dbpassword
	logsend.Conf.DBName = *database
	logsend.Conf.ContinueWatch = !*stopContinue
	logsend.Conf.UDP = *udp
	logsend.Conf.WatchDir = *logDir
	logsend.Conf.Memprofile = *memprofile

	if *check {
		_, err := logsend.LoadConfig(*config)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		fmt.Println("ok")
		os.Exit(0)
	}

	logsend.WatchFiles(*logDir, *config)
}
