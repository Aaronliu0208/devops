#!/bin/bash

# try to generate builder num
echo $(date +"%Y%m%d%H%M").$(git rev-parse --short HEAD).$(git symbolic-ref --short -q HEAD)