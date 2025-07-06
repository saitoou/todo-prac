
### JWTのペイロードだけ確認（Base64URLデコード）
```
echo '<token>' | cut -d '.' -f2 | base64 -d
```

### ユーザー登録
```
curl -X POST http://localhost:8080/api/v1/auth/signup \
  -H "Content-Type: application/json" \
  -d '{
    "name": "Taro",   
    "email": "taro@example.com",    
    "password": "password123"           
  }'
```  

```
curl -X POST http://localhost:8080/api/v1/todos \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer <TOKEN>" \
  -d '{
    "title": "買い物リスト",
    "content": "卵、牛乳、パン"
  }'
```