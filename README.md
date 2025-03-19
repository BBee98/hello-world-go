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
En mi caso, utilizaré `IntelliJ`, pero esta decisión no afecta en nada al desarrollo de la aplicación.


````
Puedes probar un breve ejemplo en ``go`` utilizando el código en esta página

https://gobyexample.com/hello-world
````

## 2. Introduciendo Fiber en nuestro proyecto

### 2.1 ¿Qué es Fiber?

Fiber es un **framework** utilizando por Go para **desarrollo web**

### 2.2 Primer script con Fiber

Aquí tenemos un ejemplo básico.

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

### Analizando el script 

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

Una de las cosas principales que queremos configurar es, como dijimos anteriormente, la posibilidad de servir **ficheros ``.html``**.
Una de las cosas que podemos configurar es **qué tipo de templates queremos utilizar**. ``Go`` soporta templates como ``.pug``, pero en este caso
queremos usar ``html``

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

Sin embargo, si hemos añadido css a nuestro fichero ``html``, comprobaremos que **éste no puede ser leído**. Eso es
porque debemos indicarle a ``Go`` que sirva como estáticos los ficheros que le indiquemos. 

Añadamos la siguiente línea:

``	app.Static("/", "./views") ``

Y así ``Go`` podrá servir los ficheros `.css` y que éstos sean reconocidos.

El fichero final quedaría así:

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

	app.Static("/", "./views")

	app.Get("/", func(c *fiber.Ctx) error {
		// Render index template
		return c.Render("index", fiber.Map{})
	})

	app.Listen(":3000")

}
````

Solo queda importarlo en nuestro fichero ``main.go``:

````
package main

import (
	"hello-world/config"
)

func main() {
	config.InitConfig()
}

````

Levantamos el servidor con la instrucción `go run main.go` y ¡listo!.


### Notas importantes

- Como dijimos anteriormente, ``Go`` se mueve por **packages (paquetes), no por directorios**. Cuando efectuamos una **importación** de
una función o de una variable, es importante que tengamos en cuenta que, al contrario de lo que ocurre con otros lenguajes, **no tenemos que mencionar
el fichero que estamos importando**. **Solo necesitamos llamar al directorio en el cual se encuentra el fichero**. En el caso de la función
``InitConfig``, por ejemplo, el **directorio** se llama ``config`` (que, por suerte o por desgracia, coincide con el fichero ``config.go`` en cuanto a nombre).
Si el directorio se llamara ``configuration``, deberíamos usar la instrucción de esta manera:

````
package main

import (
	"hello-world/configuration"
)

func main() {
	config.InitConfig()
}

````

## 3. Routing en Fiber

Si atendemos a la documentación de ``Fiber``:

```
https://docs.gofiber.io/guide/routing
```

Nos indica que ``Fiber`` nos proporciona la siguiente lista de métodos para configurar nuestras rutas:

````
// HTTP methods
func (app *App) Get(path string, handlers ...Handler) Router
func (app *App) Head(path string, handlers ...Handler) Router
func (app *App) Post(path string, handlers ...Handler) Router
func (app *App) Put(path string, handlers ...Handler) Router
func (app *App) Delete(path string, handlers ...Handler) Router
func (app *App) Connect(path string, handlers ...Handler) Router
func (app *App) Options(path string, handlers ...Handler) Router
func (app *App) Trace(path string, handlers ...Handler) Router
func (app *App) Patch(path string, handlers ...Handler) Router
````

Esto puede ser un poco "confuso", así que vamos a explicarlo **a la vez que modificamos un poco nuestro código**.

### 3.1 Configurando la ruta base "/"

Si nos fijamos en el ejemplo que nos facilita ``Fiber``:


```
// Simple GET handler
app.Get("/api/list", func(c *fiber.Ctx) error {
  return c.SendString("I'm a GET request!")
})
```

Nos indica que el método `Get` requiere de dos atributos:

- El primero: la **definición de la ruta**.
- El segundo: el ``handler``, que se trata de un callback de este mismo tipo
que decidirá _qué hacer_ cuando se acceda a esta ruta.

Vemos que el método lo usa desde la variable `app`. Es decir, tenemos que usar 
**el objeto `app` que creamos al principio de todo para configurar la aplicación**:



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

Sin embargo, **hay un problema**, y es que la variable **está declarada localmente**. Recordemos que ``Fiber``
interpreta que una variable o función es **exportable cuando ésta está en mayúscula**.

Por tanto, tenemos que hacer un pequeño cambio en el código. 

1º Tenemos que **exportar la variable `app`**:

```
package config

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
)

var App *fiber.App

func InitConfig() {

	engine := html.New("./views", ".html")

	App = fiber.New(fiber.Config{
		Views: engine,
		// Default global path to search for views (can be overriden when calling Render())
	})

	App.Static("/", "./views")

	App.Get("/", func(c *fiber.Ctx) error {
		// Render index template
		return c.Render("index", fiber.Map{})
	})

	App.Listen(":3000")

}
```

Para ello, **declaramos** la variable **fuera** de la función:

``var App *fiber.App``

- **¿Por qué usamos el asterisco ( * )?**. Porque esto nos permite utilizar un mero **puntero**, 
siendo así **más eficientes** al utilizar **menos recursos**. Es decir, en lugar de **hacer una copia** de todo el
objeto `App`, lo que hacemos es decir "oye, solo quiero la **referencia** para poder usar los métodos".

