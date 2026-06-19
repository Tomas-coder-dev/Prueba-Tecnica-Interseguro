const { getStatistics } = require('../services/statistics.service');

describe('Statistics Service Logic', () => {
  test('debería calcular correctamente max, min, avg y sum', () => {
    const data = {
      q: [[1, 0], [0, 1]],
      r: [[2, 2], [0, 2]]
    };
    
    const stats = getStatistics(data);
    
    expect(stats.max).toBe(2);
    expect(stats.min).toBe(0);
    expect(stats.sum).toBe(8);
    expect(stats.average).toBe(1); // (1+0+0+1+2+2+0+2) / 8 = 8/8 = 1
  });

  test('debería detectar si una matriz es diagonal', () => {
    const data = {
      q: [[5, 0], [0, 5]], // Es diagonal
      r: [[1, 2], [3, 4]]  // No es diagonal
    };
    
    const stats = getStatistics(data);
    expect(stats.hasDiagonalMatrix).toBe(true);
  });
});