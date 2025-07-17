# Monime CLI

![GitHub release (latest by date)](https://img.shields.io/github/v/release/monimesl/monime-cli)

**Power up your integration game with terminal-driven control, testing, and management.**

> **⚠️ Alpha Release** — This project is in active development.  
> Features may be incomplete, and APIs may change without notice.

## Introduction

The Monime CLI is your command-line companion for building, testing, and managing your Monime integrations with unparalleled efficiency. Designed for developers and power users, it brings the full power of Monime's APIs and services directly to your terminal.

This CLI also manages two complementary GUI applications:

* **USSD Simulator**, providing a visual interface for testing payment codes and USSD flows locally.
* **Webhook Inspector**, a tool to easily inspect and debug incoming webhooks during your development.

## Installation

### MacOS

The Monime CLI can be easily installed via Homebrew:

```bash
brew tap monimesl/monime-cli
brew install monimesl/monime-cli/monime
````

### Linux

Follow these steps to install Monime CLI on Linux:

1. Download the latest Linux `.tar.gz` file for your system architecture from our [GitHub Releases](https://github.com/monimesl/monime-cli/releases).
2. Extract the archive:
   ```bash
   tar -xvf monime_cli_x.x.x_linux_amd64.tar.gz
3. Move the extracted ./monime binary to your system's PATH
   ```bash
   sudo mv ./monime /usr/local/bin
   

### Windows

Follow these steps to install Monime CLI on Windows:

1. Download the latest Linux `.tar.gz` file for your system architecture from our [GitHub Releases](https://github.com/monimesl/monime-cli/releases).
2. Extract the archive:
   ```bash
   tar -xvf monime_cli_x.x.x_windows_amd64.tar.gz
3. Add the path to the extracted monime.exe to your **Path** environment variable.

**Note on Windows:** Our Windows binaries are currently **unsigned**. When you download and attempt to run an unsigned executable, Windows SmartScreen may present a warning such as "Windows protected your PC" or a similar security alert. To run the application despite this, please click on **"More info"** within the SmartScreen pop-up, and then select **"Run anyway"**.

## Usage

Once installed, you can start using the Monime CLI:

```bash
# Example: Get the CLI help info
monime --help
```

```bash
# Example: Check the CLI version
monime --version
```

## License

Copyright © Monime Ltd. All rights reserved.

This project is licensed under the [Apache 2.0 license](LICENSE).