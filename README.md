# cloudmobility-hackathon

## Development

### Golang (Backend)

#### Requirements
- Start the environment in a vscode dev environment (docker must be running in your local machine);


#### Start webserver
```
go run cmd/knfunc/main.go
```

#### Request function from git repo

```
curl -v -XPOST 127.0.0.1:8080/func_create --data '{"repo": {"url":"https://github.com/knative-sample/helloworld-go", "tag": "master", "path": "run/helloworld"}, "language": "go"}'
```

#### Get request status

```
curl -XGET http://127.0.0.1:8080/func_status/{request_id}
```
