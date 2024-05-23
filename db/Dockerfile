# ベースイメージの指定
FROM mysql:8.0


# MySQLの設定ファイルをコピー
COPY ./my.cnf /etc/mysql/conf.d/my.cnf

# 初期設定を実行するSQLスクリプトをコピー
COPY ./init.sql /docker-entrypoint-initdb.d/init.sql

# ポートを公開
EXPOSE 3306