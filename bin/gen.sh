#!/bin/sh

TASK_FILE="${1}.go"
TESTCASES_FILE="testcases/${TASK_FILE}"
RUN_FN=$2
TEST_CASE="T${RUN_FN}"
CONFIRM_REWRITE=$3


if [ "z${CONFIRM_REWRITE}" = "z" ]; then
    if [ -f "${TASK_FILE}" ]; then
        echo "File ${TASK_FILE} already exists."
        echo "Please confirm if you want to rewrite it:"
        echo "% ${0} ${1} ${2} yes"
        exit 1
    fi

    if [ -f "${TESTCASES_FILE}" ]; then
        echo "File ${TESTCASES_FILE} already exists."
        echo "Please confirm if you want to rewrite it:"
        echo "% ${0} ${1} ${2} yes"
        exit 1
    fi
fi

sed -e "s/NameOfTask/${TEST_CASE}/g" < tpl/_blank.go | sed -e "s/nameOfTask/${RUN_FN}/g" > "${TASK_FILE}"
echo "File ${TASK_FILE} created"
sed -e "s/NameOfTask/${TEST_CASE}/g" < tpl/_blank_testcases.go > "${TESTCASES_FILE}"
echo "File ${TESTCASES_FILE} created"
