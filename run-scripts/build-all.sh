set -exu
HERE=$(dirname $(realpath $BASH_SOURCE))
cd $HERE

bash data-splitter.sh
bash word-downloader.sh
bash kj-study.sh

cd ../kj-study-web
rm -rf build
pnpm build