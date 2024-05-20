#!/khulnasoft/bin/bash

############################################################
# Start Khulnasoft Agent Services
############################################################
echo "Start Khulnasoft services... Console is $MGMT_CONSOLE_URL"
# exec setsid /khulnasoft/usr/local/bin/start-df-services.sh &
/khulnasoft/bin/bash /khulnasoft/usr/local/bin/start-df-services.sh &
# Wait for the agent to start
/khulnasoft/bin/sleep 20
echo "Khulnasoft agent started..."

############################################################
# Start the customer application entry point below... 
############################################################
#echo "Starting the customer application entry point below..."
#cust-entry-point.sh "$@"

# Execute customer entry-point if provided as arguments
if [ "$#" -ne 0 ]; then
    echo "Application entry-point specified as arguments to khulnasoft entrypoint. Execute application entrypoint."
    echo executing -- "$@"
    "$@"
else
    echo "No application entry-point specified as arguments to khulnasoft entrypoint."
fi

#echo "Block to avoid agent container from exiting fargate. Not needed if customer application blocks"
#/khulnasoft/usr/local/bin/block-scratch.sh
echo "Agent entry-point-scratch exiting...."

