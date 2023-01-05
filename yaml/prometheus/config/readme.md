#### 获取prometheus token -> sa.token
```shell
kubectl -n kube-system describe secret $(kubectl -n kube-system describe sa myprometheus | grep 'Mountable secrets' | cut -f 2- -d ":" | tr -d " ") | grep -E '^token' | cut -f2 -d':' | tr -d '\t'
```


#### 重新假装prometheus配置
```shell
curl -X POST http://150.158.87.137:9090/-/reload
```

#### 安装adapter secret流程
```shell
openssl genrsa -out serving.key 2048
openssl req -new -key serving.key -out serving.csr -subj "/CN=serving"
openssl x509 -req -in serving.csr -CA ./ca.crt -CAkey ./ca.key -CAcreateserial -out serving.crt -days 3650
kl create secret generic cm-adapter-serving-certs --from-file=serving.crt=./serving.crt --from-file=serving.key -n custom-metrics
kl get secrets -n custom-metrics

kl get apiservice | grep custom-metrics   # 验证一下

```