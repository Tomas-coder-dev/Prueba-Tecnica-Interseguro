# Desafío Técnico TI - Interseguro 

Solución integral basada en arquitectura de microservicios para el procesamiento, rotación y cálculo estadístico de matrices. 

##  Enlaces de Producción (Despliegue en Render)

El proyecto está subido a la nube y es completamente funcional. Puedes probar la interfaz interactiva o consumir las APIs directamente:

*   **Frontend (UI Interactiva):** [https://frontend-interseguro.onrender.com]
*   **API Go (Core & QR):** [https://go-api-interseguro-rhs2.onrender.com]
*   **API Node.js (Estadísticas):** [https://node-api-interseguro-l8ya.onrender.com]


## 🏗️ Arquitectura del Sistema

La solución está dividida en tres contenedores independientes que se comunican entre sí:

1.  **API Orquestadora (Go / Fiber):**
    *   Recibe una matriz rectangular (n x m).
    *   Realiza una rotación de **90 grados en sentido horario**.
    *   Calcula la **factorización QR** utilizando la librería matemática `gonum`.
    *   Delega el cálculo estadístico enviando las matrices resultantes al microservicio de Node.js.
2.  **API Estadística (Node.js / Express):**
    *   Procesa las matrices recibidas desde Go.
    *   Calcula y retorna: **Máximo, Mínimo, Promedio, Suma Total** y valida si alguna es una **Matriz Diagonal**.
3.  **Frontend Cliente (Vue.js 3 / Vite) - *Plus Opcional*:**
    *   Interfaz web amigable para ingresar el JSON de la matriz y visualizar tanto la rotación como los datos estadísticos en tiempo real.


##  Ejecución Local con Docker

El proyecto está completamente contenedorizado para garantizar que funcione en cualquier entorno sin necesidad de instalar Go o Node.js de forma nativa.

### Requisitos previos
*   [Docker](https://www.docker.com/) y Docker Compose instalados.

### Pasos para levantar el proyecto

1. Clona este repositorio:

   git clone [https://github.com/Tomas-coder-dev/Prueba-Tecnica-Interseguro.git]
   cd Prueba-Tecnica-Interseguro