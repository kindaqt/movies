import React from 'react';
import { Link } from 'react-router-dom'

function Header() {
  return (
    <header style={headerStyle}>
      <h1>Movies</h1>
      <nav>
          <Link style={linkStyle} to="/">Home</Link>{" | "} 
          <Link style={linkStyle} to="/about">About</Link>{" | "} 
          <Link style={linkStyle} to="/health">Health</Link>
      </nav>
    </header>
  )
}

const headerStyle = {
  background: '#333',
  color: '#fff',
  textAlign: 'center',
  padding: '10px',
}

const linkStyle = {
  color: "#fff",
  textDecoration: "none",
}

export default Header;