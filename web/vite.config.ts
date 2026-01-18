import { defineConfig } from "vite";
import react from "@vitejs/plugin-react";
import tailwindcss from "@tailwindcss/vite";

// https://vite.dev/config/
export default defineConfig({
  plugins: [
    react({
      babel: { plugins: [["babel-plugin-react-compiler"]] },
    }),
    tailwindcss(),
  ],
  server: {
    proxy: {
      "/listing": "http://1.1.0.10:3000",
      "/static": "http://1.1.0.10:3000",
    },
  },
});
