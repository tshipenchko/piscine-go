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

	"github.com/01-edu/z01"

	"piscine"
)

func main() {
	piscine.DoNothing()
	fmt.Print("=== testing ===")
	z01.PrintRune('\n')
}
EOT
