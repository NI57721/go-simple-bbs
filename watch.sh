#!/usr/bin/bash -u

while true; do
  sleep 1
  modified_time=$(stat -c %Y .)
  if [ "${last_modified_time:-}" = "$modified_time" ]; then
    continue
  fi
  if [ -v pid ]; then
    pkill -P "$pid"
  fi
  last_modified_time=$modified_time
  go fmt
  go run . &
  pid=$!
  echo "Rebuilt at $(date "+%H:%M:%S")"
done

