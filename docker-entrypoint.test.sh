#!/bin/sh

PID_PATH="/var/run/docker.pid"
TEST_VALIDATE_CIIGNORE_PATH="./app/test/test-validate-ciignore.sh"
TEST_RUN_ALL_TESTS_PATH="./app/test/test-run-all-tests.sh"
TEST_CHECK_COVERAGE_ALL_METHODS_PATH="./app/test/test-check-coverage-all-methods.sh"
TEST_VALIDATE_COVERAGE_PATH="./app/test/test-validate-coverage.sh"

start_docker() {
    local isLocal="$1"

    if [ -e $PID_PATH ]; then
        echo "♻️  Reiniciando o dockerd"
        kill $(cat $PID_PATH)
        rm $PID_PATH
    fi

    if [ $isLocal = true ]; then
        echo "☄️  Iniciando e executando os testes"
        dockerd &
    else
        echo "🐳  Iniciando sem executar os testes"
        dockerd
    fi
    
    for i in $(seq 1 10); do
        if docker info >/dev/null 2>&1; then
            break
        fi
        sleep 1
    done

    if ! docker info >/dev/null 2>&1; then
        echo "❌  Falha ao iniciar o dockerd"
        exit 1
    fi
}

test_validate_ciignore() {
    chmod +x $TEST_VALIDATE_CIIGNORE_PATH
    source $TEST_VALIDATE_CIIGNORE_PATH
}

test_run_all_tests() {
    chmod +x $TEST_RUN_ALL_TESTS_PATH
    $TEST_RUN_ALL_TESTS_PATH
}

test_check_coverage_all_methods() {
    chmod +x $TEST_CHECK_COVERAGE_ALL_METHODS_PATH
    $TEST_CHECK_COVERAGE_ALL_METHODS_PATH
}

test_validate_coverage() {
    chmod +x $TEST_VALIDATE_COVERAGE_PATH
    $TEST_VALIDATE_COVERAGE_PATH
}

if [ -z $TEST ] || [ $TEST == 'local' ]; then
    start_docker true

    test_validate_ciignore
    test_run_all_tests
    test_check_coverage_all_methods
    test_validate_coverage
elif ! [ -z $TEST ] && [ $TEST == 'ci' ]; then
    start_docker false
else
    echo "❌  O valor '$TEST' informado na variavel de ambiente 'TEST' eh invalido, informe TEST=local ou TEST=ci"
    exit 1
fi
