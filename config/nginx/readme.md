

#### 启动命令
```shell
docker run --name ngx -d -p 80:80 -p 9393:9393 -v /root/demo/nginx/nginx.conf:/etc/nginx/nginx.conf nginx:1.18-alpine


docker run --name ngx -d -p 80:80 -p 9393:9393 -v /Users/drzhang/demo/go/go_base/config/nginx/nginx.conf:/etc/nginx/nginx.conf nginx:1.18-alpine
```