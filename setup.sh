#!/bin/bash

CURR_DIR=`pwd`

mkdir -p ./data && cd data

if [[ ! -d ./ml-20m ]]; then
    wget http://files.grouplens.org/datasets/movielens/ml-20m.zip && 
        unzip ml-20m.zip
fi

if [[ ! -f ./ml-20m/ml-youtube.csv ]]; then
    cd ml-20m
    wget http://files.grouplens.org/datasets/movielens/ml-20m-youtube.zip &&
        unzip ml-20m-youtube.zip
fi

cd $CURR_DIR
