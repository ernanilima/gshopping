#!/bin/sh

COVERAGE_OUT="coverage.out"

if [ ! -e $COVERAGE_OUT ]; then
    exit 1
fi

echo -e "⏳\t Validando cobertura"

COVERAGE_REQUESTED=90
COVERAGE_PROVIDED=$(go tool cover -func=$COVERAGE_OUT | grep total | awk '{print $3}' | tr -d '%')
ACHIEVED=$(echo "$COVERAGE_PROVIDED >= $COVERAGE_REQUESTED" | bc -l)

echo ".--------------------------------------------------------."
echo "| Cobertura solicitada | Cobertura fornecida | Alcancado |"
echo "| -------------------- + ------------------- + --------- |"
grep total coverage_output.txt | awk -v REQUESTED="$COVERAGE_REQUESTED%" -v ACHIEVED="$ACHIEVED" '{printf("| %-20s | %-19s | %-10s |\n", REQUESTED, $3, (ACHIEVED == 1 ? "✅" : "❌"))}'
echo "'--------------------------------------------------------'"

if [ $ACHIEVED -eq 0 ]; then
    exit 1
fi