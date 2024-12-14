#!/bin/env bash

PROJECT_FOLDER=/home/$USER/Projects

echo "-----------------  running tests -----------------"
cd ./tests/home/

printf "> momo update \n\n"
$PROJECT_FOLDER/momo/bin/linux/momo update

printf "\n--- Content of config.toml ---\n\n"
cat config.toml

printf "\n\n> momo install \n\n"
$PROJECT_FOLDER/momo/bin/linux/momo install

echo "-----------------   tests done   -----------------"
