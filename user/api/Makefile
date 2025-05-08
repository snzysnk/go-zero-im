build:
	GOOS=linux GOARCH=amd64 go build -o ./bin/api_user ./user.go
	export https_proxy=http://127.0.0.1:7897 http_proxy=http://127.0.0.1:7897 all_proxy=socks5://127.0.0.1:7897
	docker build -f dockerfile_go_zero_user_api  -t user-api .
