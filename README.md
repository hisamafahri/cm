# CM - A Commit CLI

![cover](https://i.imgur.com/LLdRKeZ.png)

A CLI to replace your `git commit` command, so your git message can partially follow [the Conventional Changelog](https://github.com/conventional-changelog/conventional-changelog) ecosystem. And yes, it is build on top of [Go](https://go.dev)

## Install

Just run this:

```bash
go install github.com/hisamafahri/cm
```

## Usage

```bash
git add <file-or-folder> # or the scope of the file you wanna commit
cm
```

To commit all of the changes in the current directory, you can easily run:

```bash
cm -a # or cm --all

# This command will substitute:
# git add .
# cm
```

By running that command, you will add all of the changes in the *current directory* and commit it automatically. :)

## Author

[Hisam A Fahri](https://hisamafahri.com): [@hisamafahri](https://github.com/hisamafahri)

## License

[MIT](LICENSE)
