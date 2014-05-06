#!/usr/bin/env bash

function check_json() {
    for f in $@;do
        jq . "${f}" 1>/dev/null || {
            echo "==> invalid json:" "${f}"
        }
    done
}

function check_xml() {
    for f in $@;do
        xmllint --noout "${f}" || {
            echo "==> invalid xml:" "${f}"
        }
    done
}

check_json in.json/*.json
check_xml xml/*.xml
