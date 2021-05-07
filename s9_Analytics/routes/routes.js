const app = (module.exports = require("express")());

async function authorizeRequest(req, res, next) {
  try {
    console.log("==========> Verifying token");
    const axios = require("axios");
    let authToken = req.get("Authorization");
    console.log(
      "===> Auth token header",
      authToken,
      " TYPe: ",
      typeof authToken
    );
    let authTknArray = authToken.split("Bearer ");

    console.log(
      "===> Auth tokenbody : ",
      authToken,
      " TYPe: ",
      typeof authToken
    );
    if (authToken == undefined || authToken == "" || authTknArray.length != 2) {
      throw "Invalid token";
    }

    authToken = authTknArray[1];

    const formUrlEncoded = (x) =>
      Object.keys(x).reduce(
        (p, c) => p + `&${c}=${encodeURIComponent(x[c])}`,
        ""
      );

    let response = await axios({
      url: "http://auth_service:8080/a/verify",
      method: "POST",
      headers: { "Content-Type": "application/x-www-form-urlencoded" },
      data: formUrlEncoded({
        token: authToken,
      }),
    });
    console.log("Response", response.data);
    if (response.status < 199 || response.status > 300) {
      throw "Invalid token";
      return;
    }
    req.id = response.data.userId;
    req.role = response.data.userType;
    next();
  } catch (err) {
    res.status(401).send("Invalid Token");
  }
}

app.get(
  "/ak/token",
  async (req, res, next) => {
    authorizeRequest(req, res, next);
  },
  async (req, res) => {
    try {
      console.log(
        "======> New token request from  ID: ",
        req.id,
        " , Role: ",
        req.role
      );
      let adal = require("adal-node");
      const tenant = "71d91dcf-7c64-41ea-9c4f-d3183cfe4a96";
      const resource = "https://analysis.windows.net/powerbi/api";
      var context = new adal.AuthenticationContext(
        "https://login.microsoftonline.com/common/oauth2/nativeclient"
      );
      context.acquireTokenWithUsernamePassword(
        resource,
        "sricharanpisupati@pgktechnologies.com",
        "Lunarbard929",
        "cdab76df-3b4d-49ad-8b94-114d88aa93ea",
        (err, tokenResponse) => {
          if (err) {
            console.log(`Token generation failed due to ${err}`);
            res.status(400).send({ message: `${err}` });
          } else {
            console.dir(tokenResponse, { depth: null, colors: true });
            res.status(200).send({
              accessToken: tokenResponse.accessToken,
              refreshToken: tokenResponse.refreshToken,
            });
          }
        }
      );
    } catch (error) {
      res.status(500).send(error);
    }
  }
);
