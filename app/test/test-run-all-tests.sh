#!/bin/sh

sleep 3

docker rm -vf $(docker ps -a -q) > /dev/null 2>&1
docker rmi -f $(docker images -a -q) > /dev/null 2>&1

echo -e "⏳\t Executando todos os testes"

COVERAGE_OUT="coverage.out"
CIIGNORE_PATH="./app/test/.ciignore"
PKGS="$(go list ./... | grep -xvf "$CIIGNORE_PATH")"
COVERPKG="$(echo "$PKGS" | tr '\n' ',')"

go test -covermode=count -coverpkg="$COVERPKG" -coverprofile="$COVERAGE_OUT" $PKGS > test_output.txt

if ! grep -q "FAIL" test_output.txt; then
    echo ".---------------------------------------------------------."
    echo "| Status | Modulo                             | Cobertura |"
    echo "| ------ + ---------------------------------- + --------- |"
    awk 'BEGIN {FS=OFS=" "} $1 != "?" {split($2,a,"gshopping"); printf("| %-6s | %-34s | %-6s %s |\n", $1, a[2], $5, ($5+0 > 80 ? "✅" : "⚠️"))}' test_output.txt
    echo "'---------------------------------------------------------'"
else
    cat test_output.txt
    rm $COVERAGE_OUT
    exit 1
fi
