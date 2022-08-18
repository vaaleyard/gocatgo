<div align="center">

<h1> GoCatGo </h1>
GoCatGo is another pastebin tool with a super focus on transparency<br>

<br>

![lines of code](https://sloc.xyz/github/vaaleyard/gocatgo) ![Code Size](https://img.shields.io/github/languages/code-size/vaaleyard/gocatgo) [![License: MIT](https://img.shields.io/badge/License-MIT-blue.svg)](./LICENSE)

</div> 

## Requirements
It's a simple tool, you only need `curl`.

# Usage
```bash
# With a file
cat file.txt | curl -F "file=@-" gcg.sh
# or
curl -F "file=@file.txt" gcg.sh
```

```bash
# Passing any string
echo "some cool code" | curl -F "file=@-" gcg.sh
```

## Privacy, Encryption and Transparency
Everything is encrypted. See [wiki](https://github.com/vaaleyard/gocatgo/wiki).

## Contribution
See [CONTRIBUTING.md](./CONTRIBUTING.md)

## FAQ
1. How do I know the code running is the same as the repository?  
  **Answer**: I've created an URL so you can check the sha256 of current code running: gcg.sh/sha256

## License
[MIT](./LICENSE)
