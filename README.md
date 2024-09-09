# ARREGLOS
ARREGLOS

Este código en ambos lenguajes, Python y Go, implementa una estructura (arreglo) para gestionar las ventas mensuales de tres departamentos: Ropa, Deportes y Juguetes. La funcionalidad incluye la capacidad de insertar y actualizar ventas para un departamento y mes específicos, buscar ventas por departamento y mes, eliminar ventas estableciendo el monto a cero, y mostrar el estado completo de las ventas organizadas por departamento y mes.

python:

Método insertar_ventas:

El método insertar_ventas sirve para actualizar el monto de ventas para un departamento y un mes específicos. Primero, el método valida si el nombre del departamento proporcionado es válido, comprobando que esté en la lista de departamentos definidos. Si el departamento no es válido, se imprime un mensaje de error y el método termina para evitar errores. Luego, el método verifica que el valor del mes esté dentro del rango permitido, que es de 0 a 11, correspondiente a los 12 meses del año. Si el mes no está en este rango, se imprime un mensaje de error y el método finaliza.
Antes de realizar cualquier actualización, el método imprime el estado actual de las ventas utilizando el método mostrar_ventas, lo que permite visualizar cómo están los datos antes del cambio. Luego, se utiliza el método index (índice) para encontrar la posición del departamento en la lista de departamentos, lo que permite identificar la sublista correcta dentro de ventas donde se deben hacer los cambios. Una vez encontrado el índice del departamento, el método actualiza el monto de ventas en la sublista correspondiente al mes especificado. Finalmente, el método imprime un mensaje confirmando que las ventas han sido actualizadas y muestra el nuevo estado de las ventas para verificar los cambios realizados.

Método buscar_ventas:

El método buscar_ventas sirve para recuperar el monto de ventas para un departamento y un mes específicos. Primero, valida si el nombre del departamento proporcionado está en la lista de departamentos válidos. Si el departamento no es válido, el método devuelve un mensaje de error indicando que el departamento no es válido. Luego, el método verifica que el valor del mes esté dentro del rango permitido (0 a 11). Si el mes está fuera de este rango, regresa un mensaje de error señalando que el mes no es válido. Una vez realizadas estas validaciones, el método determina el índice del departamento en la lista de departamentos utilizando el método index. Este índice se usa para acceder a la sublista correspondiente en ventas y obtener el monto de ventas para el mes especificado. Finalmente, el método retorna el monto de ventas encontrado para el departamento y mes indicados.

Método eliminar_ventas

El método eliminar_ventas permite eliminar el monto de ventas para un departamento y un mes específicos. Primero, verifica si el departamento proporcionado es válido, asegurándose de que esté en la lista de departamentos. Si el departamento no es válido, el método imprime un mensaje de error y finaliza para evitar modificaciones incorrectas. A continuación, valida que el mes esté dentro del rango permitido (0 a 11). Si el mes no está en este rango, el método imprime un mensaje de error y termina su ejecución. Antes de realizar la eliminación, el método imprime el estado actual de las ventas para proporcionar un antecedente antes de modificar los valores. Luego, utiliza el método index para encontrar la posición del departamento en la lista de departamentos. Con este índice, se accede a la sublista correspondiente en ventas y se establece el monto de ventas para el mes especificado a 0, eliminando el monto de ventas para ese mes. Finalmente, el método imprime un mensaje confirmando que las ventas han sido eliminadas y muestra el nuevo estado de las ventas para verificar los cambios.

Método mostrar_ventas

El método mostrar_ventas está diseñado para imprimir el estado actual de todas las ventas organizadas por departamento y mes. Recorre la lista de departamentos y, para cada departamento, imprime el nombre del departamento seguido de los montos de ventas para cada uno de los 12 meses del año. Este método proporciona una vista completa y clara de los datos de ventas almacenados en el atributo ventas, facilitando la visualización y verificación de las ventas registradas.

Método buscar_venta_especifica
El método buscar_venta_especifica permite buscar un monto específico de ventas en todos los departamentos y meses. Recorre la lista de departamentos y, para cada departamento, examina los montos de ventas para cada mes del año. Si el monto de ventas coincide con el monto buscado, el método agrega una lista con el nombre del departamento y el mes correspondiente a una lista de resultados. Al finalizar la búsqueda, el método regresa la lista de resultados que contiene todas las ubicaciones (departamento y mes) donde se encontró el monto de ventas buscado. Si no se encuentran coincidencias, la lista de resultados estará vacía.

golang (GO):

A diferencia de Python en GO se debe de iniciar con la definición del paquete main que es donde se desarrolla todo el programa y se debe de inportar la librería "fmt" que es la que te permite hacer uso de los prints oara poder mostrar información, entre otros. otra diferencia notable es la sintaxis ya que para definir una variable debes de usar la palabra reservada "var" o "func" para el caso de una función, además del uso de llaves {} para poner las sangrias en las funciones o siclos al igual que otros lenguajes como java. Para definir un arreglo en GO se le debe de poner el tipo de datos que contendrá en python no es necesario definirlo.

