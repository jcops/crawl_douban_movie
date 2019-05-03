# crawl_douban_movie
Go爬取豆瓣电影



## How to run

### Required

- Mysql
- Redis

### Ready

Create a **go database** and import [SQL](https://github.com/jcops/crawl_douban_movie/blob/master/sql/movie.sql)

创建一个 test库,然后导入sql,创建表！

### Conf

You should modify `conf/app.ini`

```
[database]
Type = mysql
User = root
Password =
Host = 127.0.0.1:3306
Name = go
```

## Installation

```
yum install go -y 
export GOPROXY=https://goproxy.io
go get github.com/jcops/crawl_douban_movie
cd $GOPATH/src/github.com/jcops/crawl_douban_movie
go build 
```

### 效果

![1556892498059](C:\Users\niejichao\AppData\Roaming\Typora\typora-user-images\1556892498059.png)

![1556892472701](C:\Users\niejichao\AppData\Roaming\Typora\typora-user-images\1556892472701.png)