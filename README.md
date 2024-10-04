# Welcome to the doman repository!

Doman is a dotfiles manager that's not really usable yet because it misses some basic functionality.
I would not recommend to run it as you dotfile manager, although I won't stop you if you do so.

<br>

# Building instructions

**Requirements**

- A machine running any linux distribution (Windows or MacOS will work too probably, I just never tested it)
- Go
- Git

When you meet the requirements, execute the following command:

```shell
git clone https://github.com/superNWHG/doman.git
cd doman/cmd/doman
go install
```

If you command doesn't work after installing, you may need to add the following to your shell config (for example .zshrc or .bashrc)

```shell
export PATH=$PATH:$USER/go/bin
```

# Current features

- Creating a new dotfiles repo and init doman in it (`doman new [flag(s)]`)
- Init in an existing dotfiles repo (`doman init [flag(s)]`)
- Add a new config folder/file (`doman add [flag(s)]`)
- See all configurations that are being tracked by doman (`doman read [flag(s)]`)
- Sync configurations with the remote repository (`doman sync [flag(s)]`)
- Symlink all files that are not symlinked already (`doman link [flag(s)]`)
- Edit a dotfile entry with your preferred editory (`doman edit [flag(s)]`)

**For a list with also flags, run `doman` with no arguments**
