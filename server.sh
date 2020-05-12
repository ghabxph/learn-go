#!/bin/bash

echo "Sets env variable"
export SFS_WEB_DIR=`pwd`/web

echo "Running bin/webserver"
bin/webserver
echo ""
