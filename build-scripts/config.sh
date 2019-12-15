#! /bin/bash

servers=(
    192.168.1.100
    192.168.1.101
    192.168.1.102
    192.168.1.103
)

ssh_pass="password"

nodes=7

APPDIR=~/apriori
APP=~/apriori/app
MAPDIR=~/apriori/maps
RESULTDIR=~/apriori/resuls
PIFILE=~/pifile
PIFILE4=~/pi4