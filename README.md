#  Desafío Técnico  - Procesador de Matrices

Este proyecto es la solución a la prueba técnica para el procesamiento de matrices. En lugar de hacer un script simple, opté por diseñar una arquitectura moderna y escalable basada en microservicios, utilizando contenedores para garantizar que el código se ejecute de manera idéntica en cualquier entorno.

##  La Arquitectura

El sistema está compuesto por tres piezas clave que se comunican de forma aislada a través de una red interna:

1. **El Motor Matemático (Go):** Construido sobre **Fiber** por su altísimo rendimiento. Este servicio actúa como la puerta de entrada, recibe la matriz original, calcula su factorización QR y delega el resto del trabajo.
2. **El Procesador Estadístico (Node.js):** Construido con **Express**. Se encarga de recibir los resultados matemáticos de Go y procesar la lógica de negocio (cálculo de máximos, mínimos, promedios, sumas y la validación de matriz diagonal).
3. **El Panel de Control (Vue.js + Nginx):** *[Implementación Opcional]* Para no depender de Postman o cURL al evaluar la prueba, construí una interfaz minimalista. Se compila estáticamente y es servida por **Nginx**, lo que la hace extremadamente rápida.

## 🛠️ Tecnologías Elegidas

* **Backend:** Go (Golang) 1.26, Node.js 20
* **Frontend:** Vue 3 (Composition API), Vite, CSS puro (sin dependencias pesadas).
* **Infraestructura:** Docker, Docker Compose.

---

## 🚦 Cómo probar el proyecto localmente

Gracias a Docker, no necesitas instalar Go, Node, ni configurar variables de entorno en tu máquina. Si tienes Docker instalado, estás a un comando de distancia.

### Pasos:

1. Clona este repositorio en tu máquina:
   ```bash
   git clone [https://github.com/TU_USUARIO/technical-challenge-ti.git](https://github.com/TU_USUARIO/technical-challenge-ti.git)
   cd technical-challenge-ti
