import React from "react";
import logo from './UCSD-logo.png'
import clipart from './Take-a-break-clipart.png'
import googleLogo from './Google-logo.png'
import './Login.css';

const Login = () => {

    return(
        <div className="container">
            <div className="left-section">
                <div className="logo-container">
                    <img src={logo} className="UCSD-logo" alt="logo" />
                </div>
                <div className="App-name-header">
                    <h1>
                        Take a Break
                    </h1>
                </div>
                <div className="clipart-container">
                    <img src={clipart} className="Login-Clipart" alt="logo" />
                </div>
                
            </div>

            <div className="right-section">
                <h2 className="Sign-in-header">
                        Sign In
                </h2>
                <p>Sign in with your UCSD Credentials</p>
                <button className="google-button"> 
                    <img src={googleLogo} className="google-logo" alt="google-logo" /> Continue with Google
                </button>
            </div>

        </div>
    );
};

export default Login;