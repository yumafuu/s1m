<div align="center">
<img height="180" alt="s1m logo" src="https://raw.githubusercontent.com/yumafuu/s1m/main/.github/assets/gopher.png">

# S1M

S1M is AWS Parameter Store TUI tool.

It allows you to manage your parameters in AWS Parameter Store.

[![GitHub release](https://img.shields.io/github/v/release/yumafuu/s1m)]()]

</div>

---

## Demo
<a href="https://gyazo.com/1eedb6e565cf6559bb5e175579674f46"><img src="https://i.gyazo.com/1eedb6e565cf6559bb5e175579674f46.gif" alt="Image from Gyazo" width="1000"/></a>

## Features

* Show all parameters in AWS Parameter Store
* Create new parameters
* Edit parameters
* Delete parameters
* Copy the value, name of parameters


## Key Bindings

s1m inspired by vim key bindings.

| Key     | Description                                 |
|---------|---------------------------------------------|
| `j`     | Move down                                   |
| `k`     | Move up                                     |
| `i`     | Edit Parameter under the cursor             |
| `d`     | Delete Parameter under the cursor           |
| `o`     | Create new Parameter                        |
| `c`     | Copy the Value under the cursor             |
| `y`     | Copy the Name of Parameter under the cursor |
| `<ESC>` | Exit from the input box                     |


## Installation

```bash
# Homebrew
brew install YumaFuu/tap/s1m

# Go
go install github.com/YumaFuu/s1m@latest
```

---

## License
s1m is released under the MIT license. See [LICENSE](

## Icon

- [GopherKon](https://www.quasilyte.dev/gopherkon/)
