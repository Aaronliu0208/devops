#!/bin/bash
# script that stop visitor

usage() {
cat << EOF
usage: $0 option

This script is stop blackbox visitor
Options:
    -h  Show the help and exit
    -n  VALUE  Name for frpc client 
EOF
}

declare BX_NAME
while getopts ":n:h" optKey; do
    case "$optKey" in
        n)
            BX_NAME=$OPTARG
			;;
        h|*)
			usage
			;;
    esac
done

if [[ -z $BX_NAME ]]; then
    usage
    exit 1
fi


# main
WORK_DIR=$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )
pushd $WORK_DIR 1>/dev/null 2>&1
# start docker
echo stop blackbox vistor $BX_NAME
ps aux | grep frpc_${BX_NAME}.ini | grep -v grep | awk '{print $2}' | xargs kill -9 1>/dev/null 2>&1

# end
popd 1>/dev/null 2>&1
