curl "http://localhost:8080/cache/Set" \
-H "Content-Type: application/json" \
-H "Authorization: Bearer $MICRO_API_TOKEN" \
-d '{
  "key": "foo",
  "value": "bar"
}'