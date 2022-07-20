#! /bin/sh

echo "CLINE: $CLINE"
exec test "$CLINE" = "a1"
