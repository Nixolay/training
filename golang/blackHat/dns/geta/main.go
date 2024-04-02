package main

import (
	"errors"
	"flag"

	"github.com/miekg/dns"
)

func main() {
	serv := flag.String("server", "8.8.8.8:53", "The DNS server to use.")
	domain := flag.String("domain", "google.com", "The domain to perform guessing against.")
	flag.Parse()

	ips, err := lookupA(*domain, *serv)
	if err != nil {
		panic(err.Error())
	}

	for _, ip := range ips {
		println(ip)
	}
}

func lookupA(fqdn, serverAddr string) ([]string, error) {
	var (
		msg dns.Msg
		ips []string
	)

	msg.SetQuestion(dns.Fqdn(fqdn), dns.TypeA)

	inMsg, err := dns.Exchange(&msg, serverAddr)
	if err != nil {
		return nil, err
	}

	if len(inMsg.Answer) < 1 {
		return nil, errors.New("no answer")
	}

	for _, answer := range inMsg.Answer {
		if a, ok := answer.(*dns.A); ok {
			ips = append(ips, a.A.String())
		}
	}

	return ips, nil
}
