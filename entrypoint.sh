#!/usr/bin/env bash

echo ">>> Creating a ~/.pcppass file for $PCP_USER"
echo *:9898:$PCP_USER:$PCP_PASSWORD > ~/.pcppass
chmod 0600 ~/.pcppass

exec "$@"
