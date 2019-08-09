#!/bin/bash
go install .
sudo packetdemo &
pid=$!
sleep 1
sudo ifconfig utun2 10.1.0.10 10.1.0.20 up
trap "sudo kill $pid" INT TERM
wait $pid
