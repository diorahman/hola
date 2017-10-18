## hola

hello, grpc!

## prepare

```
$ protoc hola/hola.proto --go_out=plugins=grpc:.
```

## play

```
$ go run main.go
$ npm install -g grpcc
$ grpcc --proto hola/hola.proto --address 127.0.0.1:8000 --insecure
$ ? What package you want to use? hola

Connecting to Hola on 127.0.0.1:8000. Available globals:

  client - the client connection to Hola
    get (GetRequest, callback) returns GetResponse

  printReply - function to easily print a server reply (alias: pr)

Hola@127.0.0.1:8000> client.get("ok", pr)
EventEmitter {}
Hola@127.0.0.1:8000>
{
  "value": "ok"
}
```

## license

MIT
