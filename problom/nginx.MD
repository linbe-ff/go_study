package problom

docker pull nginx

# 创建挂载目录
mkdir -p /home/nginx/conf
mkdir -p /home/nginx/log
mkdir -p /home/nginx/html

# 生成容器
docker run --name nginx -p 9001:80 -d nginx
# 将容器nginx.conf文件复制到宿主机
docker cp nginx:/etc/nginx/nginx.conf /home/nginx/conf/nginx.conf
# 将容器conf.d文件夹下内容复制到宿主机
docker cp nginx:/etc/nginx/conf.d /home/nginx/conf/conf.d
# 将容器中的html文件夹复制到宿主机
docker cp nginx:/usr/share/nginx/html /home/nginx/
# 将容器中的error.log文件夹复制到宿主机
docker cp nginx:/var/log/nginx/error.log /home/nginx/error.log

# 将宿主机nginx.conf文件复制到容器
docker cp /home/nginx/conf/nginx.conf nginx:/etc/nginx/nginx.conf