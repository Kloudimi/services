curl "http://localhost:8080/user/Create" \
-H "Content-Type: application/json" \
-H "Authorization: Bearer $MICRO_API_TOKEN" \
-d '{
  "email": "joe@example.com",
  "id": "usrid-1",
  "password": "mySecretPass123",
  "username": "usrname-1"
}'