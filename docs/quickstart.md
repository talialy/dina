# Quickstart

set everything in less that 5 minutes.

If you don't have it, create a `.dots` folder inside your home directory (or virtually anywhere you want).

```bash
mkdir ~/.dots
```

Make a filetree for dina to read inside `~/.dots`

```bash
mkdir ~/.dots/{config,scripts}
```

Copy all of your desired directories from `.config` to the newly created folder.

```bash
cp -r ~/.config/hypr ~/dots/config # Or just use a GUI if you want
```

Install dina and run updates

```bash
go install github.com/talialy/dina@latest && dina update
```

And last, use dina to install everything (backing up just in case)

```bash
dina install --backup
```

And that's all you need to use dina :)
In any case, you can go to the [man-page](./man-page.md)
