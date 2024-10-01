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
``` shell
git clone https://github.com/superNWHG/doman.git
cd doman
go install
```

If you command doesn't work after installing, you may need to add the following to your shell config (for example .zshrc or .bashrc)
``` shell
export PATH=$PATH:$USER/go/bin
```
