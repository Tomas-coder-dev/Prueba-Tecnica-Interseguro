<script setup>
import { ref } from 'vue';
import axios from 'axios';

// data de prueba por defecto
const inputMatrix = ref('[\n  [1, 2, 3],\n  [4, 5, 6],\n  [7, 8, 9]\n]');
const results = ref(null);
const errorMsg = ref('');
const isLoading = ref(false);

// cortamos los decimales largos para no romper el diseño de las tarjetas
const formatNumber = (num) => {
  if (typeof num !== 'number') return num;
  return Number.isInteger(num) ? num : Number(num.toFixed(4));
};

const processMatrix = async () => {
  errorMsg.value = '';
  results.value = null;
  isLoading.value = true;

  try {
    const parsedMatrix = JSON.parse(inputMatrix.value);

    // le pegamos al backend de go pasándole la matriz
    const response = await axios.post(`${import.meta.env.VITE_API_URL}/qr`, {
      matrix: parsedMatrix
    });

    // guardamos todo el payload (ahora trae rotatedMatrix y statistics)
    results.value = response.data;

  } catch (err) {
    if (err instanceof SyntaxError) {
      errorMsg.value = 'Revisa el formato del JSON. Faltan comas o corchetes por ahí.';
    } else {
      errorMsg.value = err.response?.data?.error || 'Se cayó la conexión con el backend de Go.';
    }
  } finally {
    isLoading.value = false;
  }
};
</script>

<template>
  <main class="container">
    <header>
      <h1>Desafío Técnico TI</h1>
      <p>Procesador de Matrices (Go + Node.js)</p>
    </header>
    
    <section class="panel input-panel">
      <label for="json-input">Ingresa la matriz a procesar:</label>
      <textarea id="json-input" v-model="inputMatrix" rows="8"></textarea>
      
      <button @click="processMatrix" :disabled="isLoading">
        {{ isLoading ? 'Procesando...' : 'Calcular Estadísticas' }}
      </button>
      
      <div v-if="errorMsg" class="alert error">
        {{ errorMsg }}
      </div>
    </section>

    <section v-if="results && results.rotatedMatrix" class="panel">
      <h2>Matriz Rotada (90° Horario)</h2>
      <pre class="matrix-preview">{{ JSON.stringify(results.rotatedMatrix, null, 2) }}</pre>
    </section>

    <section v-if="results && results.statistics" class="panel results-panel">
      <h2>Resultados Estadísticos</h2>
      <div class="grid">
        <div class="card">
          <span class="label">Máximo</span>
          <span class="value">{{ formatNumber(results.statistics.max) }}</span>
        </div>
        <div class="card">
          <span class="label">Mínimo</span>
          <span class="value">{{ formatNumber(results.statistics.min) }}</span>
        </div>
        <div class="card">
          <span class="label">Promedio</span>
          <span class="value">{{ formatNumber(results.statistics.average) }}</span>
        </div>
        <div class="card">
          <span class="label">Suma</span>
          <span class="value">{{ formatNumber(results.statistics.sum) }}</span>
        </div>
        <div class="card">
          <span class="label">¿Es Diagonal?</span>
          <span class="value">{{ results.statistics.hasDiagonalMatrix ? 'Sí' : 'No' }}</span>
        </div>
      </div>
    </section>
  </main>
</template>

<style scoped>
.container {
  max-width: 800px;
  margin: 2rem auto;
  padding: 0 1rem;
}

header {
  text-align: center;
  margin-bottom: 2rem;
}

header h1 {
  font-size: 2rem;
  color: #1e293b;
}

header p {
  color: #64748b;
}

.panel {
  background: white;
  padding: 1.5rem;
  border-radius: 8px;
  box-shadow: 0 4px 6px -1px rgba(0, 0, 0, 0.1);
  margin-bottom: 1.5rem;
}

.input-panel {
  display: flex;
  flex-direction: column;
  gap: 1rem;
}

textarea {
  font-family: monospace;
  padding: 1rem;
  border: 1px solid #cbd5e1;
  border-radius: 4px;
  resize: vertical;
}

button {
  padding: 0.75rem;
  background-color: #3b82f6;
  color: white;
  border: none;
  border-radius: 4px;
  font-weight: 600;
  cursor: pointer;
  transition: opacity 0.2s;
}

button:hover:not(:disabled) {
  opacity: 0.9;
}

button:disabled {
  background-color: #94a3b8;
  cursor: not-allowed;
}

.alert.error {
  padding: 1rem;
  background-color: #fef2f2;
  color: #991b1b;
  border-left: 4px solid #ef4444;
  border-radius: 4px;
}

.matrix-preview {
  background: #f1f5f9;
  padding: 1rem;
  border-radius: 6px;
  overflow-x: auto;
  font-family: monospace;
  color: #334155;
  margin-top: 1rem;
}

.grid {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(100px, 1fr));
  gap: 1rem;
  margin-top: 1rem;
}

.card {
  background: #f1f5f9;
  padding: 1rem;
  border-radius: 6px;
  display: flex;
  flex-direction: column;
  align-items: center;
  gap: 0.5rem;
  word-break: break-word;
  text-align: center;
}

.card .label {
  font-size: 0.8rem;
  color: #64748b;
  text-transform: uppercase;
  font-weight: bold;
}

.card .value {
  font-size: 1.2rem;
  font-weight: 800;
  color: #0f172a;
}
</style>