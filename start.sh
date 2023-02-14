#!/bin/sh
GOMON_IGNORE=${GOMON_IGNORE-''}

if [ "${GOMON_IGNORE}" = '' ]; then
  EXCLUDE_REGEX=''
else
  EXCLUDE_REGEX="-build.exclude_regex ${GOMON_IGNORE}"
fi

AIR_CONFIG=/app/.air.toml
if [ ! -f "$AIR_CONFIG" ];  then
    AIR_CONFIG=/.air.toml
fi
air -c "${AIR_CONFIG}" "${EXCLUDE_REGEX}"