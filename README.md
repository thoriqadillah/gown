# **Gown**
Gown (Go download) is a download manager written in Go and Vue using [Wails](https://wails.io)

**## Getting Started**
### Prerequisites
- Go
- Node
- NPM

### Installation
```bash
# install the required dependencies
sudo dnf install upx nsis cmake pkgconf clang webkit2gtk4.0-devel gtk3-devel 

# install wails
go install github.com/wailsapp/wails/v2/cmd/wails@latest

# add this in the shell rc
alias wails="/home/thoriqadillah/go/bin/wails"

# verify the installation with
wails doctor
```

### Run development
```bash
wails dev
```

### Build the project
```bash
wails build -platform <x>
```

