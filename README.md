# mzmt-machine
[学習用] 自動販売機のAPI

# usage

`docker-compose up -d`
してリクエストを送る

# end point

## [GET] /list
飲料のリストを取得する

[sample]
curl -H "Content-Type: application/json" localhost:8080/list | jq

## [POST] /new
新規に飲料を登録する

[sample]
curl -X POST -H "Content-Type: application/json" localhost:8080/new?name=coffee\&price=120\&amount=100

## [PATCH] /buy/:id
対象のidの飲料を購入する

[sample]
curl -X PUT -H "Content-Type: application/json" localhost:8080/buy/:id


## table

Drink
-------------
Name: string
Price: int
Amount: int
-------------
