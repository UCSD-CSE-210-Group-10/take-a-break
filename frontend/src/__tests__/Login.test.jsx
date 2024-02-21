import { render, screen } from '@testing-library/react'
import App from '../App';

test("Sections Render Successfully", () => {
    render(<App/>);
    const leftSec = screen.getByTestId('left-section');
    const rightSec = screen.getByTestId('right-section');
    expect(leftSec).toBeInTheDocument();
    expect(rightSec).toBeInTheDocument();
})
