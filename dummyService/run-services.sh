#!/bin/bash

# Lancer le premier service en arrière-plan
(cd call/ && go build -o call && ./call) &

# Lancer le deuxième service
(cd data/ && go build -o data && ./data)
