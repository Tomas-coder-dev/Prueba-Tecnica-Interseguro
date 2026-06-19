#  Desafío Técnico TI | Arquitectura Distribuida para Procesamiento de Matrices

Este repositorio contiene mi solución integral al Coding Challenge para la posición de TI. 

En lugar de construir scripts aislados, decidí diseñar e implementar una **arquitectura moderna de microservicios contenerizados**. La solución no solo cumple con los requisitos matemáticos obligatorios, sino que incorpora un frontend completo y buenas prácticas de infraestructura y despliegue.

---

##  Topología del Sistema

El proyecto está compuesto por tres servicios independientes que operan dentro de una red privada de Docker:

1. **Go API (Gateway & Math Engine):** * Construida con **Fiber** por su baja latencia y alto rendimiento.
   * **Responsabilidad:** Actúa como punto de entrada (recibe el POST del usuario), calcula la factorización QR de la matriz ingresada y delega el análisis estadístico al siguiente microservicio.
2. **Node.js API (Data Processor):** * Construida con **Express.js**.
   * **Responsabilidad:** Recibe de Go los resultados de la rotación/factorización y procesa la lógica de negocio secundaria (Máximo, Mínimo, Promedio, Suma Total y validación de matriz diagonal).
3. **Frontend SPA (Bonus Implementado):** * Construido con **Vue 3** y servido estáticamente a través de **Nginx**.
   * **Responsabilidad:** Proporcionar una experiencia de usuario limpia y evitar el uso obligatorio de herramientas como Postman para evaluar la prueba. Incluye validación de sintaxis JSON en el cliente.

---

## 🛠️ Stack Tecnológico

* **Lenguajes:** Go (1.26), Node.js (v20), JavaScript/HTML/CSS
* **Frameworks:** Fiber, Express, Vue 3 (Composition API), Vite
* **Infraestructura:** Docker, Docker Compose, Nginx (Docker Hub: `aylas`)
* **Comunicaciones:** RESTful HTTP, JSON

---

##  Guía de Despliegue Local

Gracias a Docker, no es necesario instalar ni configurar entornos de desarrollo complejos. Si tienes Docker corriendo en tu máquina, estás a un comando de levantar toda la infraestructura.

### Requisitos previos
* [Docker Desktop](https://www.docker.com/products/docker-desktop/) instalado.
* Puertos `80`, `8080` y `3000` libres en tu entorno local.

### Pasos de ejecución

1. Clona este repositorio y entra a la carpeta del proyecto:
   ```bash
   git clone [https://github.com/Tomas-coder-dev/Prueba-Tecnica-Interseguro.git](https://github.com/Tomas-coder-dev/Prueba-Tecnica-Interseguro.git)
   cd Prueba-Tecnica-Interseguro
