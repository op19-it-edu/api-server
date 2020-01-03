# APIサーバーの環境構築

## Golangの実行環境構築

1. Golangのダウンロード＆インストール

Mac
```sh
$ brew install go
```

Linux
```sh
$ sudo apt install golang-go
```

2. PATHを通す

bash
```sh
# ~/.bashrc に下記のコードを追加して保存
export GOPATH=$HOME/go
export PATH=$GOPATH/bin:$PATH
```

fish
```sh
# ~/.config/fish/config.fish に下記のコードを追加して保存
set -x GOPATH $HOME/go
set -x PATH $PATH $GOPATH/bin
```

3. 動作確認

```sh
#  下記のようにgoコマンドを叩いて、出力があれば環境構築完了
$ go version
go version go1.13.5 darwin/amd64
```

---

## MySQLの環境構築

### 準備

1. `docker`と`docker-compose`のインストール
2. `mysql`のインストール

### db_setupの使い方

1. `docker-compose.yml`があるディレクトリに移動する

```sh
$ cd db_setup
```

2. `docker-compose.yml`を元にコンテナを起動する
```sh
$ docker-compose up -d

Creating network "db_setup_default" with the default driver
Creating todotweet_db ... done
```

3. コンテナが起動できているか確認する
```sh
$ docker-compose ps

    Name                 Command             State                 Ports              
--------------------------------------------------------------------------------------
todotweet_db   docker-entrypoint.sh mysqld   Up      0.0.0.0:3306->3306/tcp, 33060/tcp
```

4. DBサーバーにアクセスできるか確認する
```sh
$ mysql -u root -p -h 127.0.0.1 -P 3306

# パスワードを要求されるので、入力（デフォルトはpassword）
Enter password:

# 下記のような出力があればOK
Welcome to the MySQL monitor.  Commands end with ; or \g.
Your MySQL connection id is 8
Server version: 8.0.18 MySQL Community Server - GPL

Copyright (c) 2000, 2019, Oracle and/or its affiliates. All rights reserved.

Oracle is a registered trademark of Oracle Corporation and/or its
affiliates. Other names may be trademarks of their respective
owners.

Type 'help;' or '\h' for help. Type '\c' to clear the current input statement.

mysql> 
```

5. `todotweet_db`というDBが作成されているか確認する

```sql
mysql> show databases;
+--------------------+
| Database           |
+--------------------+
| information_schema |
| mysql              |
| performance_schema |
| sys                |
| todotweet_db       |
+--------------------+
5 rows in set (0.01 sec)
```

6. `todotweet_db`に切り替えて、tableを確認する
```sql
/* 使用するDBを切り替える */
mysql> use todotweet_db;
Database changed
/* tableを表示する */
mysql> show tables;
Empty set (0.02 sec)
```