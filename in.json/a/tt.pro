
TEMPLATE=lib
CONFIC += plugin
QT += qml dbus

TARGET = PolicyKit1
DESTDIR = lib

OBJECTS_DIRS = tmp
MOC_DIR = tmp

HEADERS += plugin.h dbus-authenticationAgent.h dbus-authority.h 


test.depends = lib/$(TARGET)
test.commands = (qmlscene -I . test.qml)
QMAKE_EXTRA_TARGETS += test
