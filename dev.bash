#!/bin/bash
set -e

# pnpm create docusaurus@latest docusaurus classic

# rm -rf docusaurus/docs/*

# mkdir -p docusaurus/docs/article/

cp ./article/.config/gsk/docusaurus/* ./docusaurus/docs/article/

cd docusaurus && pnpm start