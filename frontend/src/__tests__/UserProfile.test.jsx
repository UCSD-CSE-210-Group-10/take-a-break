// UserProfile.test.js
import React from 'react';
import { render } from '@testing-library/react';
import UserProfile from './UserProfile';

test('renders user profile with dummy data', () => {
  const dummyUser = {
    name: 'John Doe',
    email: 'john.doe@example.com',
    imageUrl: './UCSD-logo.png',
  };

  const { getByText, getByAltText } = render(<UserProfile user={dummyUser} />);

  // Check if the user information is rendered
  const nameElement = getByText(/John Doe/i);
  const emailElement = getByText(/john\.doe@example\.com/i);
  const avatarElement = getByAltText('User Avatar');

  expect(nameElement).toBeInTheDocument();
  expect(emailElement).toBeInTheDocument();
  expect(avatarElement).toBeInTheDocument();
});

