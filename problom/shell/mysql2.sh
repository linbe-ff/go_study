#!/bin/bash

# 定义容器名、数据库用户、密码、数据库名和备份保存目录

# 数据库相关信息
CONTAINER_NAME="mymariadb"
DB_USER="root"       # 数据库用户名
DB_PASSWORD="a1234"  # 数据库密码
DB_NAME="test_backcup"      # 要备份的数据库名
HOST="localhost"            # 数据库地址，默认为本地
BACKUP_DIR="/usr/local/backup"

TIMESTAMP=$(date +%Y%m%d%H%M%S)

# 定义备份文件名
BACKUP_FILE=${BACKUP_DIR}/${DB_NAME}_${TIMESTAMP}.sql.gz

# 确保备份目录存在
mkdir -p ${BACKUP_DIR}

# 执行备份操作
echo "开始备份数据库 ${DB_NAME} 到 ${BACKUP_FILE}"
docker exec ${CONTAINER_NAME} sh -c "exec mysqldump -u${DB_USER} -p${DB_PASSWORD} ${DB_NAME} | gzip" > ${BACKUP_FILE}
docker exec ${CONTAINER_NAME} mysqldump -u${DB_USER} -p${DB_PASSWORD} ${DB_NAME} | gzip > ${BACKUP_FILE}

# 检查备份是否成功
if [ $? -eq 0 ]; then
    echo "备份成功"
else
    echo "备份失败"
fi

# 可选：删除N天前的备份（例如，保留7天内的备份）
# 注意：这行命令非常强力，确保路径正确无误，避免误删重要文件
 find $BACKUP_DIR -type f -name "*.sql.gz" -mtime +7 -exec rm -f {} \; >> "$LOG_FILE" 2>&1