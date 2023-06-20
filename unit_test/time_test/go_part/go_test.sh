#!/bin/bash

# Compilation du script Go
go build -o myscript ./go_part/test_questions.go

# Création d'un fichier de test
echo "yes" > data_answer.txt
echo "yes" >> data_answer.txt
echo "yes" >> data_answer.txt

# Exécution du script à tester
output=$(./myscript data_answer.txt)

# Vérification du résultat
expected_output="Score final du patient : 120"
if [[ "$output" == "$expected_output" ]]; then
    echo "Go Test 1 réussi -> $output"
else
    echo "Go Go Test 1 échoué -> Le résultat obtenu ($output) ne correspond pas au résultat attendu ($expected_output)"
fi

# Nettoyage des fichiers temporaires
rm -f data_answer.txt


################################

echo "yes" > data_answer.txt
echo "yes" >> data_answer.txt
echo "no" >> data_answer.txt

output=$(./myscript data_answer.txt)

expected_output="Score final du patient : 110"
if [[ "$output" == "$expected_output" ]]; then
    echo "Go Test 2 réussi -> $output"
else
    echo "Go Test 2 échoué -> Le résultat obtenu ($output) ne correspond pas au résultat attendu ($expected_output)"
fi

rm -f data_answer.txt


################################

echo "yes" > data_answer.txt
echo "no" >> data_answer.txt
echo "no" >> data_answer.txt

output=$(./myscript data_answer.txt)

expected_output="Score final du patient : 10"
if [[ "$output" == "$expected_output" ]]; then
    echo "Go Test 3 réussi -> $output"
else
    echo "Go Test 3 échoué -> Le résultat obtenu ($output) ne correspond pas au résultat attendu ($expected_output)"
fi

rm -f data_answer.txt


################################

echo "yes" > data_answer.txt
echo "no" >> data_answer.txt
echo "yes" >> data_answer.txt

output=$(./myscript data_answer.txt)

expected_output="Score final du patient : 20"
if [[ "$output" == "$expected_output" ]]; then
    echo "Go Test 4 réussi -> $output"
else
    echo "Go Test 4 échoué -> Le résultat obtenu ($output) ne correspond pas au résultat attendu ($expected_output)"
fi

rm -f data_answer.txt

################################

echo "no" > data_answer.txt
echo "no" >> data_answer.txt
echo "no" >> data_answer.txt

# Exécution du script à tester
output=$(./myscript data_answer.txt)

# Vérification du résultat
expected_output="Score final du patient : 0"
if [[ "$output" == "$expected_output" ]]; then
    echo "Go Test 5 réussi -> $output"
else
    echo "Go Test 5 échoué -> Le résultat obtenu ($output) ne correspond pas au résultat attendu ($expected_output)"
fi

# Nettoyage des fichiers temporaires
rm -f data_answer.txt


################################

echo "no" > data_answer.txt
echo "no" >> data_answer.txt
echo "yes" >> data_answer.txt

output=$(./myscript data_answer.txt)

expected_output="Score final du patient : 10"
if [[ "$output" == "$expected_output" ]]; then
    echo "Go Test 6 réussi -> $output"
else
    echo "Go Test 6 échoué -> Le résultat obtenu ($output) ne correspond pas au résultat attendu ($expected_output)"
fi

rm -f data_answer.txt


################################

echo "no" > data_answer.txt
echo "yes" >> data_answer.txt
echo "yes" >> data_answer.txt

output=$(./myscript data_answer.txt)

expected_output="Score final du patient : 110"
if [[ "$output" == "$expected_output" ]]; then
    echo "Go Test 7 réussi -> $output"
else
    echo "Go Test 7 échoué -> Le résultat obtenu ($output) ne correspond pas au résultat attendu ($expected_output)"
fi

rm -f data_answer.txt


################################

echo "no" > data_answer.txt
echo "yes" >> data_answer.txt
echo "no" >> data_answer.txt

output=$(./myscript data_answer.txt)

expected_output="Score final du patient : 100"
if [[ "$output" == "$expected_output" ]]; then
    echo "Go Test 8 réussi -> $output"
else
    echo "Go Test 8 échoué -> Le résultat obtenu ($output) ne correspond pas au résultat attendu ($expected_output)"
fi

rm -f myscript data_answer.txt

