curl "http://localhost:8080/time/Zone" \
-H "Content-Type: application/json" \
-H "Authorization: Bearer $MICRO_API_TOKEN" \
-d '{
  "location": "London"
}'