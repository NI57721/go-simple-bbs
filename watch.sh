#!/usr/bin/bash -u

pre_hash=
pid=

while true; do
  sleep 1

  hash=$(sha1sum main.go)

  if [ "$pre_hash" = "$hash" ]; then
    continue
  fi

  if [ ! -z "$pid" ]; then
    pkill -P $pid
  fi

  sha1sum main.go > hash
  pre_hash=$(cat hash 2> /dev/null)

  go fmt
  go run . &
  pid=$!
  echo "Hot Reload has done. $(date)"
done

