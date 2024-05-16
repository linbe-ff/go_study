# EXPLAIN

## 从第100万行数据开始取10条数据”的需求
    WITH cte AS (
         SELECT *, ROW_NUMBER() OVER (ORDER BY id) AS row_num
         FROM table_name
     )
     SELECT * 
     FROM cte 
     WHERE row_num BETWEEN 1000001 AND 1000010

## 联合索引
    a b
    where a         能命中
    where b         不能命中
    where b and a   能命中，不如where a
###    where a or b   不能命中
    优化
    全连接
    (*** where a)
    all union
    (*** where b)

    使用IN或EXISTS：
    将OR条件转化为IN或EXISTS子查询，有时可以促使优化器做出更有效的决策。
    
    SELECT * FROM your_table
    WHERE a IN (some_value)
    OR EXISTS (SELECT 1 FROM your_table t2 WHERE t2.b = another_value AND t2.id = your_table.id)