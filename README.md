# free-osx

> unix `free` on osx in golang

## Install / Usage

Run `go get github.com/adamdecaf/free-osx` and then you can `go install` to run `free-osx`.

```bash
~/go/src/github.com/adamdecaf/free-osx $ ./free-osx -h
6.17GB
```

## Inspiration

I often want to just see the [free / cached ram on my mac book pro](https://apple.stackexchange.com/questions/4286/is-there-a-mac-os-x-terminal-version-of-the-free-command-in-linux-systems), but am often unable to without using `top` or `htop`. There's a handy [bash](https://github.com/vigo/dotfiles-universal/blob/master/prompts/free_memory) and [ruby](https://github.com/vigo/dotfiles-universal/blob/master/prompts%2Ffree_memory.rb) scripts written, but this is a port to golang.
