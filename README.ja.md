

# dok（ドック）

`dok` は、`peco` によるインタラクティブな選択を使って Docker コマンドを簡単に実行できる CLI ツールです。

## 主な機能

- `dok ps` — 起動中またはすべてのコンテナを選択表示
- `dok exec` — 実行中のコンテナを選び、bash や sh などのコマンドを実行
- `dok run` — イメージやポート、コマンドを選択してコンテナを起動
- `dok rm` — 停止中を含むコンテナの削除
- `dok rmi` — イメージの削除
- `dok stop` — コンテナの停止
- `dok start` — イメージからコンテナを起動

## インストール

### 前提条件

- [Go](https://golang.org/doc/install)
- [peco](https://github.com/peco/peco) — インタラクティブな選択に必要です

```bash
brew install peco
go install github.com/inamuu/dok@latest
```

## 設定ファイル

`dok init` を実行すると、以下の設定ファイルが `~/.dok.config` に作成されます。  
追加したコマンドがexecやrunで選択できます。

```ini
[Commands]
/bin/sh
/bin/bash
ls -la
ps
```

このファイルは `dok exec` や `dok run` で使用するコマンドをカスタマイズできます。

## 必要なツール

- [Go](https://golang.org/doc/install)
- [peco](https://github.com/peco/peco)

## ライセンス

MIT