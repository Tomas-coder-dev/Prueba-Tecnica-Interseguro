const express = require("express");

const statisticsRoutes = require("./routes/statistics.routes");

const app = express();

app.use(express.json());

app.use("/statistics", statisticsRoutes);

// Usamos el puerto que Render nos da, o el 3000 si estamos en local
const PORT = process.env.PORT || 3000;

app.listen(PORT, () => {
  console.log(`Node API running on port ${PORT}`);
});