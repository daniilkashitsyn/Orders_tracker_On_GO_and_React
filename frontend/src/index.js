import React from "react";
import ReactDOM from "react-dom/client";
import { BrowserRouter as Router, Route, Routes } from "react-router-dom";
import "bootstrap/dist/css/bootstrap.css";
import Clients from "./clients/clients";
import Navigation from "./components/navigation";
import AddClient from "./clients/addClient";


const APP = () => {
    return (
        <Router>
            <Navigation></Navigation>
            <Routes>
                <Route path="/clients" element={<Clients />}></Route>
                <Route path="/clients/add" element={<AddClient />}></Route>
            </Routes>
        </Router>
    )
}

const root = ReactDOM.createRoot(document.getElementById('root'))

root.render(
    <React.StrictMode>
        <APP/>
    </React.StrictMode>
)