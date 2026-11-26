# Conversor Babilónico: de numeración sexagesimal a decimal en Go

[![Go](https://img.shields.io/badge/Go-1.22%2B-00ADD8?logo=go&logoColor=white)](https://go.dev/)
[![Pruebas](https://img.shields.io/badge/pruebas-go%20test-success)](#pruebas)
[![Cobertura](https://img.shields.io/badge/cobertura-85%25-green)](#pruebas)
[![Sin dependencias](https://img.shields.io/badge/dependencias-ninguna-lightgrey)](go.mod)

Programa de consola que convierte números escritos en el sistema de numeración babilónico (base 60) a su valor decimal. Resuelve el procedimiento explicado en el libro del profesor Peña para pasar de notación sexagesimal a decimal.

## Cómo funciona

1. Se divide la entrada en dos cadenas usando `.` como separador: la parte entera y la parte fraccionaria.
2. Cada parte se divide de nuevo usando `;` como separador entre dígitos.
3. Los dígitos se recorren y se multiplican por la potencia de 60 que corresponde a su posición: `60^0`, `60^1`... hacia la izquierda en la parte entera y `60^-1`, `60^-2`... hacia la derecha en la parte fraccionaria.
4. Cada dígito debe estar entre 1 y 59 y la notación debe ser correcta. En caso contrario el programa muestra una alerta y vuelve a pedir la entrada.

### Ejemplos

| Entrada | Cálculo | Resultado |
| :--- | :--- | :--- |
| `1;30` | 1 x 60 + 30 | `90` |
| `1;2;3` | 1 x 3600 + 2 x 60 + 3 | `3723` |
| `1;30.15` | 1 x 60 + 30 + 15 / 60 | `90.25` |
| `1.30` | 1 + 30 / 60 | `1.5` |

## Estructura

```
.
├── babilonico.go       # ParseSexagesimal: validación y conversión a decimal
├── babilonico_test.go  # Pruebas de la conversión y del ciclo de la consola
├── main.go             # Consola interactiva: pide números hasta escribir "salir"
└── go.mod
```

## Requisitos

- Go 1.22 o superior.

## Ejecución local

```bash
git clone git@github.com:KevinAndrey96/conversor-babilonico.git
cd conversor-babilonico
go run .
```

Sesión de ejemplo:

```
Número babilónico (ej. 1;30.15) o 'salir': 1;30.15
Valor decimal: 90.25
Número babilónico (ej. 1;30.15) o 'salir': 99
Alerta: integer part: digit out of range
Número babilónico (ej. 1;30.15) o 'salir': salir
```

## Pruebas

```bash
go test -race -cover ./...
```

Las pruebas cubren dígitos sueltos, varios dígitos, partes fraccionarias, espacios sobrantes y los errores de entrada vacía, dígito fuera de rango y notación inválida, además de la repetición de la solicitud tras una entrada inválida.

## Licencia

Sin licencia de código abierto. Todos los derechos reservados.
