- Chạy `docker-composer up` để tạo container
- Chạy `go mod tidy` và `go mod vendor` để `add missing and remove unused modules` thêm vào go.mod file (thư viện sẽ đc tự thêm khi sử dụng trong code), chạy `go mod vendor` để `make vendored copy of dependencies`. Để sử dụng go mod cần enable `export GO111MODULE=on`. Để kiểm tra dùng `go env` (tất cả các lệnh này đều chạy trên docker). Sau khi sử dụng xong go mod, chạy lệnh `export GO111MODULE=off` để tắt go mod
- Chạy `air -c .air.toml` để live-reload khi dev xem thêm tại `https://github.com/cosmtrek/air`

Tài liệu tham khảo
- Fasthttp: `https://github.com/valyala/fasthttp` sử dụng call api
- Fiber: `https://docs.gofiber.io/` source core
- Gorm: `https://gorm.io/docs/` truy vấn mysql orm
- jsoniter: `https://github.com/json-iterator/go` parse json string
- Go: `https://tour.golang.org/` base core
- Air: `https://github.com/cosmtrek/air` hot reload khi dev code
- MongoDB: `https://github.com/mongodb/mongo-go-driver` truy vấn mongodb
- html engine: `https://curtisvermeeren.github.io/2017/09/14/Golang-Templates-Cheatsheet#actions` template engine html sử dụng 

Testing & Benchmark:
- Go testing: `https://golang.org/pkg/testing/`
- Để run test 1 API cụ thể: `go test -run GetOne` sẽ test func `TestGetOne` của file `user_test.go`
- Để test toàn bộ: `go test`
- Để Benchmark toàn dự án: `go test -bench .` để benchmark 1 file test `go test -bench ten_file` VD `go test -bench User`
- Để benchmark một func cụ thể: `go test -bench GetOne` sẽ tính điểm Benchmark của func `BenchmarkGetOne` của file `user_test.go`