# run server and web build. need tmux already open

set -exu
HERE=$(dirname $(realpath $BASH_SOURCE))
cd $HERE/..

tmux rename-window spawn

tmux new-window -n run -c $HERE/../kj-study-web
tmux send "pnpm watch" Enter

tmux split-window -h -c $HERE
tmux send "bash kj-study.sh run" Enter