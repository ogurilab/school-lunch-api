# school-lunch-api

## Description

Docker と Make Command がインストールされている環境であれば、どの環境でも動作するようにしています。

## Usage

1. リポジトリをクローンします。

```bash
git clone https://github.com/ogurilab/school-lunch-api.git
```

2. サブモジュールを更新します。

```bash
git submodule update --init --recursive
```

3. .env ファイルの作成

```bash
cp app/.env.example app/.env
```

ローカルの場合は、`WIKIMEDIA_USERNAME`と`WIKIMEDIA_PASSWORD`を設定すれば OK です。
(2024/1/18 現在この開発環境では、`WIKIMEDIA_USERNAME`と`WIKIMEDIA_PASSWORD`を設定しなくても動作します。)

4. Docker コンテナを起動します。

```bash
make prod
```

5. localhost:8080/v1/swagger/ にアクセスすると Swagger UI が表示されます。

[http://localhost:8080/v1/swagger/](http://localhost:8080/v1/swagger/)
