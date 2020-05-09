#!/usr/bin/python
# -*- coding: utf-8 -*-


import argparse
def add(args):
    """ add monitor site """
    print('call add')
    print(args)

def main():
    parser = argparse.ArgumentParser()

    parser.add_argument('-c', '--config', help='config file path for cli', required=False)
    subparsers = parser.add_subparsers(help='manage monitor target for site')
    parser_add = subparsers.add_parser('add', help='add monitor site')
    parser_remove = subparsers.add_parser('remove', help='remove monitor site by name')
    parser_list = subparsers.add_parser('list', help='list monitor site')

    parser_add.add_argument('-n', '--name', help='name of site', dest='name')
    parser_add.set_defaults(func=add)

    args = parser.parse_args()
    try:
        args.func(args)
    except AttributeError:
        parser.print_help()
        parser.exit()

if __name__ == "__main__":
    main()