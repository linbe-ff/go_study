
# 安装到docker
```
docker pull gitea
docker run -d --restart=always --name=gitea -p 10022:22 -p 3000:3000 -v /var/lib/gitea:/data gitea/gitea:latest

```