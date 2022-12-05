import { Route, Routes } from "@solidjs/router";
import App from "./App";
import Login from "./pages/Login";
import Register from "./pages/Register";
import TodoPage from "./pages/TodoPage";

const MyRoute = () => {
  return (
    <Routes>
      <Route path="/" component={TodoPage} />
      <Route path="/register" component={Register} />
      <Route path="/login" component={Login} />
    </Routes>
  );
};

export default MyRoute;
