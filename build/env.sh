#!/bin/sh

set -e

if [ ! -f "build/env.sh" ]; then
    echo "$0 must be run from the root of the repository."
    exit 2
fi

# Create fake Go workspace if it doesn't exist yet.
workspace="$PWD/build/_workspace"
root="$PWD"
xpsdir="$workspace/src/github.com/xpaymentsorg"
if [ ! -L "$xpsdir/go-xpayments" ]; then
    mkdir -p "$xpsdir"
    cd "$xpsdir"
    ln -s ../../../../../. go-xpayments
    cd "$root"
fi

# Set up the environment to use the workspace.
GOPATH="$workspace"
export GOPATH

# Run the command inside the workspace.
cd "$xpsdir/go-xpayments"
PWD="$xpsdir/go-xpayments"

# Launch the arguments with the configured environment.
exec "$@"
