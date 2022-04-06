IFS=$'\n' read -d '' -r -a lines < resources.txt

for i in "${lines[@]}"
do
   echo "$i"
   export URL=$i
   docker run --platform linux/amd64 -d  alpine/bombardier -c 1000 -d 60000h -l $URL 
 done
