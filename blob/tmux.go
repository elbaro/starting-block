package blob

const Tmux = `#### 
#### PLUGINS
#### 
# List of plugins
set -g @plugin 'tmux-plugins/tpm'
set -g @plugin 'tmux-plugins/tmux-sensible'
set -g @plugin 'tmux-plugins/tmux-cpu'

#### 
#### Basics
####
set -g default-terminal "screen-256color"
set-option -g default-shell /bin/zsh
bind r source-file ~/.tmux.conf


# numbering from 1
set -g base-index 1
setw -g pane-base-index 1



#### 
#### Keys
#### 
bind -n M-Left select-pane -L
bind -n M-Right select-pane -R
bind -n M-Up select-pane -U
bind -n M-Down select-pane -D

#### 
#### Statusbar
#### 
set -g status-interval 5
# set -g status-right-length 200
# set -g status-right '#(battery -t) | %H:%M %m-%d-%y' 

set-option -g status-fg "#666666"
set-option -g status-bg default
set-option -g status-attr default
set-window-option -g window-status-fg "#666666"
set-window-option -g window-status-bg default
set-window-option -g window-status-attr default
set-window-option -g window-status-current-fg red
set-window-option -g window-status-current-bg default
set-window-option -g window-status-current-attr default
set-option -g message-fg white
set-option -g message-bg black
set-option -g message-attr bright
set -g status-left " "
set -g status-justify left
setw -g window-status-format         ' #W '
setw -g window-status-current-format ' #W '
set -g status-right 'CPU: #{cpu_bg_color}#{cpu_percentage}'



# Initialize TMUX plugin manager (keep this line at the very bottom of tmux.conf)
run '~/.tmux/plugins/tpm/tpm'`
