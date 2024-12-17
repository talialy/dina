# 🦕 Dina

#### Dotfiles installer not (that) awesome.
A complicated solution for simplifying your dotfiles. 

**Set it up in less than 5 minutes or get your money back!**

- [Installation](#installation)
	* [With go](#go)
	* [with bash](#bash)
- [Getting Started](#getting-started)
    * [config.toml](#config.toml)
- [Roadmap](#roadmap)

## Installation
##### go
```bash
go install https://github.com/talialy/dina
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
[stow]
[stow.yazi]
[stow.hypr]
scripts = [ "fish.sh" ]
dependencies = ["fish", "blight"]
[stow.kitty]
scripts = [ "fonts.sh" ]
flatpak = [
    "com.spotify.Client",
    "org.mozilla.firefox"
]
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


## Roadmap
🟢 Working on it. 🟡 Planned. 🔴 Unsure

- [ ] 🟢 Add external package managers
- [ ] 🟡 downloading fonts support
- [ ] 🟡 Add snap package support 
- [ ] 🟡 Support for multiple users
- [ ] 🔴 Export config to script 

