#!/bin/bash

Help()
{
   echo "Add description of the script functions here."
   echo
   echo "Syntax: generator.sh -f file_sql_name"
   echo "options:"
   echo "-f     SQL file name separate with an underscore if more than 1 word"
   echo
}

Create()
{
    TIME=`date +"%Y%m%d%H%M"`

    FILE_PATH="./etc/sql/"
    FILE_PATH+="$TIME"
    FILE_PATH+="_$NAME"
    FILE_PATH+=".sql"

    cat << 'EOF' > $FILE_PATH
-- +migrate Up


-- +migrate Down


EOF
}

while getopts f: flag
do
    case "${flag}" in
        f) NAME=${OPTARG};;
        \?) exit;;
    esac
done

if [ -z "$NAME" ] 
then
    Help  
else
    Create
    echo "File created"
fi 