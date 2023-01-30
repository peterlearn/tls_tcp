# tls_tcp
# go 开始tcp tls server
### 参考地址 https://blog.csdn.net/wecloud1314/article/details/122409562
### https://blog.csdn.net/weixin_33862188/article/details/90255086


1，生成 CA 私钥
openssl genrsa -out ca.key 2048
2，生成CA证书
openssl req -x509 -new -nodes -key ca.key -subj “/CN=ca_host” -days 5000 -out ca.crt
3 生成服务端私钥
openssl genrsa -out server.key 2048
4 生成服务端证书认证请求
openssl req -new -key server.key -subj "/CN=www.test.com" -out server.csr
5  添加hosts文件  
 notepad C:\Windows\System32\drivers\etc\hosts
 vim /etc/hosts
 127.0.0.1 www.test.com
6 生成服务端证书
openssl x509 -req -in server.csr -CA ca.crt -CAkey ca.key -CAcreateserial -out server.crt -days 5000
