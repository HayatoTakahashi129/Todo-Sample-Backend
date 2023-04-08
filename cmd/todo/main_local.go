//go:build !release

package main

/**
* ローカル実行用。
* go run . で実行されたときのみ読み込まれる想定。
 */
func main() {
	app := CreateRoute()
	app.Listen(":8080")
}
