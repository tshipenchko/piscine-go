#/usr/bin/env bash
# from https://stackoverflow.com/a/246128
SCRIPT_DIR=$( cd -- "$( dirname -- "${BASH_SOURCE[0]}" )" &> /dev/null && pwd )

cd $SCRIPT_DIR
rm -rf test
mkdir -p test
cd test
echo "*" > .gitignore
cat > main.go << EOT
package main

import (
	"fmt"
)

func main() {
	fmt.Println()
}
EOT
