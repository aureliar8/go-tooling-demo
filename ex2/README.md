```
grpcurl -proto ./helloworl.proto -plaintext -d '{"name": "Aurélien DEROIDE"}' localhost:1234 helloworld.Greeter/SayHello 

```

```
❯ ghz --proto ./helloworl.proto --insecure -d '{"name": "Aurélien DEROIDE"}' --call helloworld.Greeter/SayHello localhost:1234 -z 20s
```
