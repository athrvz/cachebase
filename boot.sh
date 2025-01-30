set -e # Exit early if any commnads fail. 

(
    cd "$(dirname "$0")" # cd to repo dir
    go build -o /tmp/cachebase app/*.go
)

exec /tmp/cachebase "$@"