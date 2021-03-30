Covered query

IXSCAN/COLLSCAN.   索引扫描/集合扫描 （尽量是使用 索引扫描，而不是集合扫描）

Query shape:  查询条件用了哪些字段

Selectivity   过滤性

**组合索引** 正确使用方式：ESR原则

E: (Equal) 精确匹配的字段放最前面

S:（Sort）排序条件放中间

R:(Range) 匹配的字段放最后



索引使用后台创建，避免影响当前业务。

