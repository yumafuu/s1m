<div align="center">
<img height="300" alt="s1m logo" src="https://raw.githubusercontent.com/yumafuu/s1m/main/.github/assets/gopher.png">

# S1M

S1M is AWS Parameter Store TUI tool.

It allows you to manage your parameters in AWS Parameter Store.

[![GitHub release](https://img.shields.io/github/v/release/yumafuu/s1m)]()

</div>

---

## Demo
<a href="https://i.gyazo.com/1556778fc4d5fa6bcf7abd33ec7d40f5.gif">
<img src="https://i.gyazo.com/1556778fc4d5fa6bcf7abd33ec7d40f5.gif" alt="Image from Gyazo" />
</a>

## Features

* List parameters
* Create new parameters
* Edit parameters
* Delete parameters
* Copy the Value, Name of parameters


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

## Icon

- Make with [GopherKon](https://www.quasilyte.dev/gopherkon/)
