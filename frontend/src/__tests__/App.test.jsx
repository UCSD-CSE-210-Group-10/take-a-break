import { render, screen } from '@testing-library/react'
import App from '../App';

test("App renders successfully", () => {
    render(<App/>);
    const element = screen.getByTestId('app-head');
    expect(element).toBeInTheDocument();
})

test('Login renders successfully', () => {
    render(<App />);
    const element = screen.getByText('Take a Break');
    expect(element).toBeInTheDocument();
  });