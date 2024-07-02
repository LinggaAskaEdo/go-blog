#!/bin/bash

# Define a timestamp function
# timestamp() {
#   date +"%Y%m%d%N" # current time
# }

PATH="./etc/sql/"
# TIME=$(date +"%Y%m%d%N")
# FILE=timestamp${name}

echo $PATH`date +"%Y%m%d%N`

# cat << 'EOF' > $PATH$FILE
# -- +migrate Up


# -- +migrate Down

# EOF

# # do something...
# timestamp # print timestamp
# # do something else...
# timestamp # print another timestamp
# # continue...