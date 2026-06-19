const express = require("express");
const router = express.Router();

const { calculateStatistics } = require("../controllers/statistics.controller");

// Único endpoint (POST /statistics) para procesar los datos enviados desde Go
router.post("/", calculateStatistics);

module.exports = router;