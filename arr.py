class Ventas:
    def __init__(self):
        self.ventas = [
            [5000, 4000, 6000, 7000, 5500, 8000, 6500, 7200, 6800, 5900, 7300, 7700],  # Ropa
            [3000, 3500, 2500, 4000, 3200, 3700, 3400, 3900, 3600, 3300, 3100, 3200],  # Deportes
            [7000, 7500, 8000, 8500, 7800, 8200, 7900, 8600, 8300, 8000, 7800, 8100]   # Juguetes
        ]
        self.departamentos = ['Ropa', 'Deportes', 'Juguetes']

    def insertar_ventas(self, departamento, mes, monto):
        if departamento not in self.departamentos:
            print("Departamento no válido.")
            return
        if mes < 0 or mes > 11:
            print("Mes no válido. Debe estar entre 0 y 11.")
            return

        print("Estado antes de insertar ventas:")
        self.mostrar_ventas()

        indice = self.departamentos.index(departamento)
        self.ventas[indice][mes] = monto
        print(f"Ventas de {departamento} en mes {mes+1} actualizadas a {monto}.")

        print("Estado después de insertar ventas:")
        self.mostrar_ventas()

    def buscar_ventas(self, departamento, mes):
        if departamento not in self.departamentos:
            return "Departamento no válido."
        if mes < 0 or mes > 11:
            return "Mes no válido. Debe estar entre 0 y 11."

        indice = self.departamentos.index(departamento)
        return self.ventas[indice][mes]

    def eliminar_ventas(self, departamento, mes):
        if departamento not in self.departamentos:
            print("Departamento no válido.")
            return
        if mes < 0 or mes > 11:
            print("Mes no válido. Debe estar entre 0 y 11.")
            return

        print("Estado antes de eliminar ventas:")
        self.mostrar_ventas()

        indice = self.departamentos.index(departamento)
        self.ventas[indice][mes] = 0
        print(f"Ventas de {departamento} en mes {mes+1} eliminadas.")

        print("Estado después de eliminar ventas:")
        self.mostrar_ventas()

    def mostrar_ventas(self):
        for i in range(len(self.departamentos)):
            departamento = self.departamentos[i]
            print(f"Ventas en {departamento}:")
            for mes in range(12):
                print(f"Mes {mes+1}: {self.ventas[i][mes]}")
            print()  # Línea en blanco para separar departamentos

    def buscar_venta_especifica(self, monto):
        resultados = []
        for i, departamento in enumerate(self.departamentos):
            for mes in range(12):
                if self.ventas[i][mes] == monto:
                    resultados.append((departamento, mes))
        return resultados
ventas = Ventas()
ventas.insertar_ventas('Ropa', 0, 5500)  
ventas.insertar_ventas('Deportes', 1, 3200)  
ventas.insertar_ventas('Juguetes', 2, 7500)  

monto_buscado = 5000
resultados = ventas.buscar_venta_especifica(monto_buscado)
if resultados:
    print(f"Ventas de {monto_buscado} encontradas en:")
    for resultado in resultados:
        departamento, mes = resultado
        print(f"Departamento: {departamento}, Mes: {mes + 1}")
else:
    print(f"No se encontraron ventas con el monto {monto_buscado}.")

ventas.eliminar_ventas('Deportes', 1)  


