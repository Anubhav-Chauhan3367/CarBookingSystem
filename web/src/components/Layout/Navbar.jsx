import React from "react";
import { NavLink } from "react-router-dom"; // Import NavLink

import "./Navbar.css";

function Navbar() {
	return (
		<nav>
			<ul>
				<li>
					<NavLink exact to="/" activeClassName="active">
						{" "}
						{/* Use NavLink */}
						Home
					</NavLink>
				</li>
				<li>
					<NavLink to="/about" activeClassName="active">
						About
					</NavLink>
				</li>
				<li>
					<NavLink to="/contact" activeClassName="active">
						Contact
					</NavLink>
				</li>
				<li>
					<NavLink to="/cars" activeClassName="active">
						Cars
					</NavLink>
				</li>
				<li>
					<NavLink to="/signup" activeClassName="active">
						Sign Up
					</NavLink>
				</li>
				<li>
					<NavLink to="/login" activeClassName="active">
						Login
					</NavLink>
				</li>
				{/* Add more navigation links as needed */}
			</ul>
		</nav>
	);
}

export default Navbar;
