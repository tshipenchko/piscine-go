#!/usr/bin/env bash
curl -s https://01.alem.school/assets/superhero/all.json | jq ".[] | select(.id==70) | .name"
