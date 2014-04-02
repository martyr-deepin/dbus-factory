#dbus-factory#

This repository is only used for saving the dbus.in.json configure xml data files

Use [dbus-proxyer-gen](https://github.com/linuxdeepin/go-lib/tree/master/dbus/proxyer) to generate target code (supported QML/PyQT/GoLang).


NOTE:
1. ObjectName 不要和这个dbus接口提供的函数名相同，否则QML/CPP版本无法生成正确的代码
2. 尽量不要有 函数名称=Set属性名、函数名称=Get属性名  这种情况，否则golang版本会添加一个_后缀。  (QML版本无次问题)


