# Man Page

- [Quickstart](#quickstart)
  
  - [Update](#update)
  	- [Flags](#flags)
  - [Install](#install)
	- [Flags](#flags)

## Update

```bash
dina update
```

Creates or edit the config.toml file inside the dotfiles.

It's output can be determinate by the number of flags added.

### Flags

```bash
dina update \
    --flatpak # adds all the apps installed with flatpak
    --cargo # adds binaries installed with cargo (WIP)
    --go #adds binaries installed with golang (WIP)
```

Those will change the output of the file. After updating, it is recommended to check the file itself and modified.

## Install

```bash
```

Using the config file, it stows every directory and file inside it. Then installing the apps and dependencies.

### Flags

```bash
dina install \
	--type=omit,backup,force
		omit # Just skips if the folder is already there
     	backup # creates a backup of every folder found, then it installs
    	force # why would you do this? Deletes the folders, then installs
	--no-flatpaks
	--only-dotfiles
```

By default, it will stop at the sight of a folder already inside .config. For better use, the flags are there to make you aware

## Backup

```bash
dina backup
```

It moves all the files inside a folder with the current date and time created at `.dina/`.

### Commands

```bash
dina backup \
    list # shows the last 6 entries on the backup folders
```

```bash
dina backup --list
```
