# CM - A Commit CLI

![cover](https://i.imgur.com/LLdRKeZ.png)

A CLI to replace your `git commit` command, so your git message can partially follow [the Conventional Changelog](https://github.com/conventional-changelog/conventional-changelog) ecosystem. And yes, it is build on top of [Go](https://go.dev)

## Install

Just run this:

```bash
go install github.com/hisamafahri/cm
```

## Usage

### Commit Changes

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

## Pushing Changes

```bash
cm p # or cm push
```

This command will push your ***current branch*** into your remote repo. 

If there is only *one* remote repo, it will push it there autommatically. If there are multiple online repo, it will prompt you to choose:

```bash
$ cm p

? Which repository you want to push?:  [Use arrows to move, type to filter]
> remote1: https://github.com/hisamafahri/remote1
  remote2: https://github.com/hisamafahri/remote2
  remote3: https://github.com/hisamafahri/remote3
```

## Author

[Hisam A Fahri](https://hisamafahri.com): [@hisamafahri](https://github.com/hisamafahri)

## License

[MIT](LICENSE)
