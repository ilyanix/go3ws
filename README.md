# go3ws
switch between worspaces in i3wm using rofi

## Prerequisites
- i3wm
- rofi

## Usage
add to your i3wm configuration file

- for switch to worspace or create new one
```
bindsym $mod+p exec "$HOME/go/src/go3ws/go3ws" 
```

- for rename current worspace
```
bindsym $mod+n exec "$HOME/go/src/go3ws/go3ws -rename"
```
