package main

import (
	"flag"
	"DOIpUpdate"
	"fmt"
	"os"
)

var doKey = flag.String("doKey", "", "Digital Ocaean API Key");
var domain = flag.String("domain", "", "Domain for update");

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
		panic(err)
	}


	ip := DOIpUpdate.GetIP()
	if (ip == recordData.DomainRecord.Data) {
		fmt.Println("IPs are same, do not need update");
		os.Exit(0);
	}

	err = DOIpUpdate.UpdateRecord(client, recordData, ip);
	if (err != nil) {
		panic(err)
	}

	fmt.Println("IP update done to: " + ip);
}