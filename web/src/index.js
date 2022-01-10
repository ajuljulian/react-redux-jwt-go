import React from "react";
import ReactDOM from "react-dom";
import { BrowserRouter, Routes, Route } from "react-router-dom";
import { Provider } from "react-redux";
import "./index.css";
import App from "./App";
import store from "./store";
import reportWebVitals from "./reportWebVitals";
import "bootstrap/dist/css/bootstrap.min.css";

import Login from "./components/Login";
import Register from "./components/Register";
import Home from './components/Home';
import Profile from './components/Profile';
import NoPage from './components/NoPage';
import BoardUser from "./components/BoardUser";
import BoardModerator from "./components/BoardModerator";
import BoardAdmin from "./components/BoardAdmin";

ReactDOM.render(
  // Wrap the entire <App /> with a redux <Provider />
  // and pass the redux store to it as a prop
  <React.StrictMode>
    <BrowserRouter>
      <Provider store={store}>
      <Routes>
          <Route path="/" element={<App />} >
            <Route path="home" element={<Home />} />
            <Route path="profile" element={<Profile />} />
            <Route path="login" element={<Login />} />
            <Route path="register" element={<Register />} />
            <Route path="user" element={<BoardUser />} />
            <Route path="mod" element={<BoardModerator />} />
            <Route path="admin" element={<BoardAdmin />} />
            <Route path="*" element={<NoPage />} />
          </Route>
        </Routes>
      </Provider>
    </BrowserRouter>
  </React.StrictMode>,
  document.getElementById('root')
);

// If you want to start measuring performance in your app, pass a function
// to log results (for example: reportWebVitals(console.log))
// or send to an analytics endpoint. Learn more: https://bit.ly/CRA-vitals
reportWebVitals();
