<div align="center">
<img height="300" alt="s1m logo" src="https://raw.githubusercontent.com/yumafuu/s1m/main/.github/assets/gopher.png">

# S1M

S1M is AWS Parameter Store TUI tool.

It allows you to manage your parameters in AWS Parameter Store.

![Release](https://github.com/yumafuu/s1m/actions/workflows/release.yaml/badge.svg)
![Test](https://github.com/yumafuu/s1m/actions/workflows/test.yaml/badge.svg)


</div>

---

## Demo

[![Image from Gyazo](https://i.gyazo.com/391912839a7a9cd66a935e54a37e4c15.gif)](https://gyazo.com/391912839a7a9cd66a935e54a37e4c15)

## Features

* List parameters
* Create new parameters
* Edit parameters
* Delete parameters
* Copy the Value, Name of parameters


## Key Bindings

s1m inspired by vim key bindings.

| Key     | Description                                  |
|---------|----------------------------------------------|
| `j`     | Move down                                    |
| `k`     | Move up                                      |
| `i`     | Edit Parameter under the cursor              |
| `d`     | Delete Parameter under the cursor            |
| `o`     | Create new Parameter                         |
| `c`     | Copy the Value of Parameter under the cursor |
| `y`     | Copy the Name of Parameter under the cursor  |
| `<ESC>` | Exit from the input box                      |


## Installation

```bash
# Homebrew
brew install yumafuu/tap/s1m

# Go
go install github.com/yumafuu/s1m@latest
```

---

## Icon

- Make with [GopherKon](https://www.quasilyte.dev/gopherkon/)
