// UserProfile.test.js
import React from 'react';
import { render } from '@testing-library/react';
import { BrowserRouter } from 'react-router-dom';  // Import BrowserRouter
import UserProfile from '../UserProfile';

test('renders user profile with dummy data', () => {
  const dummyUser = {
    name: 'John Doe',
    email: 'john.doe@example.com',
    imageUrl: './UCSD-logo.png',
  };

  const { getByText, getByAltText } = render(
    <BrowserRouter>  {/* Wrap your component rendering with BrowserRouter */}
      <UserProfile user={dummyUser} />
    </BrowserRouter>,
  );

  // Check if the user information is rendered
  const nameElement = getByTestId('user-name');
  const emailElement = getByTestId('user-email');
  const avatarElement = getByAltText('User Avatar');

  expect(nameElement).toBeInTheDocument();
  expect(emailElement).toBeInTheDocument();
  expect(avatarElement).toBeInTheDocument();
});
