#!/bin/bash
rm -rf /tmp/BTCD-temp

BTCD --devnet --appdir=/tmp/BTCD-temp --profile=6061 &
KASPAD_PID=$!

sleep 1

infra-level-garbage --devnet -alocalhost:42611 -m messages.dat --profile=7000
TEST_EXIT_CODE=$?

kill $KASPAD_PID

wait $KASPAD_PID
KASPAD_EXIT_CODE=$?

echo "Exit code: $TEST_EXIT_CODE"
echo "BTCDd exit code: $KASPAD_EXIT_CODE"

if [ $TEST_EXIT_CODE -eq 0 ] && [ $KASPAD_EXIT_CODE -eq 0 ]; then
  echo "infra-level-garbage test: PASSED"
  exit 0
fi
echo "infra-level-garbage test: FAILED"
exit 1
