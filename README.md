# DBus Factory

**Description**: DBus Factory is a project which generates static dbus bindings for Golang and QML at build-time.

## Dependencies

### Build dependencies

- [dbus-generator](https://gitcafe.com/Deepin/go-dbus-generator) to generate the static dbus bindings

### Runtime dependencies

#### Golang Target

-  [pkg.deepin.io/lib/dbus](https://github.com/linuxdeepin/go-lib/tree/master/dbus)

#### QML Target

- Qt5DBus
- Qt5Core
- Qt5Gui
- Qt5Qml
- libstdc++6

## Getting help

Any usage issues can ask for help via

* [Gitter](https://gitter.im/orgs/linuxdeepin/rooms)
* [IRC channel](https://webchat.freenode.net/?channels=deepin)
* [Forum](https://bbs.deepin.org)
* [WiKi](http://wiki.deepin.org/)

## Getting involved

We encourage you to report issues and contribute changes

* [Contribution guide for users](http://wiki.deepin.org/index.php?title=Contribution_Guidelines_for_Users)
* [Contribution guide for developers](http://wiki.deepin.org/index.php?title=Contribution_Guidelines_for_Developers).

## Notice

1. If ObjectName was same as the method name supported by dbus service,
   QML/Cpp target can't be generated correctly.

2. Don't name your method in form like Set*$PropertyName*, Get*$PropertyName*. Otherwise the generated code will append a \_ suffix.

## License

DBus Factory is licensed under [GPLv3](LICENSE).
