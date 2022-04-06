IFS=$'\n' read -d '' -r -a lines < resources.txt

for i in "${lines[@]}"
do
   echo "$i"
   export URL=$i
   export NAME=`echo $i| sed 's/\///g'|sed 's/://g'`
   envsubst < "bombardier.yaml" > "bombardier_dst.yaml"
   kubectl apply -f bombardier_dst.yaml
done
