#!/bin/bash
# script try to run blackbox with given uniq name

usage() {
cat << EOF
usage: $0 option

This script is start a docker container to run blackbox with frpc config and unique blackbox name for frpc
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
echo current blackbox name is $BX_NAME

# main
WORK_DIR=$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )
pushd $WORK_DIR 1>/dev/null 2>&1
# create frpc.ini
rm -rf frpc.ini

cat << EOF >> frpc.ini
[common]
server_addr = 106.74.152.39
server_port = 9000
token = fuckhack\$@

[bx_${BX_NAME}]
type = stcp
sk = ${BX_NAME}
local_ip = 127.0.0.1
local_port = 9115
EOF

# start docker

CONTAINER_NAME=monitor_$BX_NAME
docker rm -f $CONTAINER_NAME
docker run -d --name $CONTAINER_NAME \
    -v $PWD/frpc.ini:/etc/blackbox_exporter/frpc.ini \
    -v $PWD/blackbox.yml:/etc/blackbox_exporter/config.yml \
    hub.htres.cn/pub/blackbox-exporter

echo blackbox container $CONTAINER_NAME started!!!
# end
popd 1>/dev/null 2>&1
