import { Route, Routes } from "@solidjs/router";
import App from "./App";
import Login from "./pages/Login";
import Register from "./pages/Register";

const MyRoute = () => {
  return (
    <Routes>
      <Route path="/" component={App} />
      <Route path="/register" component={Register} />
      <Route path="/login" component={Login} />
    </Routes>
  );
};

export default MyRoute;
