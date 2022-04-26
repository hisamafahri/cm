# CM - A Commit CLI

![cover](https://i.imgur.com/LLdRKeZ.png)

⚠️⚠️ **IMPORTANT:** ⚠️⚠️ This project is no longer actively maintained, in favor of [nit](https://github.com/hisamafahri/nit). While this project is no longer maintained, feel free to open an issue, or create a new pull request.

---

A CLI to replace your `git commit` command, so your git message can partially follow [the Conventional Changelog](https://github.com/conventional-changelog/conventional-changelog) ecosystem. And yes, it is build on top of [Go](https://go.dev)

## Install

- Download the binary file of the latest version on the [release page](https://github.com/hisamafahri/cm/releases).

- Through `go get`
  
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

## Pushing Changes

We also provided an improved `git log` for you, all you need to do is just run:

```bash
cm l # or cm log
```

![log](https://i.imgur.com/Xjgav2R.png)

## Author

[Hisam A Fahri](https://hisamafahri.com): [@hisamafahri](https://github.com/hisamafahri)

## License

[MIT](LICENSE)
