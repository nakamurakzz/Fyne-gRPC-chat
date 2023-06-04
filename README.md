# Fyne-gRPC-chat
## ディレクトリ構成
```bash
.
├── client #クライアントアプリケーション
│   └── pb
├── proto #protocolbufの定義
│   └── pb
└── server #サーバアプリケーション
    └── pb
```
## protocolbufのコンパイル
```bash
make compile
```

## サーバーの起動
```bash
cd server
make start
```
