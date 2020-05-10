#!/usr/bin/python
# -*- coding: utf-8 -*-
import argparse
import configparser
import logging
import os
import consul
import sys

config = configparser.ConfigParser()
def add(args,client):
    """ add monitor site """
    global config
    logging.debug('begin add service')
    if not args.name:
        logging.fatal("invalid input name empty")
        sys.exit("error of input")
    if not args.url:
        logging.fatal("invalid input namurle empty")
        sys.exit("error of input")

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
    return consul.Consul(host=server, port=port)

def main():
    global config
    parser = argparse.ArgumentParser()

    parser.add_argument('-c', '--config', help='config file path for cli', required=False, default='config.ini', type=str)
    parser.add_argument('-s', '--server', help='server host for consul', required=False, type=str)
    parser.add_argument('-p', '--port', help='server port for consul', required=False, type=int)
    parser.add_argument('-d', '--debug', help='show debug message for cli', action='store_true')
    subparsers = parser.add_subparsers(help='manage monitor target for site')
    parser_add = subparsers.add_parser('add', help='add monitor site')
    parser_add.add_argument('-n', '--name', help='name of site', dest='name')
    parser_add.add_argument('-u', '--url', help='url for monitor', dest='url')
    parser_add.add_argument('')
    parser_add.set_defaults(func=add)

    parser_remove = subparsers.add_parser('remove', help='remove monitor site by name')
    parser_list = subparsers.add_parser('list', help='list monitor site')

    args = parser.parse_args("add".split())
    
    if args.debug == True:
        logging.basicConfig(level=logging.DEBUG)
    config_file = args.config
    logging.debug('Reading configuration from %s' %(config_file))
    config.read(config_file)

    client = createConsul(args)
    client.agent.Service.register()
    # run sub command
    try:
        args.func(args, client)
    except AttributeError:
        parser.print_help()
        parser.exit()

if __name__ == "__main__":
    main()