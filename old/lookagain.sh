find . -name "*.sh" -type f -printf "%f\n" | rev | cut -c4- | rev
