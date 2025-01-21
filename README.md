# go-study
English | [简体中文](https://github.com/linbe-ff/go_study/blob/main/README_zh_cn.md)


## api_optimal_record
    

## colly
    A web scraping library used for implementing crawlers.
###  Practical Usage
    In this project, Excel import/export functionality was implemented, and the scraped data was written to a database with batch inserts. 
    Some websites required proxies to scrape.

## excel
    An example of Excel import/export was written.

## go-zero-log
    Usage records for go-zero-log.

## redis
    Counter: Records like, upvote, and bookmark counts.
    Lock: Distributed lock.
    Rank: Leaderboard.
    Seckill: Flash sale functionality.

## mq
    Usage records for RabbitMQ.

## reflect
    Reflection for indexed assignment.

## retry
    Implementation of retry logic.

## task
    Scheduled task management.

## sensitive-word
    Sensitive word detection module.
### 性能
```
1、When the string is "今夜总会想起你夜总最淫官员" and repeated 10,000 times using strings.Repeat, the performance is as follows:
    text      = strings.Repeat("今夜总会想起你夜总", 10000) + "最淫官员"
    
    DFA Algorithm Performance BenchmarkDFAFilterAll
       BenchmarkDFAFilterAll           1000000000               0.002005 ns/op
       BenchmarkDFAFilterAll-2         1000000000               0.002516 ns/op
    
     Regular Loop Performance BenchmarkDFAFilterForr
       BenchmarkDFAFilterForr         1        1249947300 ns/op
       BenchmarkDFAFilterForr-2       1        1291353500 ns/op    
   
2、When the string is "今夜总会想起你夜总最淫官员" and repeated 10 times using strings.Repeat, the performance is as follows:
        text      = strings.Repeat("今夜总会想起你夜总", 10) + "最淫官员"
        
    DFA Algorithm Performance BenchmarkDFAFilterAll（op单次时间忽略不计） 
        BenchmarkDFAFilterAll-2         1000000000
    
     Regular Loop Performance BenchmarkDFAFilterForr
        BenchmarkDFAFilterForr-2        1000000000               0.001047 ns/op
    
    
Conclusion 
    The DFA algorithm performs better, especially in large text matching, significantly outperforming regular loops.
    
```

### Features Included
```
// Add new sensitive words
func (d *DFA) AddWord(word string)

// Update old sensitive words
func (d *DFA) UpdateOldWord(oldWord, newWord string)

// Delete sensitive words
func (d *DFA) DeleteWord(word string) bool

// Whether to preprocess input text
isPreprocessText bool

// Filter sensitive words from text
func (d *DFA) Filter(text string, isPreprocessText bool) string

// Check if text contains any sensitive words
func (d *DFA) Check(text string, isPreprocessText bool) error

```


