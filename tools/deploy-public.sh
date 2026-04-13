#!/bin/bash

TARGET_HOST="${1:-parking}"

# rebuild the public/ folder
go run ./cmd/gobbs-static -base-url / -out public || {
  echo "error: couldn't build 'public/'"
  exit 2
}

# push the public/ folder to the production server.
# note: it's "public/" not "public" - if we drop the slash, rsync will try
#       to create the folder on the remote server!
rsync -avz --delete public/ "${TARGET_HOST}:/var/www/forums.playbymail.dev/" || {
  echo "error: couldn't push 'public/'"
  exit 2
}

echo " info: rebuilt 'public/' and pushed to production server"
exit 0
