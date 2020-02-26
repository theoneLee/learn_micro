前置条件：
go install github.com/micro/micro/v2
etcd
mysql

curl --request POST   --url http://127.0.0.1:8080/user/login   --header 'Content-Type: application/x-www-form-urlencoded'  --data 'userName=micro&pwd=1234'
如果curl请求返回有问题，使用postman这类httpclient发送请求即可