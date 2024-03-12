import React, { useState, useEffect } from "react";
import './UserProfile.css';
import NavigationBar from './NavigationBar';

const UserProfile = ({ handleLogout }) => {
    // Dummy data for testing

    const [user, setUser] = useState([]);

    useEffect(() => {
        const jwtToken = localStorage.getItem('token');
        const fetchUser= async () => {
          try {
            const { hostname, protocol } = window.location;
            const response = await fetch(`${protocol}//${hostname}:8080/users/${jwtToken}`);
            const data = await response.json();
            console.log(data)
            if(data.error && data.error === "Auth Error") {
              handleLogout()
            }
            setUser(data);
          } catch (error) {
            console.error('Error fetching events:', error);
          }
        };
    
        fetchUser();
      }, [handleLogout]);
  
    return (
        <div>
            <NavigationBar handleLogout={handleLogout}/>
            <div className="user-profile">
                <h2>User Profile</h2>
                <div className="profile-info">
                    <img src={user.avatar} alt="User Avatar" className="avatar" />
                    <div className="details">
                        <p>Name: {user.name}</p>
                        <p>Email: {user.email_id}</p>
                        {/* Add more profile information if needed */}
                    </div>
                </div>
            </div>
        </div>

    );
  };
  
  export default UserProfile;