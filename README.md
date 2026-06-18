# Swiss 
Swiss Army Knife for developers to handle daily basis utility requirements like `url encode`, `password generation`, `json escape`, `json to yaml conversion`
without sending your data to online services. The main motivation is to handle those in your local to clear your concern.

![](assets/swiss-layout-v2.png)

### 📖 [**Browse all commands & live examples → huseyinbabal.github.io/swiss**](https://huseyinbabal.github.io/swiss/)

## Install

```
brew tap huseyinbabal/tap
brew install swiss
```

More installation options will come

## Commands

`swiss` ships 30+ commands across JSON, YAML, XML, CSV, TOML, JWT, hashing, encoding,
text, networking and more — all running 100% locally.

👉 **Browse the full command reference with live examples at [huseyinbabal.github.io/swiss](https://huseyinbabal.github.io/swiss/)**, or run `swiss --help`.

### Pipes & stdin

Any command's value flag accepts `-` to read from stdin, so commands compose:

```bash
cat data.json | swiss json toYAML --value -
echo "hello" | swiss base64 encode -v - | swiss base64 decode -v -
```
