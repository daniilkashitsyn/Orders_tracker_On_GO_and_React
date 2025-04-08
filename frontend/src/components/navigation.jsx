import React from "react";
import {Link} from "react-router-dom";

function Navigation() {
    return (
        <nav className="navbar navbar-expand-md navbar-dark text-light bg-dark sticky-top">
            <div className="container-fluid">
                <Link to="/" className="navbar-brand text-warning font-monospace">ГЕРКУЛЕС</Link>
                <div className="collapse navbar-collapse">
                    <ul className="navbar-nav m-lg-auto">
                        <li className="nav-item"><Link to="/clients" className="nav-link">Клиенты</Link></li>
                        <li className="nav-item"><Link to="/clients/add" className="nav-link">Добавить</Link></li>
                        <li className="nav-item"><Link to="/"></Link></li>
                    </ul>
                </div>
            </div>
        </nav>
    );
}

export default Navigation