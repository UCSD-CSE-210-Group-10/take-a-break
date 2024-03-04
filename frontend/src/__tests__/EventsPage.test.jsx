import { render, screen } from '@testing-library/react';
import EventsPage from '../EventsPage';

test("NavigationBar Renders Successfully", () => {
    render(<EventsPage />);
    
    const navigationBar = screen.getByTestId('navigation-bar');
    expect(navigationBar).toBeInTheDocument();
});

test("Search Bar Renders Successfully", () => {
    render(<EventsPage />);
    
    const searchBar = screen.getByPlaceholderText('Search Event');
    expect(searchBar).toBeInTheDocument();

    const tagsDropdown = screen.getByText('Tags');
    expect(tagsDropdown).toBeInTheDocument();
});

test("Event Cards Render Successfully", () => {
    render(<EventsPage />);
    
    const eventCardLinks = screen.getAllByRole('link', { name: /Event \d/ });
    eventCardLinks.forEach(link => {
        expect(link).toBeInTheDocument();
    });
});

test("Event Details Render Correctly", () => {
    render(<EventsPage />);
    
    const eventNames = screen.getAllByRole('heading', { level: 3 });
    eventNames.forEach(name => {
        expect(name).toBeInTheDocument();
    });

    const eventDates = screen.getAllByText(/2024/);
    eventDates.forEach(date => {
        expect(date).toBeInTheDocument();
    });

    const eventTimes = screen.getAllByText(/AM|PM/);
    eventTimes.forEach(time => {
        expect(time).toBeInTheDocument();
    });

    const eventOrganizations = screen.getAllByText(/Organization/);
    eventOrganizations.forEach(organization => {
        expect(organization).toBeInTheDocument();
    });
});
