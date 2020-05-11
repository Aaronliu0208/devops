#!/bin/bash
# generate visitor for frpc.ini with given name
usage() {
cat << EOF
usage: $0 option

This script is start a frpc process to proxy blackbox instance with given name and port
Options:
    -h  Show the help and exit
    -n  VALUE  Name for frpc client 
    -p  VALUE  Port for frpc local port
EOF
}

declare BX_NAME
declare BX_PORT
while getopts ":n:p:h" optKey; do
    case "$optKey" in
        n)
            BX_NAME=$OPTARG
			;;
        p)
            BX_PORT=$OPTARG
			;;
        h|*)
			usage
			;;
    esac
done

if [[ -z $BX_PORT ]]; then
    usage
    exit 1
fi

if [[ -z $BX_NAME ]]; then
    usage
    exit 1
fi

# debug message
echo current blackbox name is $BX_NAME

rm -rf frpc_${BX_NAME}.ini

cat << EOF >> frpc_${BX_NAME}.ini
[common]
server_addr = 106.74.152.39
server_port = 9000
token = fuckhack\$@

[bx_${BX_NAME}_visitor]
type = stcp
role = visitor
server_name = bx_${BX_NAME}
sk = ${BX_NAME}
bind_addr = 0.0.0.0
bind_port = ${BX_PORT}
EOF

# start visitor
ps aux | grep frpc_${BX_NAME}.ini | grep -v grep | awk '{print $2}' | xargs kill -9 1>/dev/null 2>&1
nohup ./frpc -c frpc_${BX_NAME}.ini &
echo start blackbox visitor  ${BX_NAME} with config frpc_${BX_NAME}.ini and port $BX_PORT
popd 1>/dev/null 2>&1