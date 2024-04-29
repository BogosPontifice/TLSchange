package main

import (
	"crypto/tls"
	"fmt"
	"net/http"
)

func main() {

	//supports TLS1.2 AND TLS1.3 Nmap scan output

	tlsCustomConfig := tls.Config{
		MaxVersion: tls.VersionTLS13,
		MinVersion: tls.VersionTLS13,
	}
	//mudar a configuração do TLS
	http := http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tlsCustomConfig,
		},
	}

	resp, err := http.Get("https://distopia.savi2w.workers.dev/")

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println(resp.Status)

	resp.Body.Close()

}
