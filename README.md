ginでのセッションテスト

```
go run main.go
```
ブラウザで`localhost:9001/login`にアクセス<br>
ユーザ名（なんでもいい）を入力し、ログインを押下<br>
ログイン後のメニュー画面でdevtoolを開き、<br>
初回：`ResponseHeaders`の`Set-Cookie`に`session-id`が発行されていることを確認<br>
2回目以降：`RequestHeaders`の`Cookie`に`session-id`が発行されていることを確認<br>
