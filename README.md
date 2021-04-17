# ✨ Generate Snowflake

Generate unique IDs. Inspired by Twitter's Snowflake system.

# 📦 Installation

-   Initialize your project (`go mod init example.com/example`)
-   Add package (`go get github.com/barbarbar338/snowflake`)

# 🤓 Usage

```go
package main

import "github.com/barbarbar338/snowflake"

func main() {
    snowflake := Snowflake {
		EPOCH: 1618606800, // your projects EPOCH
	}

    id := snowflake.Generate() // Some unique snowflake ID (eg: 6782465263234318336)
}
```

# 💻 How It Works?

```
EPOCH: 1618606800
Snowflake: 6782465263234318336

Binary: 101111000100000001010000011100011001100100000100000000000000000
Timestamp: 1617065730866
WorkerID: 1
ProcessID: 0
Increment: 0

●------------------------------------------------------------------------●

╔                                        Binary                          ╗
║10111100010000000101000001110001100110010║ ║00001║ ║00000║ ║000000000000║
╚               Timestamp                 ╝ ╚ WID ╝ ╚ PID ╝ ╚  Increment ╝
```

# 🧦 Contributing

Fell free to use GitHub's features.
