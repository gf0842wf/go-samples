### 字符串处理函数 ###
**这个是lib, 字符串处理相关的**  

- 函数doc [osc blog][http://my.oschina.net/1123581321/blog/192971](http://my.oschina.net/1123581321/blog/192971 "golang-string")  

- 判断是否以某字符串打头/结尾  
strings.HasPrefix(s string, prefix string) bool => 对应python的str.startswith   
strings.HasSuffix(s string, suffix string) bool => 对应python的str.endswith

- 字符串分割  
strings.Split(s string, sep string) []string => 对应python的str.split

- 返回子串索引  
strings.Index(s string, sub string) int => 对应python的str.index  
strings.LastIndex 最后一个匹配索引

- 字符串连接  
strings.Join(a []string, sep string) string =>对应python的str.join

- 字符串替换  
strings.Replace(s, old, new string, n int) string =>对应python的str.replace

- 转为大写/小写  
strings.ToUpper(s string) string  
strings.ToLower  
对应python的str.upper,str.lower
- 子串个数
strings.Count(s string, sep string) int  
对应python的 str.count

*以下是自己实现的*  
**安装**  
`go install zstr`  

- Partition  类似python的str.partition, 在udp包解析时很好用
- ToString   把多个字段转换为string并拼接
- IsSpace    是不是空字符
- Trim
- TrimBytes
- TrimExtraSpace