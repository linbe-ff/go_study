
top 展示进程 cpu占用情况等

clear 清空命令台

systemctl enable docker // 开机自启动

docker images 查看镜像

sudo docker run -d -p 3306:3306 --name mariadb -e MYSQL_ROOT_PASSWORD=a1234 mariadb

docker rmi -f d2c94e258dcb 删除镜像
docker rm 77d063de5ecd 容器

docker exec -it mariadb bash 进入容器交互式终端

GRANT ALL PRIVILEGES ON *.* TO 'root'@'%' IDENTIFIED BY 'A123'
FLUSH PRIVILEGES

yum-config-manager --add-repo https://mirrors.aliyun.com/docker-ce/linux/centos/docker-ce.repo

docker run -itd --name redis -p 6379:6379 redis

apt-get update
apt-get install vim

chmod +x sqlbackup.sh
./backup_mysql.sh
crontab -e
0 2 * * * /path/to/your/backup_mysql.sh

docker pull gitea
docker run -d --restart=always --name=gitea -p 10022:22 -p 3000:3000 -v /var/lib/gitea:/data gitea/gitea:latest
