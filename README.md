# api_optimal_record 
    接口优化记录

# colly
    一个爬虫库，借助实现爬取
##  实际使用
    在本次项目中，使用了Excel导入导出、并将爬到的数据写进数据库、分批插入
    有些网站需要代理才能爬取

# excel
    写了一个Excel导入导出的示例

# go-zero-log
    go-zero-log的使用记录

# redis
    couter: 记录喜欢 点赞 收藏次数
    lock: 分布式锁
    rank: 排行榜
    seckill：秒杀

# mq
    rabbitMQ使用记录

# reflect 
    反射赋值索引

# retry
    重试

# task
    定时任务

# sensitive-word
    敏感词检测
## 性能
```
1、当字符串为：今夜总会想起你夜总最淫官员，并且用strings.Repeat重复10000次的性能如下
    text      = strings.Repeat("今夜总会想起你夜总", 10000) + "最淫官员"
    
    DFA 算法性能
       BenchmarkDFAFilterAll           1000000000               0.002005 ns/op
       BenchmarkDFAFilterAll-2         1000000000               0.002516 ns/op
    
    普通遍历性能
       BenchmarkDFAFilterForr         1        1249947300 ns/op
       BenchmarkDFAFilterForr-2       1        1291353500 ns/op    
   
2、当字符串为：今夜总会想起你夜总最淫官员，并且用strings.Repeat重复10次的性能如下
        text      = strings.Repeat("今夜总会想起你夜总", 10) + "最淫官员"
        
    DFA 算法性能 （op单次时间忽略不计）
        BenchmarkDFAFilterAll-2         1000000000
    
    普通遍历性能 
        BenchmarkDFAFilterForr-2        1000000000               0.001047 ns/op
    
    
结论 
    使用DFA算法性能更好，特别是在成文本匹配中速度遥遥领先。
```

