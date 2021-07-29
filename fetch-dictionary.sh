#!/bin/sh
curl -Lo frek.zip http://kaino.kotus.fi/sanat/taajuuslista/vns_frek.zip
unzip frek.zip
cat vns_frek.txt | awk '{print $3}' | sort -d > dictionary.txt 

