#!/bin/zsh

export PREVIEW_ADDR=127.0.0.1:8080
export BROWSER_SYNC_PROXY=http://${PREVIEW_ADDR}

npx --yes browser-sync start --config browser-sync.config.js
