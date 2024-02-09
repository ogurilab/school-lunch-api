# school-lunch-api

Docker と Make Command がインストールされている環境であれば、どの環境でも動作するようにしています。

### Usage - API を利用する開発者向け（Line Bot 開発者向け）

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

5. localhost:8080 にアクセスすると Swagger UI が表示されます。

[http://localhost:8080](http://localhost:8080)

6. 半田市の学校給食のデータを追加する。

   - ops/docker/entrypoint/data/に init.sql を追加します。

   - make コマンドを実行して、データベースにデータを追加します。

```bash
make seed_handa
```

7. Docker のコンテナを停止する場合は、以下のコマンドを実行します。

```bash
make prod_stop
```

### Usage - API を開発する開発者向け

1. Go のインストール [https://golang.org/doc/install](https://golang.org/doc/install)

```bash
brew install go

go version
```

2. リポジトリをクローンします。

```bash
git clone https://github.com/ogurilab/school-lunch-api.git
```

3. サブモジュールを更新します。

```bash
git submodule update --init --recursive
```

4. .env ファイルの作成

```bash
cp app/.env.example app/.env
```

5. データベースの起動

```bash
make up
```

6. サーバーの起動

```bash
make start
```

7. localhost:8080 にアクセスすると Swagger UI が表示されます。

[http://localhost:8080](http://localhost:8080)

8. データベースの停止

```bash
make down
```
