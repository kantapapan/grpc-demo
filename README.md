# grpc-demo

gRPC は Google が開発しているRPCライブラリとFWで、通信層は HTTP/2 を介して行われます。

試しにGolangによるgRPCサーバー&クライアントを同一環境に立ててgRPC通信を行ってみる。


# OS

    CentOS Linux release 7.6.1810 (Core)

# Install


##  Go gRPCライブラリ

    go get -u google.golang.org/grpc


##  protocol buffer コンパイラ

- centos7用のrpmを取得してインストールします。
- protoファイルからコード生成をするコンパイラ(protoc)

    wget http://mirror.centos.org/centos/7/os/x86_64/Packages/emacs-filesystem-24.3-22.el7.noarch.rpm
    wget https://cbs.centos.org/kojifiles/packages/protobuf/3.6.1/4.el7/x86_64/protobuf-3.6.1-4.el7.x86_64.rpm
    wget https://cbs.centos.org/kojifiles/packages/protobuf/3.6.1/4.el7/x86_64/protobuf-compiler-3.6.1-4.el7.x86_64.rpm

    rpm -ivh *.rpm


## Protocol Buffersのプラグインをインストール

    go get github.com/golang/protobuf/protoc-gen-go


## gRPCミドルウェア

    go get -u github.com/grpc-ecosystem/go-grpc-middleware
    go get -u github.com/sirupsen/logrus


## demoソースの構成

    grpc-demo/
    ├── README.md
    ├── client.go
    ├── go.mod
    ├── go.sum
    ├── main.go
    ├── pb
    │   └── cat.pb.go
    ├── proto
    │   └── cat.proto
    └── service
        └── cat.go


## サーバの実装

IDLを使って `proto`配下の `.proto`にAPI定義を記述します。

本家ドキュメント
https://developers.google.com/protocol-buffers/docs/proto3#scalar

## .proto ファイルからserver、client,interface等のコードを生成する

```console
$ pwd #~/go/src/github.com/kantapapan/grpc-demo
$ protoc  -I ./proto --go_out=plugins=grpc:./pb cat.proto

```
実行後 `pb`配下にcat.pb.goが生成されています。


## サービス実装

作られた*.pb.goのinterfaceを満たすように実装します。
`service` 配下のcat.go を参照してください。

## gRPC Server 実装

main.goを参照してください。

## gRPC Client実装

client.goを参照してください。


## go run で確認
一通り必要なものができたので、実際にgo runして動作を確認する。


- grpcサーバー起動

```console
go run main.go
```

- grpcクライアントからテストしてみる

```console
go run client.go
```

- 結果

    result:&cat.MyCatResponse{Name:"mike", Kind:"Norwegian Forest Cat", XXX_NoUnkeyedLiteral:struct {}{}, XXX_unrecognized:[]uint8(nil), XXX_sizecache:0} 
    error::<nil> 

以上で、ローカル環境でgRPCの動作確認が完了


## 参考

https://qiita.com/marnie_ms4/items/4582a1a0db363fe246f3
https://github.com/golang/protobuf
https://cbs.centos.org/koji/buildinfo?buildID=26113
https://centos.pkgs.org/7/centos-x86_64/emacs-filesystem-24.3-22.el7.noarch.rpm.html




