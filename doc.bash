#!/bin/bash
set -e

SRC=./article/.config/gsk/docusaurus/
BASE=~/Dev/work/gostartkit/www/gostartkit.com/
DIST=$BASE/docusaurus/docs/golang/example/article-ce/


gsk doc

mkdir -p $DIST

cp $SRC/* $DIST

# cd $BASE && ./dev.bash