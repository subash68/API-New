const helmet = require("helmet");
const app = require("./loaders/express");
const handleAsyncExceptions = require("./loaders/handleError");
const routes = require("./routes/routes");

app.use(helmet());
app.set("trust proxy", true);

/**Express server starts here */
async function run() {
  try {
    // mount the routes
    app.use(routes);

    // Server starts here
    app.listen(8080, "0.0.0.0", function (err) {
      if (err) {
        console.log("Failed to start the server " + err);
      }
      console.log("Asset tracking apis are running on http://0.0.0.0:8080");
    });
  } catch (err) {
    throw new Error(err);
  }
}

module.exports = run;

if (require.main === module) {
  handleAsyncExceptions();
  run();
}
