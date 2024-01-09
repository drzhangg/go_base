
# 写一个脚本，实现从docker hub下载任意一个镜像，然后上传到本地镜像仓库。
NAME=$1

docker pull $NAME

docker login localhost:5000

docker push localhost:5000:$NAME