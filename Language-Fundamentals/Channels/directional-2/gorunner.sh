#!/bin/bash

echo "GO Program Runner"
echo "================="

read -p "How many times do you want to run the program: " count

for((i = 1; i <= count; i++)); do

	echo
	echo "Run #$i"
	echo "----------------------------------"
	go run ./main.go
	echo "----------------------------------"
done
echo "Completed $count runs!"
