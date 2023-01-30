# tls_tcp
# go 开始tcp tls server
### 参考地址
https://cloud.tencent.com/developer/article/2054716 <br>
https://blog.csdn.net/wecloud1314/article/details/122409562 <br>
 https://blog.csdn.net/weixin_33862188/article/details/90255086 <br>
### openssl windows 工具 http://slproweb.com/products/Win32OpenSSL.html

1，生成 CA 私钥 <br> 
openssl genrsa -out ca.key 2048 <br>
2 生成CA证书 <br>
openssl req -x509 -new -nodes -key ca.key -subj "/CN=www.test.com" -days 5000 -out ca.crt <br>
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



## 客户端 && 服务器 开启双向认证 tls_http下的目录
&emsp;HTTPS的一种简单实现<br>
生成密钥、证书<br>
第一步，为服务器端和客户端准备公钥、私钥<br>
 生成服务器端私钥<br>
 openssl genrsa -out server.key 1024<br>
 生成服务器端公钥<br>
 openssl rsa -in server.key -pubout -out server.pem<br>
 生成客户端私钥<br>
 openssl genrsa -out client.key 1024<br>
 生成客户端公钥<br>
openssl rsa -in client.key -pubout -out client.pem<br>
第二步，生成 CA 证书 <br>
生成 CA 私钥<br>
openssl genrsa -out ca.key 1024<br>
X.509 Certificate Signing Request (CSR) Management.<br>
openssl req -new -key ca.key -out ca.csr<br>
X.509 Certificate Data Management.<br>
openssl x509 -req -in ca.csr -signkey ca.key -out ca.crt<br>
在执行第二步时<br>
Common Name (e.g. server FQDN or YOUR name) []: 这一项，是最后可以访问的域名，为了方便测试，写成 test.com<br>
如果想通过IP直接访问而不想通过域名，在创建服务端证书时：用如下语句：<br>
openssl genrsa -out server.key 2048<br>
openssl req -new -key server.key -subj "/CN=192.168.1.10" -out server.csr<br>
echo subjectAltName = IP:192.168.1.10 > extfile.cnf<br>
openssl x509 -req -in server.csr -CA ca.crt -CAkey ca.key -CAcreateserial -extfile extfile.cnf -out server.crt -days 5000<br>
第三步，生成服务器端证书和客户端证书<br>
服务器端需要向 CA 机构申请签名证书，在申请签名证书之前依然是创建自己的 CSR 文件<br>
openssl req -new -key server.key -out server.csr<br>
向自己的 CA 机构申请证书，签名过程需要 CA 的证书和私钥参与，最终颁发一个带有 CA 签名的证书<br>
openssl x509 -req -CA ca.crt -CAkey ca.key -CAcreateserial -in server.csr -out server.crt<br>
client 端<br>
openssl req -new -key client.key -out client.csr<br>
client 端到 CA 签名<br>
openssl x509 -req -CA ca.crt -CAkey ca.key -CAcreateserial -in client.csr -out client.crt<br>
第四步，客户端机器添加test.com域名<br>
sudo vim /etc/hosts<br>
添加服务端的ip信息对应的域名：192.168.1.100  test.com<br>






