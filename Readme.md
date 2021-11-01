## Short demo that demonstrate how to run Rust code in Golang

### Rust project
1. Initial rust program:
src/lib.rs
```
use std::thread;

#[no_mangle]
pub extern "C" fn process() {
    let handles:Vec<_> = (0..10).map(|_|{
        thread::spawn(||{
            let mut x = 0;
            for _ in 0..5_000_000 {
                x += 1
            }
            x
        })
    }).collect();

    for h in handles {
        println!("Thread finished with count={}",
                 h.join().map_err(|_| "Could not join thread!").unwrap());
    }
    println!("Done!");
}
```
2. Cargo.toml
```
[package]
name = "embed"
version = "0.1.0"
edition = "2021"

# See more keys and their definitions at https://doc.rust-lang.org/cargo/reference/manifest.html

[lib]
crate-type = ["cdylib"]
```
3. Build rust project
```
cargo build --release
```

### Golang project.
1. Copy `libembed.so` from rust project `/target/release` to  golang `/lib` subfolder.
2. Create `process.h` inside `/lib` sunbfolder with function signature 
```
void process();
```
2. Build
```
go build -ldflags="-r /lib" main.go
```
3. Set `LD_LIBRARY_PATH` env variable.
```
LD_LIBRARY_PATH=~/go/src/github.com/rust-go/embed/embed/lib
export LD_LIBRARY_PATH
```
4. Run
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

