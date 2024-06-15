set -exu
HERE=$(dirname $(realpath $BASH_SOURCE))
cd $HERE/..

go build -o word-downloader.exe bin/word-downloader/word_downloader.go

set +u
if [[ "$1" == "run" ]]; then
    ./word-downloader.exe
fi