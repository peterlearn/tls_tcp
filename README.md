# tls_tcp
# go 开始tcp tls server
### 参考地址 https://blog.csdn.net/wecloud1314/article/details/122409562
### https://blog.csdn.net/weixin_33862188/article/details/90255086
### openssl windows 工具 http://slproweb.com/products/Win32OpenSSL.html

1，生成 CA 私钥 <br> 
openssl genrsa -out ca.key 2048 <br>
2 生成CA证书 <br>
openssl req -x509 -new -nodes -key ca.key -subj “/CN=www.test.com” -days 5000 -out ca.crt <br>
3 生成服务端私钥 <br>
openssl genrsa -out server.key 2048<br>
4 生成服务端证书认证请求<br>
openssl req -new -key server.key -subj "/CN=www.test.com" -out server.csr <br>
5  添加hosts文件  <br>
 notepad C:\Windows\System32\drivers\etc\hosts <br>
 vim /etc/hosts <br>
 127.0.0.1 www.test.com <br>
6 生成服务端证书 <br>
openssl x509 -req -in server.csr -CA ca.crt -CAkey ca.key -CAcreateserial -out server.crt -days 5000 <br>
