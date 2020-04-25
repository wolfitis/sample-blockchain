# sample-blockchain


to be able to work with boltDB run this command and then import the package
```
$ go get github.com/boltdb/bolt/...

```
to work with ripemd160 run this command and then import the package as in wallet.go
```
go get -u golang.org/x/crypto/ripemd160/...
```

in progress - next parts will be updated soon


to run all files except tests use:
```
go run !(*_test).go [parameters]
```

### List of useful links 

[Block hashing algorithm](https://en.bitcoin.it/wiki/Block_hashing_algorithm)

[Proof of work](https://en.bitcoin.it/wiki/Proof_of_work)

[Hashcash](https://en.bitcoin.it/wiki/Hashcash)

[Bitcoin Core: Data storage](https://en.bitcoin.it/wiki/Bitcoin_Core_0.11_(ch_2):_Data_Storage)

[boltDB](https://github.com/boltdb/bolt)

[go package: encoding/gob](https://golang.org/pkg/encoding/gob/)

[go package: flag](https://golang.org/pkg/flag/)

##### (will be updated)
