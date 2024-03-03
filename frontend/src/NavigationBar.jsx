import Container from "react-bootstrap/Container";
import Nav from "react-bootstrap/Nav";
import Navbar from "react-bootstrap/Navbar";
import NavDropdown from "react-bootstrap/NavDropdown";
import {Link} from 'react-router-dom';
import logo from "./UCSD-logo.png";
import "./NavigationBar.css";

function NavigationBar() {
  return (
    <Navbar expand="lg" className="bg-body-tertiary" data-testid="navigation-bar">
      <Container>
        <img src={logo} alt="UCSD Logo" className="ucsd-logo" />
        <Navbar.Toggle aria-controls="basic-navbar-nav" />
        <Navbar.Collapse id="basic-navbar-nav">
          <Nav className="me-auto">
            <Nav.Link as={Link} to="/events">Events</Nav.Link>
            <Nav.Link as={Link} to="/health">Health</Nav.Link>
            <Nav.Link as={Link} to="/friends">Friends</Nav.Link>
          </Nav>
          <Nav className="justify-content-end">
            <NavDropdown title="Student" id="basic-nav-dropdown">
              <NavDropdown.Item href="/profile">Profile</NavDropdown.Item>
              <NavDropdown.Divider />
              <NavDropdown.Item href="/logout">Log out</NavDropdown.Item>
            </NavDropdown>
          </Nav>
        </Navbar.Collapse>
      </Container>
    </Navbar>
  );
}

export default NavigationBar;
