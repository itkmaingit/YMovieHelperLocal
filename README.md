# 概要

本リポジトリは 2024 年 5/31 まで公開されていた YMovieHelper という Web アプリケーションをご自身の PC で動かすためのリポジトリです。ライセンスの範囲内であれば、ご自身及び所属する組織の人間のみがアクセスできる形でのサーバーの公開、あるいはコードの改編を自由に行うことができます。

## 必須要件

- Windows OS 10/11
- WSL
- Docker(Desktop)
- VSCode
- git

検索をすればこれらアプリケーションのインストール方法は分かるはずです。

## 実行方法

---

**Update**

以前とはアプリケーションの起動方法が変わっています。お手数ですが、再度ご確認のほどよろしくお願いいたします。

1. 以下コマンドの実行(WSLのコマンド上)
 -> Windowsであれば、Win + Rボタンで`wsl`と入力して、出てきた画面で以下のコマンドを実行して下さい。

```bash
# 1回目
git clone https://github.com/itkmaingit/YMovieHelperLocal.git
cd YMovieHelperLocal
sudo apt install -y make
make init
```

```bash
# 2回目以降
make run
```

10分ほど待てば、<http://localhost:3000>にアクセスすればYMovieHelperを使用することができます。

※画像の再配布の禁止規約より、様々な場所で画像が表示されませんが、これはバグではありません。`YMovieHelperLocal/frontend/public`に画像を配置することで画像が表示されるようになります。

ブラウザに表示された後は、自由に YMovieHelper を無制限に使うことができます。

---

## For Developers

### Introduction

This is a repository of YMovieHelper, a support web application for creating YMovie that is called "Yukkuri Kaisetsu/Jikkyo/Douga, etc..." in Japan. Please develop it for appropriate purposes.

Please refer to [ARCHITECTURE.md](https://github.com/itkmaingit/YMovieHelperLocal/blob/master/documents/ARCHITECHTURE.md) for how to set up the environment.
