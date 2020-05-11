#!/usr/bin/python
# -*- coding: utf-8 -*-
import argparse
import configparser
import logging
import os
import sys
import consul
from tabulate import tabulate


config = configparser.ConfigParser()
TABLE_HEADER = ['Name', 'Address']

def addSites(args,client):
    """ add monitor site """
    logging.debug('begin add service')
    if not args.name:
        logging.fatal("invalid input name empty")
        sys.exit("error of input")
    if not args.url:
        logging.fatal("invalid input namurle empty")
        sys.exit("error of input")

    service_name = args.name
    address = args.url
    meta = {}
    meta['address'] =address
    meta['service_name'] = service_name
    ok = client.agent.service.register(service_name, meta=meta)
    if ok:
        print("create service ok")
        # try to print service
        services = [(service_name, address)]
        print(tabulate(services, headers=TABLE_HEADER,  tablefmt='simple'))
    else:
        print("create service fail")

def listSites(args, client):
    """ list all serivce """
    res = client.agent.services()
    services = []
    for k,v in res.items():
       services.append((k, v['Meta']['address']))
    print(tabulate(services, headers=TABLE_HEADER,  tablefmt='simple'))

def removeSites(args, client):
    """ list all serivce """
    if not args.name:
        logging.fatal("invalid input name empty")
        sys.exit("error of input")
    ok = client.agent.service.deregister(args.name)
    if ok:
        print("delete service %s  ok"%args.name)
    else:
        print("delete service %s  fail"%args.name)
    

def createConsul(args):
    """ 
    创建consul类，如果命令行参数有host就从命令行参数取，如果没有就从config里。如果没有默认127.0.0.1:8500 
    """
    if args.server is not None:
        server = args.server
    elif config.has_option('main', 'server'):
        server = config.get('main', 'host')
    else:
        server = '127.0.0.1'

    if args.port is not None:
        port = args.port
    elif config.has_option('main', 'port'):
        port = int(config.get('main', 'port'))
    else:
        port = 8500
    logging.debug('current server is %s:%d' %(server, port))
    client = consul.Consul(host=server, port=port)
    return client

def main():
    global config
    parser = argparse.ArgumentParser()

    parser.add_argument('-c', '--config', help='config file path for cli', required=False, default='config.ini', type=str)
    parser.add_argument('-s', '--server', help='server host for consul', required=False, type=str)
    parser.add_argument('-p', '--port', help='server port for consul', required=False, type=int)
    parser.add_argument('-d', '--debug', help='show debug message for cli', action='store_true')
    subparsers = parser.add_subparsers(help='manage monitor target for site')
    
    parser_add = subparsers.add_parser('add', help='add monitor site')
    parser_add.add_argument('name', help='name of site')
    parser_add.add_argument('-u', '--url', help='url for monitor', dest='url')
    parser_add.set_defaults(func=addSites)

    parser_list = subparsers.add_parser('list', help='list monitor site')
    parser_list.set_defaults(func=listSites)

    parser_remove = subparsers.add_parser('remove', help='remove monitor site')
    parser_remove.add_argument('name', help='name of site')
    parser_remove.set_defaults(func=removeSites)

    args = parser.parse_args()
    
    if args.debug == True:
        logging.basicConfig(level=logging.DEBUG)
    config_file = args.config
    logging.debug('Reading configuration from %s' %(config_file))
    config.read(config_file)

    client = createConsul(args)
    # run sub command
    try:
        args.func(args, client)
    except AttributeError as ae:
        logging.error("AttributeError: %s"%ae)
        parser.print_help()
        parser.exit()
    except Exception as ex:
        print("Exception: %s"%ex)

if __name__ == "__main__":
    main()