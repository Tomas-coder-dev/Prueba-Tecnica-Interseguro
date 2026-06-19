# Desafío Técnico TI - Interseguro 

Solución integral basada en arquitectura de microservicios para el procesamiento, rotación y cálculo estadístico de matrices. 

##  Enlaces de Producción (Despliegue en Render)

El proyecto está subido a la nube y es completamente funcional. Puedes probar la interfaz interactiva o consumir las APIs directamente:

*   **Frontend (UI Interactiva):** [https://frontend-interseguro.onrender.com]
*   **API Go (Core & QR):** [https://go-api-interseguro-rhs2.onrender.com]
*   **API Node.js (Estadísticas):** [https://node-api-interseguro-l8ya.onrender.com]

##  Arquitectura del Sistema

La solución está dividida en contenedores independientes que se comunican mediante HTTP:

1.  **API Orquestadora (Go / Fiber):** Recibe la matriz, realiza la rotación de 90° horario y calcula la factorización QR. Delega el cálculo estadístico al microservicio de Node.js.
2.  **API Estadística (Node.js / Express):** Procesa las matrices recibidas y calcula: Máximo, Mínimo, Promedio, Suma Total y validación de matriz diagonal.
3.  **Frontend Cliente (Vue.js 3 / Vite):** Interfaz para gestión de matrices y visualización de resultados en tiempo real.

##  Pruebas Unitarias

El proyecto incluye suites de pruebas automatizadas para garantizar la precisión matemática y la integridad de los datos:

*   **Go (Lógica de Rotación y Validación):** Ejecuta `go test ./utils -v` en la carpeta `go-api`.
*   **Node.js (Lógica Estadística):** Ejecuta `npm test` en la carpeta `node-api` (utilizando **Jest**).


##  Ejecución Local con Docker

El proyecto está totalmente dockerizado

### Pasos para levantar el proyecto

1. Clona este repositorio:

   git clone [https://github.com/Tomas-coder-dev/Prueba-Tecnica-Interseguro.git]
   cd Prueba-Tecnica-Interseguro

