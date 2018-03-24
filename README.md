# SingleFlightでポインタを扱うときのテスト
仕事で使ってるけど、ちょっと試したかっただけ。

# Usage

```
# 起動
dep ensure
go run main.go

# 同時に叩く
curl http://localhost:1323/1
curl http://localhost:1323/2
curl http://localhost:1323/3

# 結果
もちろん最初のgoroutineで作られたEntityがベースになるので
叩かれたurlは異なるが結果がおなじになる。
```