### Starting the server

The following path contains the starter for the server based on a configuration file. You must specify a config file like the example above the same folder for the parameter `-c`.
```
cmd/serverConf/main.go
```

From the root:
```
go run cmd/serverConf/main.go -c cmd/serverConf/server.yml
```

The other starters may be used for testing of features as an isolated way.
```
cmd/lexer/main.go
cmd/waf/main.go
cmd/parser/main.go
cmd/server/main.go
```