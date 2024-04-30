package main

import (
	"crypto/tls"
	"fmt"
	"net/http"
)

func main() {

	//supports TLS1.2 AND TLS1.3 Nmap scan output

	tlsConfiguration := tls.Config{
		MaxVersion: tls.VersionTLS13,
		MinVersion: tls.VersionTLS13,
	}
	//mudar a configuração do TLS
	http := http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tlsConfiguration,
		},
	}

	resp, err := http.Get("https://distopia.savi2w.workers.dev/")

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	resp.Body.Close()

	fmt.Println("O status da requisição foi: ", resp.Status)
	fmt.Println("O Corpo foi: ", string(body))

	

}
