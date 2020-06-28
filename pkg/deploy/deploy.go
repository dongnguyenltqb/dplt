package deploy

import (
	"context"
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/yahoo/vssh"
)

type Config struct {
	Env  string
	Ips  []string
	Cmds []string
	Pem  string
}

type execResult struct {
	ip       string
	outTxt   string
	errTxt   string
	exitCode int
}

var waitDeploy sync.WaitGroup
var execResultChan chan execResult

func Deploy(cfg Config) error {
	cmd := cfg.Cmds[0]
	if len(cfg.Cmds) > 0 {
		for _, v := range cfg.Cmds {
			cmd = cmd + " && " + v
		}
	}
	return start(cmd, cfg.Pem, cfg.Ips)
}

func exec(cmd string, pem string, ip string) (result execResult) {
	vs := vssh.New().Start()
	config, _ := vssh.GetConfigPEM("ubuntu", pem)
	vs.AddClient(ip, config, vssh.SetMaxSessions(1))
	vs.Wait()
	ctx := context.Background()
	respChan := vs.Run(ctx, cmd, 3*time.Minute)
	for resp := range respChan {
		if err := resp.Err(); err != nil {
			log.Println(err)
			continue
		}
		outTxt, errTxt, _ := resp.GetText(vs)
		exitCode := resp.ExitStatus()
		result = execResult{
			ip:       ip,
			outTxt:   outTxt,
			errTxt:   errTxt,
			exitCode: exitCode,
		}
	}
	return
}

func watch() {
	for k := range execResultChan {
		fmt.Printf("IP  %v\n", k.ip)
		fmt.Printf("OUTPUT TEXT = %v\n", k.outTxt)
		fmt.Printf("ERROR TEXT = %v\n", k.errTxt)
		fmt.Printf("EXIT CODE = %v\n", k.exitCode)
		fmt.Println("=========================================================================================")
		waitDeploy.Done()
	}
}

func start(cmd string, pem string, ips []string) error {
	waitDeploy.Add(len(ips))
	execResultChan = make(chan execResult, len(ips))
	go watch()
	fmt.Println("START")
	for _, ip := range ips {
		fmt.Println("IP = ", ip)
		go func(cmd string, pem string, ip string) {
			execResultChan <- exec(cmd, pem, ip)
		}(cmd, pem, ip)
	}
	waitDeploy.Wait()
	close(execResultChan)
	return nil
}