- **¿Por qué usamos ``var`` y no ``const``? Para usar ``const``necesitamos inicializarla con algún valor, y no es
lo que precisamos en este momento. Queremos, simplemente, declararla **fuera** para poder exportarla.

2º Lo siguiente que hemos hecho es cambiar **la manera de realizar la asignación de la variable:**

````
App = fiber.New(fiber.Config{
		Views: engine,
		// Default global path to search for views (can be overriden when calling Render())
	})
````

Esto es porque la expresión ``:=`` se usa para asignar valor **localmente**. En este caso necesitamos
asignárselo **a la variable ya creada**.

El resto permanece igual.

Para asegurarnos, lancemos ``go run main.go``. 

Si ocurre un error parecido a:

```
panic: runtime error: invalid memory address or nil pointer dereference
```

Es, posiblemente, porque la variable `App`, que estamos usando como `puntero`, no tiene el valor asignado correctamente.


Ahora, vamos a crearnos un **directorio** llamado ``routes`` y, déntro de éste, ``routes.go``.

Dentro, añadimos el siguiente script:

````
package routes

import (
	"github.com/gofiber/fiber/v2"
	"hello-world/backend/config"
)

func Home() {
	config.App.Get("/", func(c *fiber.Ctx) error {
		return c.Render("index", fiber.Map{})
	})
}

````

```
Nota

Recuerda que **cada vez que creamos un directorio, el package recibe el mismo nombre que el directorio en el cual está contenido**
```

Importamos la variable utilizando la instrucción ``"hello-world/backend/config"``, y gracias a que tenemos la variable `App` exportada,
podemos utilizar la función ``Get`` que nos facilita este objeto.


Ahora, si comparamos ambos códigos:

**Fichero ``routes.go``**
```
    
    func Home() {
	config.App.Get("/", func(c *fiber.Ctx) error {
		return c.Render("index", fiber.Map{})
	})
}

```
+**Fichero ``config.go``**

```
func InitConfig() {

	App.Static("/", "./views")

	App.Get("/", func(c *fiber.Ctx) error {
		// Render index template
		return c.Render("index", fiber.Map{})
	})

}
```

Vemos que **En ambos ficheros se define la ruta** ``/``. No solo eso, sino también
**la misma entrada ``GET``**. Esto provoca que las configuraciones **se pisen**, y al estar previamente
definido en ``config.go``, lo estipulado en ``routes.go``es ignorado.

Dado que vamos a servir el html, a partir de ahora, desde el fichero ``routes.go``**, ya no necesitamos
**estos dos bloques de código** de ``config.go``.

Podemos eliminarlos y el resultado de ``config.go`` quedaría así:

````
package config

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
)

var App *fiber.App

func InitConfig() {

	engine := html.New("./views", ".html")

	App = fiber.New(fiber.Config{
		Views: engine,
		// Default global path to search for views (can be overriden when calling Render())
	})

}
````

Ahora, si lanzamos ``go run main.go`` deberíamos tener el mismo resultado que al principio, pero utilizando nuestra nueva configuración
de rutas.


### 4. Métodos de Cliente.

### 4.1 Nuestro primer método GET - HttpGetCaption

Ya hemos comprobado que podemos servir ficheros html desde nuestro back de ``Go``. Ahora, vamos a 
recoger **algo de data** que pasar a nuestro ``html`` para cargarlo y que no quede tan vacía la página.

Vamos a utilizar la `api` de esta página: https://imgflip.com/api para recoger imágenes de memes. Para ello,
usaremos la siguiente url: ``https://api.imgflip.com/get_memes``

``Fiber`` nos provee de los métodos que necesitamos para hacer las diferentes peticiones.

Vamos a crearnos un fichero nuevo para gestionar estas peticiones. Creemos el directorio ``data`` y, dentro, el directorio ``get_caption``. Después,
dentro de éste, creamos el fichero ``get_caption.go``.

Ésta debería ser la estructura:

 ```
data/
└── get_caption/        (carpeta)
    └── get_caption.go  (fichero) 
```

Ahora, dentro del fichero ``get_caption.go``, creamos esta función:

```
package get_caption

import (
	"github.com/gofiber/fiber/v2"
)

func HttpGetCaption() {
	response := fiber.Get("https://api.imgflip.com/get_memes")
}

```

Si acudimos a la documentación oficial de ``Fiber``: `https://docs.gofiber.io/api/client#start-request`

Vemos lo siguiente:

````
// Client http methods
func (c *Client) Get(url string) *Agent
func (c *Client) Head(url string) *Agent
func (c *Client) Post(url string) *Agent
func (c *Client) Put(url string) *Agent
func (c *Client) Patch(url string) *Agent
func (c *Client) Delete(url string) *Agent
````

Es decir, que debemos crear una ``función`` "padre" (por llamarla así) que reciba
un parámetro de tipo ``Client``, y dentro de ésta, usar el método ``Get``, que recibirá
la url que queremos utilizar para la petición, y a cambio nos devolverá un objeto de tipo ``Agent``.
En ese sentido, se comporta igual que nuestro método ``Home``.

```
Nota:

De manera nativa, `Go` nos provee del package `net/http` para realizar esos métodos, pero recordemos que estamos
utilizando Fiber, y por eso vamos a hacer uso de _sus_ métodos.
```



