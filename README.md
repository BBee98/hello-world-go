# Aprendiendo Go!

## 1. Tu primer proyecto: hello-world

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

## 2. Aprendiendo Fiber

### 2.1 ¿Qué es Fiber?

Fiber es un **framework** utilizando por Go para **desarrollo web**

### 2.2 Primer script con Fiber

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

Es importante que **instalemos el paquete de `"github.com/gofiber/fiber/v2"` para poder utilizarlo.** Para ello, escribimos
en la terminal:

```go get github.com/gofiber/fiber/v2```

Y ya tendremos el paquete listo para utilizar.

Además, en ``Go``, al parecer, no es necesario **declarar la variable**; es decir, no es necesario que tenga un prefijo como
`const`, `let` o `var`. **Go infiere el tipo automáticamente**, por lo que si tenemos una instrucción como:

```
	app := fiber.New()
```

Ya sabe que `app` es una variable instanciada de la clase `New`.

### 2.3 Estableciendo los ficheros .html como templates

Ya hemos visto que, como con ``express``, podemos servir **ficheros estáticos**, pero dado que, hasta ahora, hemos servido **texto plano**, tenemos que **configurarlo**.

Vamos a aprender varias cosas:

1. A establecer la **configuración inicial** de nuestro servidor de Go.
2. A exportar **funciones**.

#### 2.3.1 Creando el fichero ``config.go``

Creamos el directorio ``config`` en la raíz del proyecto y, dentro de éste, el fichero ``config.go``.

Primero crearemos la **función** donde iniciaremos toda la configuración; la llamaremos `InitConfig`:

````
package config

func InitConfig() { }
````

#### ¿Por qué debemos usar la declaración ``package config``?

Al parecer, al contrario que ocurre con otros lenguajes de programación, ``Go`` no se organiza por archivos (aunque realmente nosotros creemos esos ficheros),
sino por ``paquetes``. Para ayudar a ``Go`` y poder usar las funciones o las variables que exportemos, deberemos utilizar esta instrucción.

#### ¿Cómo exportamos una función?

Para exportar una función tan solo debemos **declararla con la primera letra en mayúsculas**; de esta manera, podremos usar ``InitConfig`` desde el fichero
``main.go``

### Continuación de 2.3.1

Una de las cosas principales que queremos configurar es, como dijimos anteriormente, la posibilidad de servir **ficheros estáticos**. Pero no queremos generarlos **manualmente**
con ``go``, sino que queremos que sirva **ficheros ``html``**. 

Al igual que como ocurre con los frameworks de ``React``, ``Vue``, etc, necesitamos "instalar" un package que nos ayude a realizar esta configuración:

```
go get github.com/gofiber/template/html/v2
```

Y, una vez instalado, lo importamos en el fichero:

```
package config

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
)

func InitConfig() {


}
```

```
Nota: Ocurre igual con el paquete de "github.com/gofiber/fiber/v2". Recuerda que en un paso previo ya lo instalamos
```

Ahora podemos iniciar la configuración para utilizar los ficheros .html:

````
package config

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
)

func InitConfig() {

	engine := html.New("./views", ".html")

	app := fiber.New(fiber.Config{
		Views: engine,
		// Default global path to search for views (can be overriden when calling Render())
	})

	app.Get("/", func(c *fiber.Ctx) error {
		// Render index template
		return c.Render("index", fiber.Map{})
	})

	app.Listen(":3000")

}
````

Levantamos el servidor con la instrucción `go run main.go` y ¡listo!.


