# walle

simple program to set wallpapers in SwayWM. it simply wraps `swaybg`.

## Install

```bash
go install github.com/tsivinsky/walle@latest
```

## Usage

### Set wallpaper

```bash
walle -i ./path/to/image.png
```

walle saves path to image in its config file, so you can set it on sway startup.

### Restore wallpaper on sway startup

```bash
exec_always --no-startup-id walle --restore
```

### Set wallpaper without saving it in config

```bash
walle -i ./path/to/image.jpg -s
```

## Dev environment

### Build binary

```bash
./build.sh
```

### Install binary on system

```bash
./install.sh
```
