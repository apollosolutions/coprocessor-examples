import express from "express";

const app = express();

const processRouterRequestStage = async (payload) => {
  // See: https://www.apollographql.com/docs/router/customizations/coprocessor/ for more information around how you can leverage coprocessors

  // At the router request stage the body is a string, so you'll need to convert it to a JSON object
  const body = JSON.parse(payload.body)

  // Adding the variables as stringified JSON to context allows you to pass them to any other part of the request
  // For example you could pass this context variable to an OTel span to log out what is being passed by the client
  // I do want to note though that when passing variable values as in this sample, to exercise extreme caution as
  // variables could contain PII (personally identifiable information)
  payload.context.entries["variableValues"] = JSON.stringify(body.variables);

  /*
    In this example, we're setting our context variable to hold our request variables in the Router Request stage.
    This means that if we want it to show up in our trace, it needs to be emitted after the Router Request span since
    the span is emitted before the coprocessor logic is run. Here's an example setting it at the Supergraph Request stage in the router config:

    telemetry
      instrumentation:
        spans:
          mode: spec_compliant
          supergraph:
            attributes:
              graphql.document: true
              variable.values:
                request_context: "variableValues"
  */

  return payload;
};

app.post("/", express.json(), async (req, res) => {
  const payload = req.body;

  let response = payload;
  switch (payload.stage) {
    case "RouterRequest":
      response = await processRouterRequestStage(payload);
      break;
  }

  res.send(response);
});

app.listen(3007, () => {
  console.log("🚀 Server running at http://localhost:3007");
});
