#!/usr/bin/env bash

test -d gitdir && rm -rf gitdir

gitdir=${1:=gitdir}
basedir=${2:-$(pwd)}

wd=$(pwd)

cd $basedir
mkdir -p $gitdir
touch $gitdir/.none
cd $gitdir
git init
git add .
git commit -m 'Test commit'
git tag 1.0.0
git tag 1.1.0

cd $wd
