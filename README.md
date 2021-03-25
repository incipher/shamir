<div align="center">
  <img src="https://user-images.githubusercontent.com/11808903/112536517-e07bcb00-8dad-11eb-9931-10ad4fe5c1d9.png" width="200"/>

  <h1>shamir</h1>

  <p>Split and combine secrets using <a href="https://en.wikipedia.org/wiki/Shamir%27s_Secret_Sharing">Shamir's Secret Sharing</a> algorithm</p>

  <a href="https://github.com/incipher/shamir/releases/latest">
    <img src="https://img.shields.io/github/release/incipher/shamir.svg?style=for-the-badge" />
  </a>
</div>

## Usage

```
❯ shamir split -n 3 -t 2
Secret: ************************
67442ef838a57cbc3063a487d7ca861cf490b9026f5f3a41be
89ad77b930245a4a60f4698baace1ddbeaec94f0a96400a82a
9ef082cd4f3456dc4bf161460a7cd5f580ed1fd426fa3ff5d7

❯ shamir combine -t 2
Share #1: 67442ef838a57cbc3063a487d7ca861cf490b9026f5f3a41be
Share #2: 9ef082cd4f3456dc4bf161460a7cd5f580ed1fd426fa3ff5d7
SayHelloToMyLittleFriend
```

## Installation

## License

<a href="https://creativecommons.org/publicdomain/zero/1.0/">Creative Commons Zero v1.0 Universal</a>
