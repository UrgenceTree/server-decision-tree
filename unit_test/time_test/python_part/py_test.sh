#!/bin/bash

# Création d'un fichier de test
echo "yes" > data_answer.txt
echo "yes" >> data_answer.txt
echo "yes" >> data_answer.txt

# Exécution du script à tester
output=$(./python_part/tree.py -f data_answer.txt)

# Vérification du résultat
expected_output="Score final du patient : 110"
if [[ "$output" == "$expected_output" ]]; then
    echo "Py Test 1 réussi -> $output"
else
    echo "Py Test 1 échoué -> Le résultat obtenu ($output) ne correspond pas au résultat attendu ($expected_output)"
fi

# Nettoyage des fichiers temporaires
rm -f data_answer.txt


################################

echo "yes" > data_answer.txt
echo "yes" >> data_answer.txt
echo "no" >> data_answer.txt

output=$(./python_part/tree.py -f data_answer.txt)

expected_output="Score final du patient : 110"
if [[ "$output" == "$expected_output" ]]; then
    echo "Py Test 2 réussi -> $output"
else
    echo "Py Test 2 échoué -> Le résultat obtenu ($output) ne correspond pas au résultat attendu ($expected_output)"
fi

rm -f data_answer.txt


################################

echo "yes" > data_answer.txt
echo "no" >> data_answer.txt
echo "no" >> data_answer.txt

output=$(./python_part/tree.py -f data_answer.txt)

expected_output="Score final du patient : 10"
if [[ "$output" == "$expected_output" ]]; then
    echo "Py Test 3 réussi -> $output"
else
    echo "Py Test 3 échoué -> Le résultat obtenu ($output) ne correspond pas au résultat attendu ($expected_output)"
fi

rm -f data_answer.txt


################################

echo "yes" > data_answer.txt
echo "no" >> data_answer.txt
echo "yes" >> data_answer.txt

output=$(./python_part/tree.py -f data_answer.txt)

expected_output="Score final du patient : 20"
if [[ "$output" == "$expected_output" ]]; then
    echo "Py Test 4 réussi -> $output"
else
    echo "Py Test 4 échoué -> Le résultat obtenu ($output) ne correspond pas au résultat attendu ($expected_output)"
fi

rm -f data_answer.txt

################################

echo "no" > data_answer.txt
echo "no" >> data_answer.txt
echo "no" >> data_answer.txt

# Exécution du script à tester
output=$(./python_part/tree.py -f data_answer.txt)

# Vérification du résultat
expected_output="Score final du patient : 0"
if [[ "$output" == "$expected_output" ]]; then
    echo "Py Test 5 réussi -> $output"
else
    echo "Py Test 5 échoué -> Le résultat obtenu ($output) ne correspond pas au résultat attendu ($expected_output)"
fi

# Nettoyage des fichiers temporaires
rm -f data_answer.txt


################################

echo "no" > data_answer.txt
echo "no" >> data_answer.txt
echo "yes" >> data_answer.txt

output=$(./python_part/tree.py -f data_answer.txt)

expected_output="Score final du patient : 0"
if [[ "$output" == "$expected_output" ]]; then
    echo "Py Test 6 réussi -> $output"
else
    echo "Py Test 6 échoué -> Le résultat obtenu ($output) ne correspond pas au résultat attendu ($expected_output)"
fi

rm -f data_answer.txt


################################

echo "no" > data_answer.txt
echo "yes" >> data_answer.txt
echo "yes" >> data_answer.txt

output=$(./python_part/tree.py -f data_answer.txt)

expected_output="Score final du patient : 0"
if [[ "$output" == "$expected_output" ]]; then
    echo "Py Test 7 réussi -> $output"
else
    echo "Py Test 7 échoué -> Le résultat obtenu ($output) ne correspond pas au résultat attendu ($expected_output)"
fi

rm -f data_answer.txt


################################

echo "no" > data_answer.txt
echo "yes" >> data_answer.txt
echo "no" >> data_answer.txt

output=$(./python_part/tree.py -f data_answer.txt)

expected_output="Score final du patient : 0"
if [[ "$output" == "$expected_output" ]]; then
    echo "Py Test 8 réussi -> $output"
else
    echo "Py Test 8 échoué -> Le résultat obtenu ($output) ne correspond pas au résultat attendu ($expected_output)"
fi

rm -f data_answer.txt
