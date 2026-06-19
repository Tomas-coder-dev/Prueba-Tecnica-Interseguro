const express = require("express");

const statisticsRoutes = require("./routes/statistics.routes");

const app = express();

app.use(express.json());

app.use("/statistics", statisticsRoutes);

app.listen(3000, () => {
  console.log("Node API running on port 3000");
});