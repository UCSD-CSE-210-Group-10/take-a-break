
import React, { useState, useEffect } from "react";
import logo from './UCSD-logo.png'
import clipart from './Take-a-break-clipart.png'
import googleLogo from './Google-logo.png'
import { Modal, Button } from 'react-bootstrap';
import 'bootstrap/dist/css/bootstrap.min.css';
import './Login.css';

const Login = () => {
    
    const [showErrorModal, setShowErrorModal] = useState(false);
    const [errorMessage, setErrorMessage] = useState("");

    const config = {
        clientID: process.env.REACT_APP_GOOGLE_CLIENT_ID,
        authURL: process.env.REACT_APP_AUTHURL,
        tokenURL: process.env.REACT_APP_TOKENURL,
        redirectURL: process.env.REACT_APP_REDIRECT_URL,
        clientURL: process.env.REACT_APP_CLIENT_URL,
        tokenExpiration: 36000,
        postURL: process.env.REACT_APP_POSTURL,
      };

    const authParams = () => {
        const params = new URLSearchParams();
        params.set('client_id', config.clientID);
        params.set('redirect_uri', config.redirectURL);
        params.set('response_type', 'code');
        params.set('scope', 'openid profile email');
        params.set('access_type', 'offline');
        params.set('state', 'standard_oauth');
        params.set('prompt', 'consent');
        return params.toString();
    };
    
    const getAuthURL = () => {
        const authURL = `${config.authURL}?${authParams()}`;
        window.location.href = authURL;
    };


    useEffect(() => {
        
        const params = new URLSearchParams(window.location.search);
        if (params.has("code")) {
            const code = params.get("code");
            handleGoogleLogin(code)
        }

    }, []);

    const handleGoogleLogin = async (code) => {
        try {
            const response = await fetch(`${process.env.REACT_APP_SERVER_URL}/auth/token?code=${code}`, {
                method: "GET",
            });

            if (!response.ok) {
                const data = await response.json();
                if (data.error && data.redirect) {
                    setErrorMessage(data.error);
                    setShowErrorModal(true);
                    if (data.redirect) {
                        window.location.href = data.redirect;
                    }
                } else {
                    throw new Error(`HTTP error! Status: ${response.status}`);
                }
            } else {

                const data = await response.json();
                const authorized = data.authorized;
                if (authorized === false) {
                    setShowErrorModal(true); 
                    setErrorMessage("You don't have permission to log in with this account.");
                }
                else {
                    console.log(data.token);
                    localStorage.setItem("token", data.token);
                    window.location.href = `${process.env.REACT_APP_CLIENT_URL}/events`;
                }
            }
        } catch (error) {
            console.error("Error:", error);
        }
    };
    

    return(
        <div className="login-screen">
            <div className="left-section" data-testid="left-section">
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

            <div className="right-section" data-testid="right-section">
                <h2 className="Sign-in-header">
                        Sign In
                </h2>
                <p>Sign in with your UCSD Credentials</p>
                <button className="google-button" onClick={getAuthURL}> 
                    <img src={googleLogo} className="google-logo" alt="google-logo" /> Continue with Google
                </button>
            </div>

            <Modal show={showErrorModal} onHide={() => setShowErrorModal(false)}>
                <Modal.Header closeButton>
                    <Modal.Title>Error</Modal.Title>
                </Modal.Header>
                <Modal.Body>{errorMessage}</Modal.Body>
                <Modal.Footer>
                    <Button variant="secondary" onClick={() => setShowErrorModal(false)}>
                        Close
                    </Button>
                </Modal.Footer>
            </Modal>
        </div>
    );
};

export default Login;
