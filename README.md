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
    // create a new factory with your projects EPOCH and machineID
    s := snowflake.NewFactory(1420070400000, 0)

    id := s.Generate() // Some unique snowflake ID (eg: 17447384661725548544)
}
```

# 💻 How It Works?

```
EPOCH: 1420070400000
Snowflake: 17447384661725548544

Binary: 1011110001000000010100000111000110011001000001000000000000
Timestamp: 4159708641244
MachineID: 1
Sequence: 0

●-----------------------------------------------------------------●

╔                                Binary                          ╗
║10111100010000000101000001110001100110010║ ║00001║ ║000000000000║
╚               Timestamp                 ╝ ╚ MID ╝ ╚  Sequence  ╝
```

# 🧦 Contributing

Fell free to use GitHub's features.
