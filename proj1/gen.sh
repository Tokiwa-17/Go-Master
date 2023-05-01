#! /bin/sh

set -x

num=$#
if [ $# -eq 0 ]
then
  utils/mac-intel/bin/gensort exmaple1.dat 1mb
else
  utils/mac-intel/bin/gensort exmaple1.dat $1
fi