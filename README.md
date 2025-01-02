# 🦕 Dina

#### Dotfiles installer not (that) awesome.

A complicated solution for trying to simplify your dotfiles. 


- [Installation](#installation)
  * [With go](#go)
  * [with bash](#bash)
- [Getting Started](#getting-started)
  * [config.toml](#config.toml)
- [Roadmap](#roadmap)

## Installation

Warning: It is still in a work in progress. It only functions as a stow replacement at the moment

##### go

```bash
go install github.com/talialy/dina@latest
```

##### bash

```bash
# coming
```

## Getting Started

### config.toml

We love building those, but now, we shouldn't

```bash
dina update
```

The best way to use `dina` is by this.
It will create the file config.toml and organize it using the folder structure. If flags like `--flatpaks` are passed, it will be added too, use --help to filter and make it more tailored to your likes.

A basic config.toml file is as this:

```toml
[[stow]]
name = "yazi"
scripts = [ ]
dependencies = [ ]

[[stow]]
name = "hypr"
dependencies = [ "blight", "fish" ]

[[stow]]
name = "kitty"
scripts = [ "fonts.sh" ]

flatpak = [ "com.spotify.Client", "org.mozilla.firefox" ]
```

The above example would use the next folder structure:

```md
# .dotfiles
config/
    hypr/
        hyprland.conf
        .scripts/
            fish.sh
        .dependencies # text file
    kitty/
        .scripts/
            fonts.sh
        kitty.conf
    yazi/
        yazi.toml
config.toml
```

## Man Page

- [Quickstart](#quickstart)

- [Basics](#basics)
  
  - [Update](#update)
  
  - [Flags](#flags)

- [update]



### Update

```bash
dina update
```

Creates or edit the config.toml file inside the dotfiles.

It's output can be determinate by the number of flags added.

##### Flags

```bash
dina update \
    --flatpak # adds all the apps installed with flatpak
    --cargo # adds binaries installed with cargo (WIP)
    --go #adds binaries installed with golang (WIP)
```

Those will change the output of the file. After updating, it is recommended to check the file itself and modified.



#### Install

```bash
dina install
```

Using the config file, it stows every directory and file inside it. Then installing the apps and dependencies.

##### Flags

```bash
dina install \
    --omit -o # Just skips if the folder is already there
    --backup -b # creates a backup of every folder found, then it installs
    --force -f # why would you do this? Deletes the folders, then installs
```

By default, it will stop at the sight of a folder already inside .config. For better use, the flags are there to make you aware



#### Backup

```bash
dina backup
```

It moves all the files inside a folder with the current date and time created at `.dina/`.

##### Commands

```bash
dina backup \
    list # shows the last 6 entries on the backup folders
    
```



```bash
dina backup --list
```



## Roadmap

🟢 Working on it. 🟡 Planned. 🔴 Unsure

- [ ] 🟢 Add external package managers
- [ ] 🟡 downloading fonts support
- [ ] 🟡 Add snap package support 
- [ ] 🟡 Support for multiple users
- [ ] 🔴 Export config to script 
