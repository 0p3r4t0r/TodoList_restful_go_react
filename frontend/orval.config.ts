import { defineConfig } from "orval";

export default defineConfig({
  api: {
    input: {
      target: "./specs/swagger.json",
    },
    output: {
      target: "./src/api.ts",
      client: "react-query",
    },
    // hooks: {
    //   afterAllFilesWrite: "prettier --write",
    // },
  },
});
