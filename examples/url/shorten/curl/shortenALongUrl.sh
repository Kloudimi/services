curl "http://localhost:8080/url/Shorten" \
-H "Content-Type: application/json" \
-H "Authorization: Bearer $MICRO_API_TOKEN" \
-d '{
  "destinationURL": "https://mysite.com/this-is-a-rather-long-web-address"
}'