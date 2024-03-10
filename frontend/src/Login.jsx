
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

    useEffect(() => {
        const params = new URLSearchParams(window.location.search);
        const authorized = params.get("authorized");
        if (authorized === "false") {
            setShowErrorModal(true); 
            setErrorMessage("You don't have permission to log in with this account.");
        }
    }, []);

    const handleGoogleLogin = async () => {
        try {
            const response = await fetch(`${process.env.REACT_APP_BACKEND_SERVER}/GoogleLogin`, {
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
                // Continue with Google OAuth URL
                const data = await response.json();
                if (data.url) {
                    window.location.href = data.url;
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
                <button className="google-button" onClick={handleGoogleLogin}> 
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
