<div align="center">

<h1> GoCatGo </h1>

<a href="https://gcg.sh">GoCatGo</a> is another simple pastebin tool.<br>

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

```bash
# Upload images
curl -F "file=@image.png" gcg.sh
```

## Alias
You can add gcg alias to your shell to make it easier to upload files:
```bash
echo 'alias gcg="curl https://gcg.sh -F file=@-"' >> ~/.bashrc

# or, easier than that

echo "$(curl https://gcg.sh/alias)" >> ~/.bashrc
```

## Transparency
See [wiki](https://github.com/vaaleyard/gocatgo/wiki).

## Contribution
See [CONTRIBUTING.md](./CONTRIBUTING.md)

## FAQ
1. How do I know the code running is the same as the repository?  
   I've created an URL so you can check the sha256 of current code running: gcg.sh/sha256
2. How can I delete a paste I've created?  
   For now, you can request me at `contact@gcg.sh`. In the future it will be implemented.

## License
[MIT](./LICENSE)
