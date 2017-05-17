package main

import (
	"flag"
	"DOIpUpdate"
	"os"
	"log/syslog"
)

var doKey = flag.String("doKey", "", "Digital Ocaean API Key")
var domain = flag.String("domain", "", "Domain for update")
var logger, _ = syslog.New(syslog.LOG_INFO, "DO IP updater");

func init()  {
	flag.Parse();
	if (*doKey == "") {
		panic("Please provide Digital Ocean API key")
	}

	if (*domain == "") {
		panic("Please provide Doamin to update")
	}
}

func main()  {
	client := DOIpUpdate.GetClientWithToken(*doKey)
	recordData, err := DOIpUpdate.GetDomainRecord(client, *domain)
	if (err != nil) {
		logger.Err(err.Error())
		panic(err)
	}


	ip := DOIpUpdate.GetIP()
	if (ip == recordData.DomainRecord.Data) {
		logger.Info("IPs are same, do not need update")
		os.Exit(0)
	}

	err = DOIpUpdate.UpdateRecord(client, recordData, ip)
	if (err != nil) {
		logger.Err(err.Error())
		panic(err)
	}

	logger.Info("IP update done to: " + ip)
}