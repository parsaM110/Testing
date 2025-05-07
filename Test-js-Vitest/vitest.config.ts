import AllureReporter from "allure-vitest/reporter";
import { defineConfig } from "vitest/config";

export default defineConfig({
  test: {
    setupFiles: ["allure-vitest/setup"],
    reporters: [
      "verbose",
      [
        "allure-vitest/reporter",
        {
          resultsDir: "allure-results",
        },
      ],
    ],
  },
});