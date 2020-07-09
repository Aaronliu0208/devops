#!/usr/bin/env python
# -*- coding: utf-8 -*-
import argparse
import logging
from ansible.parsing.dataloader import DataLoader
from ansible.inventory.manager import InventoryManager
import subprocess
from collections import namedtuple
import csv

def run_cmd(cmd):
    sp = subprocess.Popen(['/bin/bash', '-c', cmd], stdout=subprocess.PIPE)
    return sp.stdout.readlines()
def rand_passwd(host):
    password = run_cmd("""ssh  -o StrictHostKeyChecking=no -t root@{} openssl rand -base64 16""".format(host))
    #print(type(password))
    return password

def update_passwd(host,password):
    cmd = """ssh -o StrictHostKeyChecking=no -t root@{0} "echo '{1}' |passwd --stdin root" """.format(host,password)
    run_cmd(cmd)


def main():
    PasswdInfo = namedtuple('PasswdInfo', ['ip', 'key'])
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

    passwd = []
    for host in hosts:
        logging.debug('==============begin detect host %s' %(host))
        p = rand_passwd(host)[0].decode("utf-8").strip()
        #print(p)
        update_passwd(host,p)
        passwd.append(PasswdInfo(host, p))

    with open('passwdinfo.csv','w') as f:
        w = csv.writer(f,quotechar='"')
        for s in passwd:
            w.writerow(s)


if __name__ == "__main__":
    main()






