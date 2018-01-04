#!/usr/bin/env bash

NoMonit=$1
HOME=/home/app

[ "$NoMonit" != "no-monit" ] && {
	monit -c /home/app/monit/monitrc unmonitor aussie-parking
}

if [[ -f $HOME/harp/aussie-parking/app.pid ]]; then
	target=$(cat $HOME/harp/aussie-parking/app.pid);
	if ps -p $target > /dev/null; then
		kill -KILL $target; > /dev/null 2>&1;
	fi
fi
cd /home/app/src/github.com/theplant/aussie
GOPATH=/home/app nohup /home/app/bin/aussie-parking >> $HOME/harp/aussie-parking/app.log 2>&1 &
echo $! > $HOME/harp/aussie-parking/app.pid

[ "$NoMonit" != "no-monit" ] && {
	monit -c /home/app/monit/monitrc monitor aussie-parking
}
