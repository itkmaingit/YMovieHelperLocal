# 概要

本リポジトリは 2024 年 5/31 まで公開されていた YMovieHelper という Web アプリケーションをご自身の PC で動かすためのリポジトリです。ライセンスの範囲内であれば、ご自身及び所属する組織の人間のみがアクセスできる形でのサーバーの公開、あるいはコードの改編を自由に行うことができます。

## 必須要件

- Windows OS 10/11
- WSL
- Docker(Desktop)
- VSCode
- Dev Containers(VSCode の拡張機能)
- git

検索をすればこれらアプリケーションのインストール方法は分かるはずです。

## 実行方法

---

1. 以下コマンドを実行

```:bash
git clone https://github.com/itkmaingit/YMovieHelperLocal.git
```

2. VSCode 上で `YMovieHelperLocal` のフォルダーを開く
3. F1 を押して Remote Containers の拡張機能のコマンドの中の`Dev Containers: Open Folder in Container...`というコマンドを実行し、`YMovieHelperLocal/api`を選択する
4. VSCode 上で Ctrl + Shift + N を押して、もう一度`Dev Containers: Open Folder in Container...`を実行し、次は`YMovieHelperLocal/frontend`を選択する
5. 10 分程度待ってから、PC のブラウザで`http://localhost:3000`にアクセスすれば、YMovieHelper にアクセスすることができます。

※画像の再配布の禁止規約より、様々な場所で画像が表示されませんが、これはバグではありません。`YMovieHelperLocal/frontend/public`に画像を配置することで画像が表示されるようになります。

ブラウザに表示された後は、自由に YMovieHelper を無制限に使うことができます。

※ localhost:3000 にアクセスし、20 分待っても表示されない場合は 3 で開いたウィンドウで Ctrl+@を押し、そこで`make init`、4 で開いたウィンドウで Ctrl+@を押してそこで`npm run dev`と入力してください。

---

## For Developers

### Introduction

This is a repository of YMovieHelper, a support web application for creating YMovie that is called "Yukkuri Kaisetsu/Jikkyo/Douga, etc..." in Japan. Please develop it for appropriate purposes.

Please refer to [ARCHITECTURE.md](https://github.com/itkmaingit/YMovieHelperLocal/blob/master/documents/ARCHITECHTURE.md) for how to set up the environment.
