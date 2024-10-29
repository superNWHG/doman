# Welcome to the doman repository!

Doman is a personal project, please do not expect working code or it having the features you would expect although I use it myself.

# Building instructions

**Requirements**
- A machine running any linux distribution (Windows or MacOS will work too probably, I just never tested it)
- Go
- Git
- Make

First clone the repository:
```shell
git clone https://github.com/superNWHG/doman.git
cd doman
```

To build a doman binary, run:
```shell
make build
```

To install doman, run:
```shell
make install
```

If you command doesn't work after installing, you may need to add the following to your shell config (for example .zshrc or .bashrc)

```shell
export PATH=$PATH:$HOME/go/bin
```

# Current features

- Creating a new dotfiles repo and init doman in it (`doman new [flag(s)]`)
- Init in an existing dotfiles repo (`doman init [flag(s)]`)
- Add a new config folder/file (`doman add [flag(s)]`)
- See all configurations that are being tracked by doman (`doman read [flag(s)]`)
- Sync configurations with the remote repository (`doman sync [flag(s)]`)
- Symlink all files that are not symlinked already (`doman link [flag(s)]`)
- Edit a dotfile entry with your preferred editor (`doman edit [flag(s)]`)

**For a list with also flags, run `doman` with no arguments**
