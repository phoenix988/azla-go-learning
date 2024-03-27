# Azla - Go Language Learning

Inspired by my GTK app Azla originally written in Lua.
But it is rewritten in Go lang and it is written mainly using
The built in libary http for go

### Installation

Requirements: Go 1.22.1 or later

Install: Go

```bash
# For Arch Linux
sudo pacman -S go

# For debian/ubuntu
sudo apt install go
```

Clone the repo

```bash
git clone https://github.com/phoenix988/azla-go-learning.git

cd azla-go-learning
```

### Build the binary

```bash
# Will make a binary file in the current directory
go build .

# Or you can just run it and it will be accessible at localhost:8080
go run .
```

That's it! You can now access the app at localhost:8080
Or the ip_address_of_the_server:8080 if you host it on a server or another machine
