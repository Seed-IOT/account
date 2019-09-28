# account

## Run service by using:
```sh
$ make run

## es 进入只读模式

```
curl -XPUT -H "Content-Type: application/json" http://localhost:9200/_all/_settings -d '{"index.blocks.read_only_allow_delete": null}'
```

```
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