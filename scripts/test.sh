#!/usr/bin/env bash
set -eo pipefail

go test -v --cover ./... -count=1 | \
    sed ''/PASS/s//$(printf "\033[32mPASS\033[0m")/'' | \
    sed ''/FAIL/s//$(printf "\033[31mFAIL\033[0m")/'' | \
    sed ''/RUN/s//$(printf "\e[96mRUN\e[0m")/'' | \
    sed ''/coverage/s//$(printf "\e[93mcoverage\e[0m")/''
