My custom linter which checks log/slog and go.uber.org/zap logs to few rules. 

Rules: 
1. Emoji's and special characters doesn't allowed
2. Message should be only in english
3. Message should start with the lower case letter
4. Sensitive data doesn't allowed


To user linter 
```sh
cd cmd
go build -o mylinter
mylinter yourgofile.go
```

To use linter as golangci-lint plugin
```sh
cd ./plugin
golangci-lint custom -v
./custom-gcl yourgofile.go
```