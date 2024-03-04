import { render, screen } from '@testing-library/react'
import Friends from '../Friends'
import { MemoryRouter } from "react-router-dom";
import FriendCard from '../FriendCard';

test("Friend Card Renders Successfully", () => {
    render(<MemoryRouter><FriendCard/></MemoryRouter>);
    
    const leftSec = screen.getByTestId('friend-card');
    expect(leftSec).toBeInTheDocument();
})

test("Friends Page Renders Successfully", () => {
    render(<MemoryRouter><Friends/></MemoryRouter>);

    const leftSec = screen.getByTestId('friends-container');
    expect(leftSec).toBeInTheDocument();
})
