package DOIpUpdate

import (
	"github.com/digitalocean/godo"
	"golang.org/x/net/context"
	"errors"
	"regexp"
)

type DomainRecordData struct {
	Domain string
	DomainRecord godo.DomainRecord
}

func GetDomainRecord(client *godo.Client, domainName string) (DomainRecordData, error) {
	var result DomainRecordData
	opt := &godo.ListOptions{}

	var re = regexp.MustCompile(`(?U)^(.*)\.{0,1}([\w-]+\.[a-z]+)$`)
	matches := re.FindAllStringSubmatch(domainName, -1)
	subDomain := matches[0][1]
	result.Domain = matches[0][2]

	for {
		records, resp, err := client.Domains.Records(context.TODO(), result.Domain, opt)
		if err != nil {
			return result, err
		}

		for _, r := range records {
			if (r.Name == subDomain && r.Type == "A") {
				result.DomainRecord = r
				return result, nil
			}
		}

		if resp.Links == nil || resp.Links.IsLastPage() {
			break
		}

		page, err := resp.Links.CurrentPage()
		if err != nil {
			return result, err
		}

		opt.Page = page + 1
	}

	return result, errors.New("No found domain")
}

func UpdateRecord(client *godo.Client, data DomainRecordData, ip string) (error)  {

	updateRequest := &godo.DomainRecordEditRequest{
		Type:     data.DomainRecord.Type,
		Name:     data.DomainRecord.Name,
		Data:     ip,
		Priority: data.DomainRecord.Priority,
		Port:     data.DomainRecord.Port,
		Weight:   data.DomainRecord.Weight,
	}

	_, _, err := client.Domains.EditRecord(context.TODO(), data.Domain, data.DomainRecord.ID, updateRequest)
	if (err != nil) {
		return err
	}

	return nil
}