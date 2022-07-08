<div align="center">
    <img height=200 src="https://user-images.githubusercontent.com/57289763/177349765-887dd049-f5cf-440f-9a57-e04161019759.png" alt="七輪の上で焼かれたお餅">
</div>

<h1 align="center">Omochi 😊</h1>

<p align="center"><strong>Full text search engine from scratch by Golangʕ◔ϖ◔ʔ (Just a toy)</strong></p>

## ✨ Features

- Omochi is an inverted index based search engine by Golang.
- If indexed correctly, any document can be searched.
- You can search documents from RESTful API.
- Supported language: English, Japanese.
<div align="center">
 <img width="673" alt="スクリーンショット 2022-07-08 11 08 15" src="https://user-images.githubusercontent.com/57289763/177902420-998af3a1-1387-4943-8332-eea8fe90aca2.png">
</div>

<!-- https://app.diagrams.net/#G15YMpFAnxuCpX0XI_Yjx7IuFPEc-k9Zf9 -->


## 📍 Requirements

- [Golang](https://golang.org/) 1.18+
- [Docker](https://www.docker.com/) 20.10+

## 📦 Setup

#### **Create network**

Create [docker network](https://docs.docker.jp/engine/reference/commandline/network_create.html)(omochi_network) by:
```
$ docker network create omochi_network
```

#### **Database migration**

Omochi uses [MariaDB](https://mariadb.org/) for storing Inverted Indexes & Documents, and [Ent](https://entgo.io/) for ORM.

For database migration, connect docker container shell by:
```
$ docker-compose run api bash
```

Then, running database migration by:

```
$ go run ./cmd/migrate/migrate.go 
```

####  **Seed data**

To try search engine, this project provides two datasets as samples in TSV Format. 

The dataset for English is a **Movie title dataset**, and the dataset for Japanese is a **Doraemon comic title dataset**.

At first, connect docker container shell by:

```
$ docker-compose run api bash
```

Then, seed data by:

```
$ go run {path to seed.go}
```

If you initialize with a Japanese dataset, `{path to seed.go}` should be `./cmd/seeds/en/seed.go `. On the other hand, for English, `./cmd/seeds/eng/seed.go `.


## 🏇 Start Application

After completing setup, you can start application by running:

```
$ docker-compose up
```

This app starts a RESTful API and listens on port 8081 for connections

## 🌎 How to use & Demo

After seeding data , you can search documents by send GET request to `/v1/document/search` . 

Query parameters are as follow:

- **`"keywords"`**: Keywords to search. If there are multiple search terms, specify them separated by commas like `"hoge,fuga,piyo"`
- **`"mode"`**: Search mode. The search modes that can be specified are `"And"` and `"Or"`

#### Demo

- **Doraemon comic title dataset**

After data seeding by **Doraemon comic title dataset**, you can search documents which include "ドラえもん" by: 
```
$ curl "http://localhost:8081/v1/document/search?keywords=ドラえもん" | jq . 
{
  "documents": [
    {
      "id": 12054,
      "content": "ドラえもんの歌",
      "tokenized_content": [
        "ドラえもん",
        "歌"
      ],
      "created_at": "2022-07-08T12:59:49+09:00",
      "updated_at": "2022-07-08T12:59:49+09:00"
    },
    {
      "id": 11992,
      "content": "恋するドラえもん",
      "tokenized_content": [
        "恋する",
        "ドラえもん"
      ],
      "created_at": "2022-07-08T12:59:48+09:00",
      "updated_at": "2022-07-08T12:59:48+09:00"
    },
    {
      "id": 11230,
      "content": "ドラえもん登場！",
      "tokenized_content": [
        "ドラえもん",
        "登場"
      ],
      "created_at": "2022-07-08T12:59:44+09:00",
      "updated_at": "2022-07-08T12:59:44+09:00"
    },
    ... 
```

- **Movie title dataset**

After data seeding by **Movie title dataset**, you can search documents which include "toy" and "story" by: 
```
$ curl "http://localhost:8081/v1/document/search?keywords=toy,story&mode=And" | jq .
{
  "documents": [
    {
      "id": 1,
      "content": "Toy Story",
      "tokenized_content": [
        "toy",
        "story"
      ],
      "created_at": "2022-07-08T13:49:24+09:00",
      "updated_at": "2022-07-08T13:49:24+09:00"
    },
    {
      "id": 39,
      "content": "Toy Story of Terror!",
      "tokenized_content": [
        "toy",
        "story",
        "terror"
      ],
      "created_at": "2022-07-08T13:49:34+09:00",
      "updated_at": "2022-07-08T13:49:34+09:00"
    },
    {
      "id": 83,
      "content": "Toy Story That Time Forgot",
      "tokenized_content": [
        "toy",
        "story",
        "time",
        "forgot"
      ],
      "created_at": "2022-07-08T13:49:53+09:00",
      "updated_at": "2022-07-08T13:49:53+09:00"
    },
    {
      "id": 213,
      "content": "Toy Story 2",
      "tokenized_content": [
        "toy",
        "story"
      ],
      "created_at": "2022-07-08T13:50:35+09:00",
      "updated_at": "2022-07-08T13:50:35+09:00"
    },
    {
      "id": 352,
      "content": "Toy Story 3",
      "tokenized_content": [
        "toy",
        "story"
      ],
      "created_at": "2022-07-08T13:51:23+09:00",
      "updated_at": "2022-07-08T13:51:23+09:00"
    }
  ]
}
```

## 📚 Reference

#### Dataset

- Fujiko.F.Fujio,Doraemon(Tentomushi Comics) 1~45, Shogakukan , 1974～1996
- ROUNAK BANIK."The Movies Dataset".kaggle.https://www.kaggle.com/datasets/rounakbanik/the-movies-dataset. Accessed on 07/08

#### Book

- [Information Retrieval: Implementing and Evaluating Search Engines](https://www.amazon.co.jp/Information-Retrieval-Implementing-Evaluating-Engines/dp/0262026511)
- [情報検索アルゴリズム](https://www.amazon.co.jp/%E6%83%85%E5%A0%B1%E6%A4%9C%E7%B4%A2%E3%82%A2%E3%83%AB%E3%82%B4%E3%83%AA%E3%82%BA%E3%83%A0-%E5%8C%97-%E7%A0%94%E4%BA%8C/dp/4320120361/ref=pd_lpo_3?pd_rd_i=4320120361&psc=1)
- [Pythonではじめる 情報検索プログラミング](https://www.amazon.co.jp/Python%E3%81%A7%E3%81%AF%E3%81%98%E3%82%81%E3%82%8B-%E6%83%85%E5%A0%B1%E6%A4%9C%E7%B4%A2%E3%83%97%E3%83%AD%E3%82%B0%E3%83%A9%E3%83%9F%E3%83%B3%E3%82%B0-%E4%BD%90%E8%97%A4-%E9%80%B2%E4%B9%9F/dp/4627818610)
- [WEB+DB PRESS Vol.126. 特集 Goで作って学ぶ検索エンジン](https://www.amazon.co.jp/WEB-DB-PRESS-Vol-126-%E7%9C%9F%E5%A3%81/dp/4297125390)
- [検索エンジン自作入門 ~手を動かしながら見渡す検索の舞台裏](https://www.amazon.co.jp/%E6%A4%9C%E7%B4%A2%E3%82%A8%E3%83%B3%E3%82%B8%E3%83%B3%E8%87%AA%E4%BD%9C%E5%85%A5%E9%96%80-%E6%89%8B%E3%82%92%E5%8B%95%E3%81%8B%E3%81%97%E3%81%AA%E3%81%8C%E3%82%89%E8%A6%8B%E6%B8%A1%E3%81%99%E6%A4%9C%E7%B4%A2%E3%81%AE%E8%88%9E%E5%8F%B0%E8%A3%8F-%E5%B1%B1%E7%94%B0-%E6%B5%A9%E4%B9%8B/dp/4774167533)


## 🧑‍💻 License

[MIT](https://github.com/YadaYuki/omochi/blob/yadayuki/add-readme/LICENSE)
