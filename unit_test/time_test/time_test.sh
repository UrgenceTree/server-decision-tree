#!/bin/bash

echo "############ GO TIME TEST ################"
time ./go_part/go_test.sh
echo ""

echo "############ PYTHON TIME TEST ################"
time ./python_part/py_test.sh
