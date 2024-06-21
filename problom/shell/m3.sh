#!/bin/bash

#容器ID
container_id="638b1ac8b54d"
#登录用户名
mysql_user="root"
#登录密码(注意 如果密码包含特殊符号 前面要用'\')
mysql_password="a1234"
mysql_port="3306"
#备份的数据库名
mysql_database="test_backcup"
# 备份文件存放地址(根据实际情况填写)
backup_dir="/usr/local/backup"
# 是否删除过期数据
expire_backup_delete="true"
#过期天数
expire_days=3
backup_time=`date +%Y%m%d%H%M`

# 备份指定数据库中数据
docker exec  $container_id mysqldump  -P$mysql_port -u$mysql_user -p$mysql_password  $mysql_database > $backup_dir/bak-$mysql_database-$backup_time.sql
docker exec  638b1ac8b54d mysqldump  -P$3306 -u$root -p$a1234  $test_backcup > $/usr/local/backup/test.sql

# 删除过期数据
if [ "$expire_backup_delete" == "true" -a "$backup_dir"!="" ];then
        `find $backup_dir/ -type f -mtime +$expire_days | xargs rm -rf`
        echo "Expired backup data delete complete!"
fi