# GetSubdomain
使用在线工具获取子域名信息<br\>
这是一个用Go语言编写的、通过在线工具网站（目前支持：`https://site.ip138.com/`和`https://site.ip138.com/`）获取子域名信息的工具

# 目录结构
```
- GetSubdomain
-- main.go # 源码
-- xx.log # 日志文件
-- 20220428-out.txt # 启用结果保存模式后默认的输出文件
-- test.txt # 指定输出位置时的保存的文件
```

# Usage
```
Usage of getsubdomain:
  -o string
        -o 将结果保存到到指定位置,
                如果不指定位置但使用了该参数, 请输入default, 将保存在默认位置: [./20220428-out.txt]
                必须对参数值添加单或双引号！ (default "nil")
  -s string
        -s url后缀补充 (default "/domain.htm")
  -t string
        -t 要查询的域名 (default "baidu.com")
  -u string
        -u 使用的工具网站 (default "https://site.ip138.com/")

```

# 示例
## 使用默认`ip138.com`获取，不保存结果
```
./get_sub_domain-linux -o default -t "abc.com"
# 使用默认“https://site.ip138.com/”，查询"abc.com"，结果输出到default（即./[当前日期].txt）
```

## 使用`chaziyu.com`获取，保存结果
```
./get_sub_domain-linux -t "abc.com" -s "" -u "https://chaziyu.com/"
# 使用”https://chaziyu.com/"，查询"abc.com"，结果不保存，仅输出到控制台
# 注意这里由于默认会在url后加domain.htm，所以需要指定-s(suffix)为""
```
## 使用默认查询网站，输出到指定位置
```
./get_sub_domain-linux -t "abc.com" -o "/home/kali/temp/get_sub_domain/test.txt"
# 使用默认网址，查询"abc.com"，将结果保存到”/home/kali/temp/get_sub_domain/test.txt"
```
