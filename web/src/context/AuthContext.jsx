import React, { createContext, useContext, useState, useEffect } from "react";

// Create a context to manage authentication state and provide it to components
export const AuthContext = createContext();

// Custom hook to access the authentication context
export const useAuth = () => useContext(AuthContext);

// AuthProvider component to manage authentication state
export const AuthProvider = ({ children }) => {
	const [user, setUser] = useState(null);
	const [loading, setLoading] = useState(true);

	// Function to log in a user
	const login = async (username, password) => {
		try {
			const response = await fetch("/api/login", {
				method: "POST",
				headers: {
					"Content-Type": "application/json",
				},
				body: JSON.stringify({ username, password }),
			});

			if (response.ok) {
				const userData = await response.json();
				setUser(userData);
			} else {
				throw new Error("Login failed");
			}
		} catch (error) {
			console.error("Login error:", error);
			throw error;
		}
	};

	// Function to log out a user
	const logout = async () => {
		try {
			await fetch("/api/logout", {
				method: "POST",
			});
			setUser(null);
		} catch (error) {
			console.error("Logout error:", error);
		}
	};

	// Check if the user is authenticated on component mount
	useEffect(() => {
		const checkAuthentication = async () => {
			try {
				const response = await fetch("/api/check-auth", {
					method: "GET",
				});

				if (response.ok) {
					const userData = await response.json();
					setUser(userData);
				}
			} catch (error) {
				console.error("Authentication check error:", error);
			} finally {
				setLoading(false);
			}
		};

		checkAuthentication();
	}, []);

	return (
		<AuthContext.Provider value={{ user, login, logout, loading }}>
			{children}
		</AuthContext.Provider>
	);
};
