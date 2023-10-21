import React, { useEffect, useState } from "react";
import { Link, useLocation } from "react-router-dom";
import { useAuth } from "../../context/AuthContext";
import "./Navbar.css";

function Navbar() {
	const { user } = useAuth();
	const location = useLocation();
	const [isNavbarHidden, setNavbarHidden] = useState(false);

	const isActive = (path) => {
		return location.pathname === path ? "active" : "";
	};

	useEffect(() => {
		let timeoutId;

		const handleMouseMovement = () => {
			setNavbarHidden(false);
			clearTimeout(timeoutId);
			timeoutId = setTimeout(() => {
				setNavbarHidden(true);
			}, 2000); // Adjust the delay as needed
		};

		handleMouseMovement();

		window.addEventListener("mousemove", handleMouseMovement);

		return () => {
			window.removeEventListener("mousemove", handleMouseMovement);
		};
	}, []);

	return (
		<nav className={`navbar ${isNavbarHidden ? "hidden" : ""}`}>
			<ul className="nav-links">
				<li>
					<Link to="/" className={isActive("/")}>
						Home
					</Link>
				</li>
				<li>
					<Link to="/about" className={isActive("/about")}>
						About
					</Link>
				</li>
				<li>
					<Link to="/contact" className={isActive("/contact")}>
						Contact
					</Link>
				</li>
				{user ? (
					<>
						<li>
							<Link to="/cars" className={isActive("/cars")}>
								Cars
							</Link>
						</li>
						<li>
							<Link to="/logout" className={isActive("/logout")}>
								Logout
							</Link>
						</li>
					</>
				) : (
					<>
						<li>
							<Link to="/signup" className={isActive("/signup")}>
								Sign Up
							</Link>
						</li>
						<li>
							<Link to="/login" className={isActive("/login")}>
								Login
							</Link>
						</li>
					</>
				)}
			</ul>
		</nav>
	);
}

export default Navbar;
