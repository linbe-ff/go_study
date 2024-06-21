#!/bin/bash

# 数据库相关信息
DB_USER="root"       # 数据库用户名
DB_PASSWORD="a1234"  # 数据库密码
DB_NAME="test_backcup"      # 要备份的数据库名
HOST="localhost"            # 数据库地址，默认为本地

# 备份文件存储路径
BACKUP_DIR="/usr/local/backup"

# 当前日期时间，用于构建备份文件名
TIMESTAMP=$(date +%Y%m%d_%H%M%S)

# 备份文件全路径
BACKUP_FILE="$BACKUP_DIR/db_backup_$TIMESTAMP.sql.gz"

# 日志文件
LOG_FILE="$BACKUP_DIR/backup_log.txt"

# 检查备份目录是否存在，不存在则创建
if [ ! -d "$BACKUP_DIR" ]; then
    mkdir -p "$BACKUP_DIR"
    echo "Backup directory created: $BACKUP_DIR" >> "$LOG_FILE"
fi

# 执行备份
echo "Starting backup at $(date)" >> "$LOG_FILE"
mysqldump -h$HOST -u$DB_USER -p$DB_PASSWORD $DB_NAME | gzip > "$BACKUP_FILE" 2>> "$LOG_FILE"

# 检查备份是否成功
if [ $? -eq 0 ]; then
    echo "Backup completed successfully at $(date)" >> "$LOG_FILE"
else
    echo "Backup failed at $(date)" >> "$LOG_FILE"
fi

# 可选：删除N天前的备份（例如，保留7天内的备份）
# 注意：这行命令非常强力，确保路径正确无误，避免误删重要文件
 find $BACKUP_DIR -type f -name "*.sql.gz" -mtime +7 -exec rm -f {} \; >> "$LOG_FILE" 2>&1