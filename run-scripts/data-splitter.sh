set -exu
HERE=$(dirname $(realpath $BASH_SOURCE))
cd $HERE/..

go build -o data-splitter.exe bin/data-splitter/data_splitter.go

set +u
if [[ "$1" == "run" ]]; then
    ./data-splitter.exe
fi