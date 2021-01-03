import React from 'react';
import {withRouter} from "react-router-dom";
import {ACCESS_TOKEN_NAME} from '../../constants/apiConstants';

function Header(props) {
    const capitalize = (s) => {
        if (typeof s !== 'string') return ''
        return s.charAt(0).toUpperCase() + s.slice(1)
    }
    let title = capitalize(props.location.pathname.substring(1, props.location.pathname.length))
    if (props.location.pathname === '/') {
        title = 'Home'
    }

    function renderLogout() {
        const token = localStorage.getItem(ACCESS_TOKEN_NAME);
        if (props.location.pathname === '/' && token) {
            return (
                <div className="ml-auto">
                    <button className="btn btn-danger" onClick={() => handleLogout()}>Logout</button>
                </div>
            )
        }
    }

    function handleLogout() {
        localStorage.removeItem(ACCESS_TOKEN_NAME)
        props.history.push('/login')
    }

    function renderLogin() {
        // dummy login check by looking the local storage
        const token = localStorage.getItem(ACCESS_TOKEN_NAME);
        if (props.location.pathname === '/' && !token) {
            return (
                <div className="ml-auto">
                    <button className="btn btn-primary" onClick={() => handleLogin()}>Login</button>
                </div>
            )
        }
    }

    function handleLogin() {
        props.history.push('/login')
    }

    function renderHome() {
        if (props.location.pathname === '/login' || props.location.pathname === '/register') {
            return (
                <div className="ml-auto">
                    <button className="btn btn-primary" onClick={() => handleHome()}>Home</button>
                </div>
            )
        }
    }

    function handleHome() {
        props.history.push('/')
        title = 'Home'
    }

    return (
        <nav className="navbar navbar-dark bg-primary">
            <div className="row col-12 d-flex justify-content-center text-white">
                <span className="h3">{title || props.title}</span>
                {renderHome()}
                {renderLogout()}
                {renderLogin()}
            </div>
        </nav>
    )
}

export default withRouter(Header);
