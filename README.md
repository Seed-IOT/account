# 用户中心
> 用户中心，用于管理用户，提供一系列服务，例如SSO、OAuth等

## 依赖

* ELK（loadsh暂时不需要）
* mysql
* redis

## 启动
$ make run

## es 进入只读模式
> 使用一下命令可以解决，或者配置更多的内存和硬盘

```

curl -XPUT -H "Content-Type: application/json" http://localhost:9200/_all/_settings -d '{"index.blocks.read_only_allow_delete": null}'

curl -XPUT -H 'Content-Type: application/json' http://localhost:9200/_cluster/settings -d '
{
    "persistent" : {
        "cluster.routing.allocation.disk.threshold_enabled" : false
    }
}'

curl -XPUT -H 'Content-Type: application/json' http://localhost:9200/_all/_settings -d '
{
    "index.blocks.read_only_allow_delete": null
}'

```

## 接口文档

安装swagger

```
go get github.com/swaggo/swag/cmd/swag
```
运行

```
swag init
```

地址
```
http://localhost:8080/swagger/index.html
```