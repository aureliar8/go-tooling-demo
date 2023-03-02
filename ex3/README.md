```
❯ grpcurl -proto ./helloworl.proto -plaintext -d '{"name": "aurelien"}' localhost:1234 helloworld.Greeter/SayHello 
{
  "message": "Hello aurelien"
}
```

```
❯ ghz --proto ./helloworl.proto --insecure -d '{"name": "aurelien"}' --call helloworld.Greeter/SayHello localhost:1234 -z 20s
```


```go
 _ "net/http/pprof"
 
 go func() {
		http.ListenAndServe("localhost:6666", nil)
	}()
 ```

```
curl 'http://localhost:8080/debug/pprof/profile?seconds=15' > cpu.pprof
```
