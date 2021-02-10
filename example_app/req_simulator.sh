#!/bin/bash

for i in {0..10000};
do
   echo  "Post $i" | faas-cli invoke sequence-orchestrator;
   sleep 5;
done;
