// Aplanamos Q y R en un solo array para que sacar el max/min sea más directo
const flattenMatrices = (q, r) => {
  return [...q.flat(), ...r.flat()];
};

// Valida si la matriz es diagonal.
//  Usamos una tolerancia (1e-10) en vez de comparar con 0 absoluto 
// porque los cálculos en Go (float64) suelen dejar decimales residuales.
const isDiagonal = (matrix) => {
  const rows = matrix.length;
  const cols = matrix[0].length;

  if (rows !== cols) {
    return false;
  }

  for (let i = 0; i < rows; i++) {
    for (let j = 0; j < cols; j++) {
      if (i !== j && Math.abs(matrix[i][j]) > 1e-10) {
        return false;
      }
    }
  }

  return true;
};

// Calcula las métricas principales procesando ambas matrices juntas
const getStatistics = (data) => {
  const { q, r } = data;

  const values = flattenMatrices(q, r);

  const sum = values.reduce((acc, value) => acc + value, 0);

  // Prevenimos un NaN si por algún motivo las matrices llegan vacías
  const average = values.length > 0
    ? sum / values.length
    : 0;

  const max = Math.max(...values);
  const min = Math.min(...values);

  return {
    max,
    min,
    average,
    sum,
    // Verificamos si al menos una de las dos cumple
    hasDiagonalMatrix: isDiagonal(q) || isDiagonal(r)
  };
};

module.exports = {
  getStatistics
};