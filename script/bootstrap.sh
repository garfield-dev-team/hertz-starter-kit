
#!/bin/bash
CURDIR=$(cd $(dirname $0); pwd)
BinaryName="hertz_app"
echo "$CURDIR/bin/${BinaryName}"
exec $CURDIR/bin/${BinaryName}
