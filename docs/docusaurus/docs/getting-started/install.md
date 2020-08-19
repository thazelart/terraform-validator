---
id: install
title: Install
sidebar_label: Install
---

Prerequisite: install [Go 1.11+](https://golang.org/).
## Get the last version from releases
You can [download from here](https://github.com/thazelart/terraform-validator/releases) the binary. Move it into a directory in your `$PATH` to use it. For example:
```bash
version_latest=$(curl -s https://api.github.com/repos/thazelart/terraform-validator/releases/latest | grep -oP '"tag_name": "\K(.*)(?=")')
wget "https://github.com/thazelart/terraform-validator/releases/download/${version_latest}/terraform-validator_Linux_x86_64.tar.gz"
tar -zxvf terraform-validator_Linux_x86_64.tar.gz
chmod +x terraform-validator
sudo mv terraform-validator /usr/local/bin
```

## Install from code:
To add terraform-validator, clone this repository and then get :
then you can build it:
```
go build
```
move it into a directory in your `$PATH` to use it. For example:
```
mv terraform-validator /usr/local/bin
```
