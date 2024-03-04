import { render, screen } from '@testing-library/react'
import NavigationBar from '../NavigationBar'

test("Navigation Bar Renders Successfully", () => {
    render(<NavigationBar/>);
    
    const leftSec = screen.getByTestId('navigation-bar');
    expect(leftSec).toBeInTheDocument();
})

test("Links Render Successfully", () => {
    render(<NavigationBar/>);
  
    const logoElement = screen.getByAltText('UCSD Logo');
    expect(logoElement).toBeInTheDocument();
  
    const eventsLink = screen.getByText('Events');
    expect(eventsLink).toBeInTheDocument();
  
    const healthLink = screen.getByText('Health');
    expect(healthLink).toBeInTheDocument();
  
    const friendsLink = screen.getByText('Friends');
    expect(friendsLink).toBeInTheDocument();
  
    const studentDropdown = screen.getByText('Student');
    expect(studentDropdown).toBeInTheDocument();
})

test("Logo Renders Successfully", () => {
    render(<NavigationBar/>);
  
    const logoElement = screen.getByAltText('UCSD Logo');
    expect(logoElement).toBeInTheDocument();
})