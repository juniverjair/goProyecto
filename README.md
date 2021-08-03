# Proyecto Go

## Propuesta del proyecto

Traductor de texto por país o ciudad, donde selecciono un país o ciudad de origen y uno de destino finalmente traduzco dependiendo el idioma del origen y destino.

Se consumiran 2 API's:

1. Ciudades y lenguajes
   
    [https://restcountries.eu](https://restcountries.eu)

2. Google Translate
   
    [https://cloud.google.com/translate](https://cloud.google.com/translate/docs/basic/quickstart)


## Ejecución

### Ejecutar el proyecto

```bash
go mod init master.project/go
go get $(librerias espacificadas en el main.go)
go run main.go
```

### Servicios

```bash
GET http://localhost:9000/getCountries
GET http://localhost:9000/getTranslate Params: text, lang
```

### Url Base

```bash
http://localhost:9000/home
```

### Visualización en ejecución

![home-screen](/static/home.png)