#!/usr/bin/env bash
set -eo pipefail

refreshdb() {
    # Stop postgres if already running
    docker-compose down

    # Clear old volumes
    docker volume rm devicecommander_postgres || true
    docker volume rm devicecommander_pgadmin || true

    # Start Up Postgres DB
    docker-compose up -d

    # Run Migrations
    go run main.go migrate-db
}

if [[ "$1" == "true" ]]; then
    refreshdb
fi

go test -v --cover ./... -count=1 |
    sed ''/PASS/s//$(printf "\033[32mPASS\033[0m")/'' |
    sed ''/FAIL/s//$(printf "\033[31mFAIL\033[0m")/'' |
    sed ''/RUN/s//$(printf "\e[96mRUN\e[0m")/'' |
    sed ''/coverage/s//$(printf "\e[93mcoverage\e[0m")/''


