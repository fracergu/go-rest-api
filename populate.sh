#!/bin/bash

for i in {1..100}
do
   curl -X POST http://localhost:8000/api/v1/users \
        -H 'Content-Type: application/json' \
        -d "{
            \"name\": \"Nombre$i\",
            \"surname\": \"Apellido$i\",
            \"username\": \"NombreDeUsuario$i\",
            \"email\": \"correo$i@example.com\",
            \"password\": \"contraseña$1\"
        }"
done
