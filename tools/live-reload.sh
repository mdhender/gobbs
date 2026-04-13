#!/bin/zsh

export BROWSER_SYNC_HOST=${BROWSER_SYNC_HOST:-gobbs.test}
export BROWSER_SYNC_PROXY=http://gobbs.test:8080

npx --yes browser-sync start --config browser-sync.config.js
