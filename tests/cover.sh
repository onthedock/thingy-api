#!/bin/bash
cover () {
    local t=$(mktemp -t cover_XXX)
    go test $COVERFLAGS -coverprofile=$t $@ \
        && go tool cover -func=$t \
        && go tool cover -html=$t -o /workspaces/thingy-api/cover.html \
        && unlink $t
}