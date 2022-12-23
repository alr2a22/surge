threshols='[
    {
        "coefficient": 1,
        "threshold": 0
    },
    {
        "coefficient": 1.2,
        "threshold": 5
    },
    {
        "coefficient": 1.4,
        "threshold": 10
    },
    {
        "coefficient": 1.8,
        "threshold": 20
    },
    {
        "coefficient": 2,
        "threshold": 35
    },
    {
        "coefficient": 2.2,
        "threshold": 55
    }
]'

user=admin
password=admin

res_login=$(curl -XPOST -s localhost:3000/login -H  "Content-Type: application/json" -d "{ \"username\": \"$user\", \"password\": \"$password\" }")
token=$(echo $res_login | jq .token | tr -d '"')
echo token is $token

send () {
  curl -XPOST localhost:3000/thresholds -H "Content-Type: application/json" -H  "Authorization: Bearer $token" -d "{ \"coefficient\": $1, \"threshold\": $2 }"
}

echo "$threshols" | jq -r '.[]|[.coefficient, .threshold] | @tsv' |
while IFS=$'\t' read -r coefficient threshold; do
send $coefficient $threshold
done