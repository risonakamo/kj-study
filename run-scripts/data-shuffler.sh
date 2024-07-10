set -exu
HERE=$(dirname $(realpath $BASH_SOURCE))
cd $HERE/..

go build -o data-shuffler.exe bin/data-shuffler/data-shuffler.go

set +u
if [[ "$1" == "run" ]]; then
    ./data-shuffler.exe
fi