Método InsertarVentas 

El método insertarVentas sirve para actualizar el monto de ventas para un departamento y un mes específicos. Primero, verifica si el nombre del departamento es válido mediante el método departamentoValido, que recorre la lista de departamentos y devuelve true (verdadero) si el departamento está presente. Si el departamento no es válido, imprime un mensaje de error y termina el método. Luego, valida que el mes esté dentro del rango permitido (0 a 11) usando una condición if. Si el mes no es válido, imprime un mensaje de error y termina la ejecución. Antes de actualizar el monto, el método imprime el estado actual de las ventas llamando a mostrarVentas. Después, encuentra el índice del departamento en la lista de departamentos utilizando el método indiceDepartamento, que recorre la lista y devuelve la posición del departamento. Con este índice, el método actualiza el monto en la sublista correspondiente de ventas. Finalmente, imprime un mensaje confirmando la actualización y muestra el estado actualizado de las ventas.

Método buscarVentas

El método buscarVentas se utiliza para recuperar el monto de ventas de un departamento y un mes específicos. Primero, valida que el nombre del departamento sea uno de los válidos utilizando el método departamentoValido, que revisa si el departamento está en la lista. Si el departamento no es válido, imprime un mensaje de error y retorna -1. Luego, verifica que el mes esté dentro del rango permitido (0 a 11) con una condición if. Si el mes no es válido, imprime un mensaje de error y retorna -1. Con el índice del departamento encontrado mediante indiceDepartamento, el método accede a la sublista correspondiente en ventas para obtener el monto de ventas del mes especificado y lo regresa.

Método eliminarVentas

El método eliminarVentas sirve para eliminar el monto de ventas de un departamento y un mes específicos, estableciendo el monto a cero. Primero, valida que el nombre del departamento sea uno de los válidos usando el método departamentoValido. Si el departamento no es válido, imprime un mensaje de error y termina el método. Luego, verifica que el mes esté dentro del rango permitido (0 a 11) con una condición if. Si el mes no es válido, imprime un mensaje de error y finaliza la ejecución. Antes de eliminar el monto, imprime el estado actual de las ventas utilizando mostrarVentas. Luego, encuentra el índice del departamento con indiceDepartamento y establece el monto de ventas en la sublista correspondiente a cero. Finalmente, imprime un mensaje confirmando la eliminación y muestra el estado actualizado de las ventas.

Método mostrarVentas

El método mostrarVentas imprime el estado actual de todas las ventas organizadas por departamento y mes. Recorre la lista de departamentos usando un bucle for y, para cada departamento, imprime su nombre seguido de los montos de ventas correspondientes a cada uno de los 12 meses del año. Utiliza fmt.Printf para formatear la salida y fmt.Println para imprimir cada monto de ventas y separar la información entre departamentos.

Método buscarVentaEspecifica

El método buscarVentaEspecifica busca todas las coincidencias de un monto específico de ventas en todos los departamentos y meses. Recorre la lista de departamentos usando un bucle for, y para cada departamento, revisa los montos de ventas para cada mes usando otro bucle for. Si el monto de ventas coincide con el monto buscado, el método agrega una lista con el nombre del departamento y el mes a una lista de resultados. Al finalizar la búsqueda, el método retorna la lista de resultados. Si no se encuentran coincidencias, la lista de resultados estará vacía.

Métodos auxiliares 

departamentoValido e indiceDepartamento

El método departamentoValido verifica si un departamento dado es válido revisando si está en la lista de departamentos. Recorre la lista de departamentos con un bucle for y retorna true si el departamento se encuentra en la lista. Si no se encuentra, retorna false.
El método indiceDepartamento encuentra el índice de un departamento en la lista de departamentos. Utiliza un bucle for para comparar el nombre del departamento dado con cada elemento en la lista. Si encuentra una coincidencia, retorna el índice del departamento. Si no encuentra ninguna coincidencia, retorna -1.

diferencias entre Python y GO

En Python, las clases se definen con class y los métodos con def. Los atributos se inicializan en el constructor __init__, y las listas y diccionarios son estructuras de datos comunes. Los métodos acceden a los atributos usando self, y se utilizan métodos como index para buscar elementos en listas. Los errores se manejan con excepciones y print. Python es dinámico en cuanto a tipos de datos y usa print para salida con formateo opcional.

En Go, las estructuras se definen con struct y los métodos se asocian usando un receptor de tipo. El constructor es una función como NewVentas, y las colecciones se manejan con slices. La búsqueda en slices se hace con bucles for, y los errores se gestionan mediante condiciones if y fmt.Println. Go es estático en cuanto a tipos y usa fmt.Printf para salida formateada. Los arrays en Go tienen tamaño fijo, y los métodos se definen con el receptor de la estructura.
