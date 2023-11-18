#!/bin/sh

set -e

echo "start the app"

# DockerfileのCMDに指定されたコマンドを実行する
exec "$@"