<div align="center">
    <h1>FJM — Fast JDK Manager</h1>
    <img src="images/logo.png" width="200" alt="Logo">
</div>

fjm is a fast, lightweight CLI tool to install and manage JDKs and their certificates across multiple shell sessions.

## Installation

Download the latest binary for your platform from the [GitHub Releases](https://github.com/Fedsam/fjm/releases) page.

### macOS / Linux

```bash
curl -Lo fjm https://github.com/your-username/fjm/releases/latest/download/fjm_Linux_x86_64.tar.gz
tar -xzf fjm_Linux_x86_64.tar.gz
chmod +x fjm
mv fjm /usr/local/bin/
```

### Windows

Download the `.zip` from the releases page, extract it, and add the binary to your `PATH`.

---

## Shell Setup

Add the following to your shell profile so fjm can manage your Java environment per session.

### bash (`~/.bashrc`)

```bash
eval "$(fjm env)"
```

### zsh (`~/.zshrc`)

```zsh
eval "$(fjm env)"
```

### PowerShell (`$PROFILE`)

```powershell
fjm env | Invoke-Expression
```

---

## Quick Start

```bash
# Install a JDK
fjm install 21

# Switch to a specific version
fjm use 21

# Check the current version
java -version
```

---

## Commands

### `fjm install <version>`

Downloads and installs the latest JDK for the given major version from Adoptium.

```bash
fjm install 21
fjm install 17
```

### `fjm use <version>`

Switches the active JDK for the current shell session.

```bash
fjm use 21
```
### `fjm cert add <domain>`

Adds a certificate to the active JDK's truststore.

```bash
fjm cert add <domain-url> --name <alias>
```
