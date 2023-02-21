echo $HERO_ID
curl -s https://01.alem.school/assets/superhero/all.json | jq -r "[.[] | select(.id==$HERO_ID)][0] | .connections.relatives"
