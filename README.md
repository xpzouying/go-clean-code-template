# go-clean-code-template

This project is a clean architecture template implemented in Go. The concept of clean architecture can be referred to through the link provided below.

- [go-clean-arch](https://github.com/xpzouying/go-clean-arch)。

## Deploy

```bash

# systemd config
cp ./examples/go-template-project.service /etc/systemd/system/go-template.service

# bin
cp ./dist/go-template-project /usr/local/bin/go-template

# config
mkdir /etc/go-template
cp ./examples/config.json /etc/go-template/config.json
```

Enable our service,

```bash
systemctl daemon-reload
systemctl enable go-template
systemctl start go-template
systemctl status go-template

# uninstall
systemctl stop go-template
systemctl disable go-template
systemctl daemon-reload
```

启动后查看对应的日志，

```bash
journalctl -u go-template -f
```
