/* @refresh reload */
import { render } from "solid-js/web";

import "./index.css";
import { Router } from "@solidjs/router";
import MyRoute from "./route";
import { HopeThemeConfig, HopeProvider } from "@hope-ui/solid";
import { globalCss } from "@hope-ui/solid";

const config: HopeThemeConfig = {
  initialColorMode: "dark",
  darkTheme: {
    colors: {
      background: "#282c34",
      primary1: "#1d131d",
    },
  },
};

render(
  () => (
    <HopeProvider config={config}>
      <Router>
        <MyRoute />
      </Router>
    </HopeProvider>
  ),
  document.getElementById("root") as HTMLElement
);
