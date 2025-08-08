#!/bin/sh

sleep 3

echo -e "⏳\t Validando ciignore"

CIIGNORE_PATH="./app/test/.ciignore"

go list ./... > list_output.txt

cat list_output.txt | grep -xvf $CIIGNORE_PATH > list_without_ciignore.txt
cat $CIIGNORE_PATH | grep -xvf list_output.txt > invalid_ciignore.txt

print_title() {
  printf "| %-3s %-53s |\n" "$1" "$2"
}

if [ ! -s "invalid_ciignore.txt" ]; then
    echo ".---------------------------------------------------------."
    print_title "🗑️" "Pacotes a serem ignorados com .ciignore"
    echo "| ------------------------------------------------------- |"
    cat $CIIGNORE_PATH | awk '{printf("| %-55s |\n", substr($1, 1, 55))}'
    echo "'---------------------------------------------------------'"

    echo ".---------------------------------------------------------."
    print_title "📣" "Pacotes a serem testados"
    echo "| ------------------------------------------------------- |"
    awk '{printf("| %-55s |\n", substr($1, 1, 55))}' list_without_ciignore.txt
    echo "'---------------------------------------------------------'"
else
    echo ".---------------------------------------------------------."
    print_title "❌" "Pacote(s) invalido(s)"
    echo "| ------------------------------------------------------- |"
    awk '{printf("| %-55s |\n", substr($1, 1, 55))}' invalid_ciignore.txt
    echo "'---------------------------------------------------------'"
    exit 1
fi
