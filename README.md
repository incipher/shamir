<div align="center">
  <img src="https://user-images.githubusercontent.com/11808903/112536517-e07bcb00-8dad-11eb-9931-10ad4fe5c1d9.png" width="200"/>

  <h1>shamir</h1>

  <p>Split and combine secrets using <a href="https://en.wikipedia.org/wiki/Shamir%27s_Secret_Sharing">Shamir's Secret Sharing</a> algorithm</p>

  <a href="https://github.com/incipher/shamir/releases/latest">
    <img src="https://img.shields.io/github/release/incipher/shamir.svg?style=for-the-badge" />
  </a>
</div>

## Table of Contents

- [Table of Contents](#table-of-contents)
- [Description](#description)
- [Background](#background)
- [Installation](#installation)
- [Usage](#usage)
  - [Interactive](#interactive)
  - [Non-interactive](#non-interactive)
- [License](#license)

## Description

Featuring UNIX-style composability, this command-line tool facilitates splitting and combining secrets using [HashiCorp Vault's implementation](https://github.com/hashicorp/vault/blob/main/shamir/shamir.go) of [Shamir's Secret Sharing](https://en.wikipedia.org/wiki/Shamir%27s_Secret_Sharing) algorithm.

## Background

[What is Shamir's Secret Sharing algorithm?](./doc/background.md)

## Installation

| Platform     | Package manager             | Command                              |
| ------------ | --------------------------- | ------------------------------------ |
| Linux, macOS | [Homebrew](https://brew.sh) | `$ brew install incipher/tap/shamir` |

## Usage

### Interactive

![A GIF showing how to use shamir interactively](./doc/assets/interactive-usage.gif)

### Non-interactive

```
$ echo "SayHelloToMyLittleFriend" | shamir split -n 5 -k 3 > shares.txt
Secret: ************************
```

```
$ head -n 3 shares.txt | shamir combine -k 3
SayHelloToMyLittleFriend
```

## License

<a href="https://creativecommons.org/publicdomain/zero/1.0/">CC0</a>
