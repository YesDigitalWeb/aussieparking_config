check host aussie-parking with address 127.0.0.1
  start program = "/bin/bash -c '/home/app/script/restart.sh no-monit'"
  stop program = "/bin/sh -c 'pkill -c -f  aussie-parking; exit0;'"
  if failed host 127.0.0.1 port 7000 then start
  if 5 restarts within 5 cycles then timeout