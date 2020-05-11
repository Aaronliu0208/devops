#!/bin/bash
# script that stop container

usage() {
cat << EOF
usage: $0 option

This script is stop container for blackbox with given name
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

# debug message

# main
WORK_DIR=$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )
pushd $WORK_DIR 1>/dev/null 2>&1
# start docker
CONTAINER_NAME=monitor_$BX_NAME
echo stop blackbox $BX_NAME  with container $CONTAINER_NAME
docker rm -f $CONTAINER_NAME

# end
popd 1>/dev/null 2>&1
