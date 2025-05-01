# Java Coprocessor

This service is to serve as an example for how to store request variables from the router into the request context, so that those variables can be referenced/used elsewhere. The main logic is in the `RouterRequestHandler.java` file with the entry point being in `CoprocessorController.java`, and some helpful models in the `models` folder.

## Run

  1. _(Optional)_ Install [SDK Man](https://sdkman.io/)
      - This can help with installing and managing different Java versions
  1. Start the coprocessor using `mvn`
      ```shell
      ./mvnw spring-boot:run
      ```
  2. Start the router
