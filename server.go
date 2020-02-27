package main

import (
	m "api-server/model"
	r "api-server/route"
)

// ルーティング設定とサーバーの起動を定義

// エントリーポイント（ここから処理が始まる）
func main() {

	// migrationを実行
	m.MigrateDB()

	//  ルーティングを設定
	e := r.NewRouter()

	// サーバー起動(ここでportを指定する)
	e.Start(":1323")
}
