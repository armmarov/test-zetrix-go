# Zetrix Go


Install dependencies

```
go get github.com/armmarov/zetrix-sdk-go-fork
go get github.com/golang/protobuf/proto
go get github.com/myENA/secureRandom
go get github.com/teserakt-io/golang-ed25519/edwards25519
go get google.golang.org/protobuf/reflect/protoreflect
go get google.golang.org/protobuf/runtime/protoimpl

```

Change FILL_IN_DEST_ADDRESS_HERE, FILL_IN_SOURCE_ADDRESS_HERE, and FILL_IN_PRIVATE_KEY_HERE value in zetrix.go

Run go script

```
go run zetrix.go
```