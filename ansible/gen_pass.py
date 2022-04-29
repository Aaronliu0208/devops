#!/usr/bin/env python
# -*- coding: UTF-8 -*- 
from __future__ import (absolute_import, division, print_function)
__metaclass__ = type

import os, random, string, argparse
from ansible.inventory.manager import InventoryManager
from ansible.parsing.dataloader import DataLoader
from ansible.vars.manager import VariableManager

SCRIPT_DIR = os.path.dirname(os.path.abspath(__file__))
ROOT_DIR = os.path.dirname(SCRIPT_DIR)
length = 13
lower = string.ascii_lowercase
upper = string.ascii_uppercase
num = string.digits
symbols = string.punctuation
chars = lower + upper + num
random.seed = (os.urandom(1024))

def main():
    parser = argparse.ArgumentParser(description='inventory password generator')
    parser.add_argument('input', type=str, help='inventory file')
    parser.add_argument('-o', '--output', type=str, required=False,
        default="output.csv",
        metavar="/path/to/output/",
        help='output file for new inventory')
    args = parser.parse_args()
    inventory_file = args.input
    dl = DataLoader()
    im = InventoryManager(loader=dl, sources=[inventory_file])
    vm = VariableManager(loader=dl, inventory=im)
    print('[all]')
    for host in im.get_hosts():
        #print("{} password={}".format(host.address, ''.join(random.sample(chars,length))))
        print(vm.get_vars(host=host))

if __name__ == "__main__":
    main()