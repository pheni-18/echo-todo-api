# echo-todo-api

## echo-todo-api は Golang, Echo, Gorm の勉強を兼ねたシンプルな WebAPI です。

## Versions

```
$ cat go.mod
```

## Docker

API の実行の前に Docker で Mysql を起動しておく

```
$ docker
$ docker-compose up -d
```

## API

実行方法

```
$ go run main.go
```

動作確認

```
$ curl -X POST localhost:1323/todo -H 'Content-Type: application/json' -d '{"title": "todo1", "description": "todo1 description", "done": 0}'
$ curl -X GET localhost:1323/todo
```
