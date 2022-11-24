import { Route, Routes } from "@solidjs/router";
import App from "./App";
import Register from "./pages/Register";

const MyRoute = () => {
  return (
    <Routes>
      <Route path="/" component={App} />
      <Route path="/register" component={Register} />
    </Routes>
  );
};

export default MyRoute;
