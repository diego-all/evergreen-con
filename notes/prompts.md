# Prompts



0. Generar prompts preconstruidos para Golang, arquitectura MVC y templates html

        Requiero algunos prompts preconstruidos para generar una aplicación en lenguaje Golang utilizando una arquitectura MVC. Requiero un frontend para capturar datos del lado cliente utilizando el motor de templates de Golang.​


1. Crear archivo Dockerfile

        Requiero un Dockerfile para crear una base de datos mysql ultima version cuya variable de entorno MYSQL_ROOT_PASSWORD sea igual a password y copie el archivo up.sql al momento de crear el contenedor.


2. Crear archivo up.sql (Definición del modelo físico de datos)

        Requiero construir un modelo de datos fisico para las tablas device_types cuyas columas son (id, name),tabla devices cuyas columnas son (id, name, location, parameters, type_id, model), considerando que type_id es llave primaria en la tabla device_types y tabla applications cuyas columnas son (id, identifier, name, port, status, type, language). Teniendo en cuenta que todas las tablas una llave primaria id autoincrementable y llevan campos de creado y actualizado para auditoria.


3. Crear conexión con la base de datos

        Requiero crear la conexion a una base de datos mysql en golang utilizando el driver correspondiente utilizando un par de funciones, una llamada connectDatabase que reciba como parametro un string conexion DSN para conectarse a una base de datos mysql, dicha funcion invoca una funcion llamada TestDB que ejecuta un ping hacia la base de datos en caso de ser exitoso retorna un mensaje Pinged database successfully! de lo contrario manejar el error.


4. Generar datos dummy para tabla devices

        Podrias generarme data dummy (5 inserts por tabla.) para las tablas en mysql device_types, devices y applications basado en el modelo fisico proporcionado, considerando que es una aplicación de agricultura moderna. 


5. Generar estructura de carpetas del proyecto

        Requiero generar la estructura de carpetas para una aplicacion en golang utilizando el patron MVC con los siguientes paquetes: assets (imagenes), config, controllers (application controller,  devicecontroller, homecontroller), db-artifacts, entities, models (applicationmodel, deviceTypemodelo, devicemodel), views (application, device), cada una de las vistas con index.html, create.html, detail.html, edit.html.


6. Generacion de modelos (estructura)

        Requiero generar el modelo para devices y applications alineado con     las opraciones CRULD. Create, Read, Update, List, Delete para este  modelo fisico de datos.


7. **Utilizando técnica chunking** (Sugerencias de Copilot)

        How can I implement the Create operation for the Device model in Go?
        How can I implement the Read operation for the Device model in Go?
        How can I implement the Update operation for the Device model in Go?
        How can I implement the Delete operation for the Device model in Go?
        How can I implement the List operation for the Device model in Go?
        To implement the Create operation for the Device model in Go, you can follow these steps:


8. Generar las vistas utilizando motor de templates de Go

        Requiero que generes las vistas utilizando el motor de templates de     Golang para estas operaciones Create, Read, Update, Delete y List   alineado con los archivos index.html, create.html, detail.html, edit.     html. Utilizando boostrap generar un navbar para index.html que este         en todas las plantillas, al igual que el footer. Adicionar estetica     con fondos y colores de boostrap con colores adecuados para una     aplicacion ecologica.


9. Generar controladores

        Requiero que generes los controllers para estas vistas y retornar al cliente la plantilla a renderizar con la data correspondiente.


