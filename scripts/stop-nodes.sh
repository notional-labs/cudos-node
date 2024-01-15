#!/bin/bash
# kill all sessions
kill_screen_sessions() {
    local session_id=$(screen -ls | grep 'validator[1-3]')
    if [ -z "$session_id" ]; then
        echo "No screen session found."
    else
        local validator_pids=($(screen -ls | grep 'validator[1-3]' | awk -F '[ .]+' '{print $1}'))
        for pid in "${validator_pids[@]}"; do
            screen -S $pid -X quit
            sleep 2
        done
    fi
}

# Function to kill processes bound to ports with 'cudo' in their name
kill_bound_ports() {
    while true; do
        local process_id=$(lsof -i -P | grep -m 1 cudo | awk '{print $2}')
        if [ -z "$process_id" ]; then
            break
        else
            kill -9 $process_id
            echo "Killed process $process_id"
            sleep 1
        fi
    done
}
kill_screen_sessions
kill_bound_ports