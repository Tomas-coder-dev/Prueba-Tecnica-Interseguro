const { getStatistics } = require("../services/statistics.service");

// calculateStatistics recibe las matrices Q y R enviadas por la API en Go,
// delega los cálculos al servicio y retorna el resultado en formato JSON.
const calculateStatistics = (req, res) => {

  const result = getStatistics(req.body);

  res.json(result);
};

module.exports = {
  calculateStatistics
};