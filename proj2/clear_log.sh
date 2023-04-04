#! /bin/sh

set -x

for i in $(seq 0 3)
do
	docker rm proj2-server$i-1
done
