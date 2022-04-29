#!/usr/bin/env python
# -*- coding: UTF-8 -*- 

# 使用 inventroy以及variables
from __future__ import (absolute_import, division, print_function)
__metaclass__ = type
import os, random, string
from ansible.inventory.manager import InventoryManager
from ansible.parsing.dataloader import DataLoader
from ansible.vars.manager import VariableManager

SCRIPT_DIR = os.path.dirname(os.path.abspath(__file__))
ROOT_DIR = os.path.dirname(SCRIPT_DIR)

print("script dir is "+SCRIPT_DIR)
print("root dir is "+ROOT_DIR)

length = 13
lower = string.ascii_lowercase
upper = string.ascii_uppercase
num = string.digits
symbols = string.punctuation
chars = lower + upper + num
random.seed = (os.urandom(1024))

print (''.join(random.sample(chars,length)))

dl = DataLoader()
im = InventoryManager(loader=dl, sources=[os.path.join(ROOT_DIR,'hushi/icii_hosts')])
print(im.get_hosts())
print(im.list_groups())

vm = VariableManager(loader=dl, inventory=im)
print(vm.get_vars())

my_host = im.get_host('10.153.51.82')
print(type(my_host))
print(vm.get_vars(host=my_host))