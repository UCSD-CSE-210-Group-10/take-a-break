import { render, screen } from '@testing-library/react'
import NavigationBar from '../NavigationBar'
import { MemoryRouter } from "react-router-dom";

test("Navigation Bar Renders Successfully", () => {
    render(<MemoryRouter><NavigationBar/></MemoryRouter>);
    
    const leftSec = screen.getByTestId('navigation-bar');
    expect(leftSec).toBeInTheDocument();
})

test("Links Render Successfully", () => {
    render(<MemoryRouter><NavigationBar/></MemoryRouter>);

  
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
    render(<MemoryRouter><NavigationBar/></MemoryRouter>);

  
    const logoElement = screen.getByAltText('UCSD Logo');
    expect(logoElement).toBeInTheDocument();
})