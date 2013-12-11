
import DBus.Com.Deepin.TestDBusGenerator 1.0
import QtQuick 2.0
import QtQuick.Controls 1.0

Item {
    Tester {
        id: "testerID"
        // path: "/com/deepin/TestDBusGenerator"
    }
    width: 400; height: 400
    TabView {
        anchors.fill  : parent
        Tab {
            title: "Tester"
            Column {

                Row {
                    Label {
                        text: "A:"
                    }
                    Text {
                        text: {
                            testerID.a = [123, 27, ["beat benas"], true, 3, 16, 162, 32,322, 64,644, 3.1415926, "/a/b/c"]
                            JSON.stringify(testerID.a)
                        }
                    }
                }
            }
        }
    }
}
