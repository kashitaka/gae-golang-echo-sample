## ローカル環境構築

### go version

```
$ go version
go version go1.8.7 darwin/amd64
```

### その他

* dep
* Cloud SDK
* google.golang.org/appengine
  * `$ go get google.golang.org/appengine`

## ローカルサーバー起動

### 依存パッケージのダウンロード

```
$ dep ensure
```

### Cloud SDKから起動

```
$ dev_appserver.py app.yaml
```

[http://localhost:8080/sample](http://localhost:8080/sample)へアクセス

※ `can't find import: "google.golang.org/appengine"` が出て依存が解決できない場合はcloud sdk内の`GOPATH`に シェルが実行するGOの `GOPATH` のリンクを貼ってあげる。GOPATHは `go env`コマンドで確認。

```
$ cd /your-cloud-sdk-path/platform/google_appengine/gopath/
$ ln -s /your-go-path/src src
```

## Cloud デプロイ

```
gcloud app deploy
```
