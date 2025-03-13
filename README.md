# Aprendiendo Go!

## Tu primer proyecto: hello-world

Para inicializar un proyecto en ``go`` debemos de, antes de nada, **instalarlo**. 

```
En la página oficial:
https://go.dev/dl/

Encontrarás los enlaces disponibles dependiendo del sistema operativo que tengas.
```

Una vez descargado e instalado, podemos crear una carpeta llamada, por ejemplo ``hello-world``, mediante el siguiente comando:

```
mkdir hello-world
```

Nos situamos dentro de la carpeta:

```
cd hello-world
```

Y ejecutamos el siguiente comando:

```
go mod init hello-world
```

Esto creará un fichero llamado ``go.mod``, que aparecerá dentro de nuestro proyecto cuando lo abramos en nuestro editor elegido.
Para este caso, utilizaremos ``IntelliJ``.



````
Puedes probar un breve ejemplo en ``go`` utilizando el código en esta página

https://gobyexample.com/hello-world
````

## Aprendiendo Fiber

### ¿Qué es Fiber?

Fiber es un **framework** utilizando por Go para **desarrollo web**

### Primer script con Fiber

```
package main

import "github.com/gofiber/fiber/v2"

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	app.Listen(":3000")
}

```

Aquí tenemos un ejemplo básico.

Según ChatGPT, la siguiente instrucción:

````
package main
````

Indica **cuál es el punto de entrada de nuestra aplicación**. Es decir, sería el equivalente a ``main.tsx`` en una aplicación de React.

Por otro lado:

```
import "github.com/gofiber/fiber/v2"
```

Nos permite **acceder** a las funciones de **Fiber** para utilizarlas y desarrollar los aspectos generales y principales de la misma.
Es decir, sería igual que en este código:

````
const express = require('express')
const app = express()
const port = 3000

app.get('/', (req, res) => {
  res.send('Hello World!')
})

app.listen(port, () => {
  console.log(`Example app listening on port ${port}`)
})
````

Utilizar la instrucción ``const express = require('express')``.

En ``Go``, al parecer, no es necesario **declarar la variable**; es decir, no es necesario que tenga un prefijo como
`const`, `let` o `var`. **Go infiere el tipo automáticamente**, por lo que si tenemos una instrucción como:

```
	app := fiber.New()
```

Ya sabe que `app` es una variable instanciada de la clase `New`.

# hello-world-go
