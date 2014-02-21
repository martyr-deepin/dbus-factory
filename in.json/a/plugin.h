
#ifndef __DBUS_H__
#define __DBUS_H__


#include "dbus-authenticationAgent.h"
#include "dbus-authority.h"
#include <QQmlExtensionPlugin>
#include <qqml.h>

class DBusPlugin: public QQmlExtensionPlugin
{
    Q_OBJECT
        Q_PLUGIN_METADATA(IID "com.deepin.dde.daemon.DBus")

    public:
        void registerTypes(const char* uri) { 
             qmlRegisterType<Authority>(uri, 1, 0, "Authority");
             qmlRegisterType<AuthenticationAgent>(uri, 1, 0, "AuthenticationAgent");
        }
};

inline
int getTypeId(const QString& sig) {
    //TODO: this should staticly generate by xml info
    if (0) { 