curl "http://localhost:8080/currency/History" \
-H "Content-Type: application/json" \
-H "Authorization: Bearer $MICRO_API_TOKEN" \
-d '{
  "code": "USD",
  "date": "2021-05-30"
}'