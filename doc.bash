#!/bin/bash
set -e

SRC=./article/.config/gsk/docusaurus/
BASE=~/Dev/work/gostartkit/www/gostartkit.com/docusaurus
DIST=$BASE/docs/golang/example/article-ce/


gsk doc

mkdir -p $DIST

cp $SRC/* $DIST

cd $BASE && pnpm start