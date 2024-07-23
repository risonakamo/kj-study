set -exu
HERE=$(dirname $(realpath $BASH_SOURCE))
cd $HERE/..

go build -o kj-study.exe bin/kj-study/kj_study.go

set +u
if [[ "$1" == "run" ]]; then
    ./kj-study.exe
fi