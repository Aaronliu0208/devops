#!/usr/bin/env python
# -*- coding: utf-8 -*-
from nginx.config.builder import NginxConfigBuilder
from nginx.config.api import Config, Section, Location
nginx = NginxConfigBuilder(daemon='on')
with nginx.add_server() as server:
    server.add_route('/foo', proxy_pass='upstream').end()
print(nginx)

events = Section('events', worker_connections='1024')
http = Section('http', include='../conf/mime.types')
http.sections.add(
    Section(
        'server',
        Location(
            '/foo',
             proxy_pass='upstream',
        ),
        server_name='_',
    )
)
nginx2 = Config(
    events,
    http,
    worker_processes='auto',
    daemon='on',
    error_log='var/error.log',
)
print(nginx2)
