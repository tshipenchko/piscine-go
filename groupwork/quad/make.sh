#/usr/bin/env bash
# from https://stackoverflow.com/a/246128
SCRIPT_DIR=$( cd -- "$( dirname -- "${BASH_SOURCE[0]}" )" &> /dev/null && pwd )
set -e
cd $SCRIPT_DIR

rm -rf build
mkdir -p build
cd build

go build -o cmd ../cmd

ln -s cmd quadA
ln -s cmd quadB
ln -s cmd quadC
ln -s cmd quadD
ln -s cmd quadE
