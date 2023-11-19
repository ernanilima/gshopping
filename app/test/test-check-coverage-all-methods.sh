#!/bin/sh

COVERAGE_OUT="coverage.out"

if [ ! -e $COVERAGE_OUT ]; then
    exit 1
fi

echo -e "⏳\t Verificando a cobertura dos metodos"

go tool cover -func=$COVERAGE_OUT > coverage_output.txt

echo ".-------------------------------------------------------------------------------------------------."
echo "| Modulo                                                  | Metodo/Funcao             | Cobertura |"
echo "| ------------------------------------------------------- + ------------------------- + --------- |"
awk 'BEGIN {FS=OFS=" "} $2 !~ /statements/ {split($1,a,"app/"); printf("| %-55s | %-25s | %-6s %s |\n", substr(a[2], 1, 55), substr($2, 1, 25), $3, ($3+0 > 80 ? "✅" : "⚠️"))}' coverage_output.txt
echo "'-------------------------------------------------------------------------------------------------'"
