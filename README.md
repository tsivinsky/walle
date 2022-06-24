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

### Set wallpaper from internet

```bash
walle -i "https://images.unsplash.com/photo-1567447343911-56a8b455318e?ixlib=rb-1.2.1&ixid=MnwxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8&auto=format&fit=crop&w=1170&q=80"
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
