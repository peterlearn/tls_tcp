package main

import (
	"bytes"
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	pool := x509.NewCertPool()
	caCertPath := "/var/run/test/ca.crt"

	caCrt, err := os.ReadFile(caCertPath)
	if err != nil {
		fmt.Println("ReadFile err", err)
		return
	}
	pool.AppendCertsFromPEM(caCrt)

	cliCrt, err := tls.LoadX509KeyPair("/var/run/test/client.crt", "/var/run/test/client.key")
	if err != nil {
		fmt.Println("LoadX509KeyPair err:", err)
		return
	}
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{
			RootCAs:            pool,
			Certificates:       []tls.Certificate{cliCrt},
			InsecureSkipVerify: true, //客户端关闭对服务端的验证
		},
	}
	client := &http.Client{Transport: tr}
	jsonStr := "{\"name\":\"wang\",\"age\":29}"
	req := bytes.NewBuffer([]byte(jsonStr))
	body_type := "application/json;charset=utf-8"
	resp, err := client.Post("https://test.com:8088/post", body_type, req)
	if err != nil {
		fmt.Println("Get error:", err)
		return
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	fmt.Println(string(body))
}
