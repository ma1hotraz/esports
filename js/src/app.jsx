import axios from "axios";
import React from "react";
import { createRoot } from "react-dom/client";
import { Toaster } from "react-hot-toast";
import { Provider } from "react-redux";
import { RouterProvider } from "react-router-dom";
import { BASE_URL } from "./constants";
import Layout from "./layout";
import Router from "./router";
import store from "./store";

axios.defaults.baseURL = BASE_URL;
axios.defaults.withCredentials = true;

export default function App() {
    return (
        <Provider store={store}>
            <RouterProvider router={Router}>
                <Layout />
            </RouterProvider>
            <Toaster />
        </Provider>
    );
}
const domNode = document.getElementById("root");
const root = createRoot(domNode);
root.render(<App />);
