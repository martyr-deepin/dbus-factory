#dbus-factory#

This repository is only used for saving the dbus.in.json configure xml data files

Use [dbus-generator](https://gitcafe.com/Deepin/go-dbus-generator) to generate target code (supported QML/PyQT/GoLang).
go get pkg.linuxdeepin.com/dbus-generator


NOTE:
1. ObjectName 不要和这个dbus接口提供的函数名相同，否则QML/CPP版本无法生成正确的代码
2. 尽量不要有 函数名称=Set属性名、函数名称=Get属性名  这种情况，否则golang版本会添加一个_后缀。  (QML版本无次问题)

3. golang中使用interface{}作为structure的类型(任何包含structure的类型都会直接使用interface{},也就是无法做到编译器类型检测)
