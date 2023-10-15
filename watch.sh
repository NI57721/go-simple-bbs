#!/usr/bin/bash -u

pre_hash=

while true; do
  sleep 1

  hash=$(sha1sum main.go)
  pid=$(ps aux | grep -v grep | grep go-simple-bbs | awk '{print $2}')

  if [ "$pre_hash" = "$hash" ]; then
    continue
  fi

  if [ ! -z "$pid" ]; then
    kill -9 $pid
  fi

  sha1sum main.go > hash
  pre_hash=$(cat hash 2> /dev/null)

  go fmt
  go run . &
  echo "Hot Reload has done. $(date)"
done

