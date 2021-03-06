package blob

import "runtime"

var Zshrc = func() string {
	conf := `autoload -Uz promptinit
setopt histignorealldups sharehistory
HISTSIZE=1000
SAVEHIST=1000
HISTFILE=~/.zsh_history
autoload -Uz compinit
compinit

zstyle ':completion::complete:*' use-cache on
zstyle ':completion::complete:*' cache-path ~/.zsh/cache/$HOST
zstyle ':completion:*' auto-description 'specify: %d'
zstyle ':completion:*' completer _expand _complete _correct _approximate
zstyle ':completion:*' format 'Completing %d'
zstyle ':completion:*' group-name ''
zstyle ':completion:*' menu select=2
zstyle ':completion:*:default' list-colors ${(s.:.)LS_COLORS}
zstyle ':completion:*' list-colors ''
zstyle ':completion:*' list-prompt %SAt %p: Hit TAB for more, or the character to insert%s
zstyle ':completion:*' select-prompt %SScrolling active: current selection at %p%s
zstyle ':completion:*' matcher-list '' 'm:{a-z}={A-Z}' 'm:{a-zA-Z}={A-Za-z}' 'r:|[._-]=* r:|=* l:|=*'
zstyle ':completion:*' menu select=long
zstyle ':completion:*' use-compctl false
zstyle ':completion:*' verbose true
zstyle ':completion:*:*:kill:*:processes' list-colors '=(#b) #([0-9]#)*=0=01;31'
zstyle ':completion:*:kill:*' command 'ps -u $USER -o pid,%cpu,tty,cputime,cmd'
setopt appendhistory autocd extendedglob nomatch notify

export LESS='--ignore-case --raw-control-chars'
export PAGER='less'

alias t='tmux attach -t 0 || tmux new'
alias d='docker exec -it notebook bash'
alias docker='sudo -g docker /usr/bin/docker'
alias nvidia-docker='sudo -g docker /usr/bin/nvidia-docker'

source "$HOME/.slimzsh/slim.zsh"
`
	switch runtime.GOOS {
	case "linux":
		conf += `alias grep='grep --color=auto'
alias egrep='egrep --color=auto'
alias fgrep='fgrep --color=auto'
alias l='ls -CF'
alias ls='ls --color=auto'
alias gpu='watch -n 1 nvidia-smi'
`
	case "mac":
		conf += `alias finder='open .'
alias cello='say -v cellos "di di di di di di di di di di di di di di di di di di di di di di di di di di"'
alias apps='mdfind "kMDItemAppStoreHasReceipt=1"'`
	}
	return conf
}()
