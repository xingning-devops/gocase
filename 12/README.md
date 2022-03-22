###1、安装istio
```sh
wget https://github.com/istio/istio/releases/download/1.12.0/istio-1.12.0-linux-amd64.tar.gz
tar zxf istio-1.12.0-linux-amd64.tar.gz
cd istio-1.12.0
cp bin/istioctl /usr/local/bin
istioctl install --set profile=demo -y
```


###2、创建namespace
```sh
kubectl create ns case
kubectl label ns case istio-injection=enabled
```


###3、创建安全证书
```sh
openssl req -x509 -sha256 -nodes -days 365 -newkey rsa:2048 -subj '/O=xing Inc./CN=*.xing.io' -keyout xing.io.key -out xing.io.crt
kubectl create -n istio-system secret tls xing-ssl --key=xing.io.key --cert=xing.io.crt
```

###4、安装jaeger
```sh
kubectl apply -f jaeger.yaml
kubectl edit configmap istio -n istio-system
set case.sampling=100

#修改clusterIP为nodeport
kubectl edit svc tracing -n istio-system
```

###5、部署httpserver
```sh
kubectl apply -f httpserver.yaml -n case
kubectl apply -f httpserver2.yaml -n case
kubectl apply -f httpserver3.yaml -n case
kubectl apply -f istio-httpserver.yaml -n case
```

###6、查看ingress ip
```sh
kubectl get svc -nistio-system

istio-ingressgateway   LoadBalancer   192.168.87.109
```

###7、发起100次请求
```sh
for i in {1..100} ; do curl --resolve server.xing.io:443:192.168.87.109 https://server.xing.io/case -v -k ;done
```

###8、查看链路
```sh
http://47.243.4.197:31901/
#查看链路图文件目录
```