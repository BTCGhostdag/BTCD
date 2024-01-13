#!/bin/bash
rm -rf /tmp/BTCD-temp

BTCD --devnet --appdir=/tmp/BTCD-temp --profile=6061 --loglevel=debug &
KASPAD_PID=$!
KASPAD_KILLED=0
function killBTCDdIfNotKilled() {
    if [ $KASPAD_KILLED -eq 0 ]; then
      kill $KASPAD_PID
    fi
}
trap "killBTCDdIfNotKilled" EXIT

sleep 1

application-level-garbage --devnet -alocalhost:42611 -b blocks.dat --profile=7000
TEST_EXIT_CODE=$?

kill $KASPAD_PID

wait $KASPAD_PID
KASPAD_KILLED=1
KASPAD_EXIT_CODE=$?

echo "Exit code: $TEST_EXIT_CODE"
echo "BTCDd exit code: $KASPAD_EXIT_CODE"

if [ $TEST_EXIT_CODE -eq 0 ] && [ $KASPAD_EXIT_CODE -eq 0 ]; then
  echo "application-level-garbage test: PASSED"
  exit 0
fi
echo "application-level-garbage test: FAILED"
exit 1