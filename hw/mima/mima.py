#!/usr/bin/env python
# -*- coding: utf-8 -*-
import argparse
import logging
from ansible.parsing.dataloader import DataLoader
from ansible.inventory.manager import InventoryManager
import subprocess
from collections import namedtuple
import csv

def runCmd(cmd):
    sp = subprocess.Popen(['/bin/bash', '-c', cmd], stdout=subprocess.PIPE)
    return sp.stdout.readlines()

def parseUsers(host):
    lines = runCmd("""ssh -t root@{} cat /etc/passwd | grep /bin/bash | awk -F ':' '{{print $1}}'""".format(host))
    users = []
    for line in lines:
        users.append(line.decode("utf-8").strip())
    return users
def parsePubkey(host, user):
    cmd = """ssh -t root@{0} <<EOF
    if [ -f /home/{1}/.ssh/authorized_keys ];then
        cat /home/{1}/.ssh/authorized_keys
    fi
EOF""".format(host, user)
    lines = runCmd(cmd)
    return lines

def main():
    SecretInfo = namedtuple('SecretInfo', ['ip', 'user', 'key'])
    parser = argparse.ArgumentParser()
    parser.add_argument('-i', '--inventory', help='inventory to fetch secret', required=False, default='hosts', type=str)
    parser.add_argument('-d', '--debug', help='show debug message', action='store_true')
    parser.add_argument("module")
    args = parser.parse_args()
    host_file=args.inventory
    if args.debug == True:
        logging.basicConfig(level=logging.DEBUG)
    logging.debug('Reading configuration from %s' %(host_file))

    data_loader = DataLoader()
    inventory = InventoryManager(loader = data_loader,
                             sources=[host_file])
    hosts = inventory.get_hosts(args.module)

    secret = []
    for host in hosts:
        logging.debug('==========begin detect host %s' %(host))
        users = parseUsers(host)
        for u in users:
            keys = parsePubkey(host, u)
            for k in keys:
                secret.append(SecretInfo(host, u, k.decode("utf-8").strip()))

    with open('secretinfo.csv','w') as f:
        w = csv.writer(f, quotechar='"')
        for s in secret:
            w.writerow(s)
        
        


if __name__ == "__main__":
    main()