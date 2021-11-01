## Small demo that runs a rust program that spins up 10 threatds in go

1. Build
```
go build -ldflags="-r /lib" main.go
```
2. Set LD_LIBRARY_PATH
```
LD_LIBRARY_PATH=~/go/src/github.com/rust-go/embed/embed/lib
export LD_LIBRARY_PATH
```
3. Run
```
./main
Thread finished with count=5000000
Thread finished with count=5000000
Thread finished with count=5000000
Thread finished with count=5000000
Thread finished with count=5000000
Thread finished with count=5000000
Thread finished with count=5000000
Thread finished with count=5000000
Thread finished with count=5000000
Thread finished with count=5000000
Done!
```

