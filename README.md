# 短链接项目

## 搭建项目的骨架

1. 建库建表

新建发号器表
```sql
CREATE TABLE `sequence`(
    `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
    `stub` varchar(1) NOT NULL,
    `timestamp` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`),
    UNIQUE KEY `idx_uniq_stub` (`stub`)
) ENGINE=MyISAM DEFAULT CHARSET=utf8mb4 COMMENT = '序号表';
```

新建长链接短链接映射表
```sql
CREATE TABLE `short_url_map`(
    `id` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键',
    `create_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `create_by` VARCHAR(64) NOT NULL DEFAULT '' COMMENT '创建者',
    `is_del` tinyint UNSIGNED NOT NULL DEFAULT '0' COMMENT '是否删除: 0正常1删除',

    `lurl` VARCHAR(2048) DEFAULT NULL COMMENT '长链接',
    `md5` CHAR(32) DEFAULT NULL COMMENT '长链接MD5',
    `surl` VARCHAR(11) DEFAULT NULL COMMENT '短链接',
    PRIMARY KEY (`id`),
    INDEX(`is_del`),
    UNIQUE(`md5`),
    UNIQUE(`surl`)
) ENGINE=INNODB DEFAULT CHARSET=utf8mb4 COMMENT = '长短链映射表';
```
2. 搭建go-zero框架的骨架

编写 `api` 文件，使用goctl命令生成代码
```go
type ConvertRequest {
    LongUrl string `json:"longUrl"`
}

type ConvertResponse{
    ShortUrl string `json:"shortUrl"`
}

type ShowRequest {
    ShortUrl string `json:"shortUrl"`
}

type ShowResponse {
    LongUrl string `json:"longUrl"`
}
service shortener-api{
    @handler ConvertHandler
    post /convert(ConvertRequest) returns(ConvertResponse)

    @handler ShowHandler
    post /:shortUrl(ShowRequest)returns(ShowResponse)
}
```

根据api文件 生成go代码
```bash 
goctl api go -api shortener.api -dir .
```
3. 根据数据表生成 model层代码

```bash
goctl model mysql datasource -url="root:103003@tcp(127.0.0.1:3306)/gorm" -table="short_url_map" -dir="./model"

goctl model mysql datasource -url="root:103003@tcp(127.0.0.1:3306)/gorm" -table="sequence" -dir="./model"
```

4. 下载项目依赖
```bash
go mod tidy
```

5. 运行项目
```bash
go run shortener.go
```
看到如下输出 表示项目成功启动啦
```bash
Starting server at 0.0.0.0:8888....
```