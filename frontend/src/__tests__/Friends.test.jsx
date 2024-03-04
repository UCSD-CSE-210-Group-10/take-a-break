import { render, screen } from '@testing-library/react'
import Friends from '../Friends'
import { MemoryRouter } from "react-router-dom";
import FriendCard from '../FriendCard';

test("Friend Card Renders Successfully", () => {
    const sampleFriend =  {
        id: 1,
        name: 'John Doe',
        image: './UCSD-logo.png',
    };

    render(<MemoryRouter><FriendCard key={sampleFriend.id} friend={sampleFriend} /></MemoryRouter>);    
    const friendCard = screen.getByTestId('friend-card');
    expect(friendCard).toBeInTheDocument();
})

test("Friends Page Renders Successfully", () => {
    render(<MemoryRouter><Friends/></MemoryRouter>);

    const friendContainer = screen.getByTestId('friends-container');
    expect(friendContainer).toBeInTheDocument();
})
