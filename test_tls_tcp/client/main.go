package main

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io"
	"os"
	"time"

	"log"
)

func main() {
	pool := x509.NewCertPool()
	caCertPath := "/var/run/test/ca.crt"
	caCrt, err := os.ReadFile(caCertPath)
	if err != nil {
		fmt.Println("ReadFile err:", err)
		return
	}
	pool.AppendCertsFromPEM(caCrt) //客户端添加ca证书

	//注意这里要使用证书中包含的主机名称
	conn, err := tls.Dial("tcp", "www.test.com:8888",
		&tls.Config{
			RootCAs:            pool,
			InsecureSkipVerify: true,
		})
	if err != nil {
		log.Fatalln(err.Error())
	}

	//tr := &http.Transport{
	//	TLSClientConfig:    &tls.Config{RootCAs: pool}, //客户端加载ca证书
	//	DisableCompression: true,
	//}
	defer conn.Close()

	log.Println("Client Connect To ", conn.RemoteAddr())
	status := conn.ConnectionState()
	fmt.Printf("%#v\n", status)
	buf := make([]byte, 1024)
	ticker := time.NewTicker(1 * time.Millisecond * 500)
	for {
		select {
		case <-ticker.C:
			{
				_, err = io.WriteString(conn, "hello")
				if err != nil {
					log.Fatalln(err.Error())
				}
				len, err := conn.Read(buf)
				if err != nil {
					fmt.Println(err.Error())
				} else {
					fmt.Println("Receive From Server:", string(buf[:len]))
				}
			}
		}
	}
}
