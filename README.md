# Clean Architecture API Server

go 言語で Clean Architecture な API Server を作って、
クリーンアーキテクチャについて理解を深めるリポジトリ

## Execution help

* main

```shell
$ go run src/app/main.go
```

or

```
$ make
```

* docker container

```shell
$ make compose/up
```

## endpoint

* user の登録 : POST `/users`

```shell
$ curl -i -H "Accept: application/json" -H "Content-type: application/json" -X POST -d '{"FirstName": "suzuki", "LastName": "taro"}' localhost:8080/users
```

* user の一覧取得 : GET `/users`

```shell
$ curl localhost:8080/users
```

## Links
* [Clean ArchitectureでAPI Serverを構築してみる](https://qiita.com/hirotakan/items/698c1f5773a3cca6193e)
