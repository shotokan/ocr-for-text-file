# ocr-for-text-file

ocr-for-text-file levanta un servidor web y queda a la espera de que se le envíe un archivo como el de entradas.txt.

Cada entrada en el archivo se conforma de 4 líneas y cada línea tiene 27 caracteres. Las primeras 3 líneas de cada
entrada contienen un número de cuenta escrito usando pipes y guiones bajos, y la cuarta línea está en blanco.
Cada número de cuenta debe tener 9 dígitos, todos los cuales deben estar en el rango 0-9. Un archivo normal
contiene alrededor de 500 entradas.

El OCR lee y genera una cadena con los números y posteriormente estos son validados en el endpoint usando la fórmula:

account number: 3 4 5 8 8 2 8 6 5
position name: d9 d8 d7 d6 d5 d4 d3 d2 d1
checksum calculation:
(1 _ d1 + 2 _ d2 + 3 _ d3 + … + 9 _ d9) mod 11 = 0

La respuesta del endpoint es la siguiente:

POST /accounts/validate

```json
[
  {
    "id": "192ac75e-be3b-4f7e-a67a-b17c23ca29e7",
    "number": "000000000",
    "checksum": "OK"
  },
  {
    "id": "058c75ef-6f53-4b86-b7bf-40606a6df747",
    "number": "111111111",
    "checksum": "ERR"
  },
  {
    "id": "6902cc7c-36cd-436b-9b8e-0439b30a66cc",
    "number": "222222222",
    "checksum": "ERR"
  },
  {
    "id": "5455ac3e-23d9-4b95-9f2a-e13a8848c850",
    "number": "333333333",
    "checksum": "ERR"
  },
  {
    "id": "c30114e7-8d50-418d-93b8-ce2477576ae1",
    "number": "444444444",
    "checksum": "ERR"
  },
  {
    "id": "f7fe9661-bdfa-4830-bcaf-ecc7e4698287",
    "number": "555555555",
    "checksum": "ERR"
  },
  {
    "id": "d68fdc56-d858-4381-84f8-85932be9e4bf",
    "number": "666666666",
    "checksum": "ERR"
  },
  {
    "id": "ade34804-a245-4ed0-b0be-1a325724018c",
    "number": "777777777",
    "checksum": "ERR"
  },
  {
    "id": "57420b26-a1cb-44ce-bff6-3bc5921ad82a",
    "number": "888888888",
    "checksum": "ERR"
  },
  {
    "id": "02ba1917-9d77-4887-b92e-500a00e94990",
    "number": "999999999",
    "checksum": "ERR"
  },
  {
    "id": "07154383-7562-4f39-8174-bfa7de3d2aad",
    "number": "123456789",
    "checksum": "OK"
  },
  {
    "id": "2762067f-bfb8-470f-ba05-20d7b8190459",
    "number": "000000051",
    "checksum": "OK"
  },
  {
    "id": "88bede6c-dfaf-4e13-aa75-ff875b2f8f57",
    "number": "49006771?",
    "checksum": "ILL"
  },
  {
    "id": "ae952ce4-29be-4ae0-8626-0413f9ebfef1",
    "number": "123??678?",
    "checksum": "ILL"
  }
]
```

TODO: testis unitatios y de integración
