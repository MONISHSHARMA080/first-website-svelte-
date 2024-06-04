#!/bin/sh

# Start the Go server
./go_server &

# Start the Svelte app (dev mode)
npm run dev &

# Wait for any process to exit
wait -n

# Exit with status of process that exited first
exit $?