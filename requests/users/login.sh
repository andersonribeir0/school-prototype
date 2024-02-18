#!/bin/bash

# URL do servidor
SERVER_URL="http://localhost:8080"

# Endpoint de login
LOGIN_ENDPOINT="$SERVER_URL/login"

# Fazendo a requisição de login e capturando a resposta completa
LOGIN_RESPONSE=$(curl -s -X POST -H "Content-Type: application/json" -d @login.json $LOGIN_ENDPOINT)

# Tentativa de extrair o token da resposta
TOKEN=$(echo $LOGIN_RESPONSE | jq -r '.body.token')

echo "Login response: $LOGIN_RESPONSE"
echo "Reveived token: $TOKEN"

# Verifica se o token foi recebido
if [ "$TOKEN" == "null" ] || [ -z "$TOKEN" ]; then
  echo "Login failure."
  exit 1
fi

# Endpoint /me
ME_ENDPOINT="$SERVER_URL/auth/me"

# Fazendo a requisição para /me usando o token e capturando a resposta
ME_RESPONSE=$(curl -s -X GET -H "Authorization: Bearer $TOKEN" $ME_ENDPOINT)

# Verifica se a resposta contém uma mensagem de erro
if echo "$ME_RESPONSE" | jq -e '.error' > /dev/null; then
  echo "Error accessing $ME_ENDPOINT: $ME_RESPONSE"
  exit 1
else
  echo "Response /me: $ME_RESPONSE"
fi

# Alteração aqui: abrir a URL no navegador com o token JWT
AUTH_URL="$SERVER_URL/auth?token=$TOKEN"
echo "Opening $AUTH_URL in the default web browser..."
open "$AUTH_URL"
