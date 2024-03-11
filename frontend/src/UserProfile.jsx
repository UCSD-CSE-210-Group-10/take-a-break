import React from 'react';
import './UserProfile.css';
import NavigationBar from './NavigationBar';

const UserProfile = ({ user, handleLogout }) => {
    // Dummy data for testing
    const dummyUser = {
        name: 'John Doe',
        email: 'john.doe@example.com',
        imageUrl: './UCSD-logo.png', // Replace with an actual image URL
      };
      
  
    // Use the provided user prop if available, otherwise use dummy data for testing
    const userData = user || dummyUser;
  
    return (
        <div>
            <NavigationBar handleLogout={handleLogout}/>
            <div className="user-profile">
                <h2>User Profile</h2>
                <div className="profile-info">
                    <img src={userData.imageUrl} alt="User Avatar" className="avatar" />
                    <div className="details">
                        <p>Name: {userData.name}</p>
                        <p>Email: {userData.email}</p>
                        {/* Add more profile information if needed */}
                    </div>
                </div>
            </div>
        </div>

    );
  };
  
  export default UserProfile;