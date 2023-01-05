#!/usr/bin/env bash

HOST_CLUSTER_NAME=${HOST_CLUSTER_NAME:-"karmada-host"}

filepath="/Users/zhang/drzhang/demo/go/go_base/text.txt"

function foo() {
    if [[ -e ${filepath} ]];then
      echo 'hello'
      return 0
    fi
    echo 'world'
   return 1
}

pwd
echo "sum: $(foo)"
echo $HOST_CLUSTER_NAME