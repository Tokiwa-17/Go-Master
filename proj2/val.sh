#! /bin/sh

set -x

input_path=src/dist/testcase1
for i in {0..3}
do 
    if [ "$i" -eq 0 ]
    then 
        cp $input_path/input-$i.dat INPUT
    else
        cat $input_path/input-$i.dat >> INPUT
    fi
done

for i in {0..3}
do
    if [ "$i" -eq 0 ]
    then 
        cp $input_path/output-$i.dat OUTPUT
    else 
        cat $input_path/output-$i.dat >> OUTPUT 
    fi 
done

utils/linux-amd64/showsort INPUT | sort > REF_OUTPUT
diff REF_OUTPUT OUTPUT