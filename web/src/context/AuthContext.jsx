import React, { createContext, useContext, useState, useEffect } from "react";

export const AuthContext = createContext();

export const useAuth = () => useContext(AuthContext);

export const AuthProvider = ({ children }) => {
	const [user, setUser] = useState(null);
	const [loading, setLoading] = useState(true);
	const login = async (username, password) => {
		try {
			const response = await fetch("http://localhost:8080/login", {
				method: "POST",
				headers: {
					"Content-Type": "application/json",
				},
				body: JSON.stringify({ username, password }),
				credentials: "include", // Include credentials (cookies) in the request
			});

			if (response.ok) {
				// Check if login was successful by checking response status
				const userData = await response.json();
				setUser(userData);
			} else {
				console.log("Login failed", response);
				throw new Error("Login failed");
			}
		} catch (error) {
			console.error("Login error:", error);
			throw error;
		}
	};

	const logout = async () => {
		try {
			// Send a request to the server to clear the session
			await fetch("http://localhost:8080/logout", {
				method: "POST",
				credentials: "include",
			});

			// Remove the session cookie from the browser
			document.cookie =
				"user-session=; expires=Thu, 01 Jan 1970 00:00:00 UTC; path=/;";

			// Update the user state
			setUser(null);
			console.log("Logout successful");
		} catch (error) {
			console.error("Logout error:", error);
		}
	};

	const register = async (username, email, password) => {
		try {
			const response = await fetch("http://localhost:8080/register", {
				method: "POST",
				headers: {
					"Content-Type": "application/json",
				},
				body: JSON.stringify({ username, email, password }),
				credentials: "include", // Include credentials (cookies) in the request
			});

			if (!response.ok) {
				throw new Error("Registration failed");
			}

			// Assuming you receive some response data (e.g., user information)
			const responseData = await response.json();
			return responseData;
		} catch (error) {
			console.error("Registration error:", error);
			throw error;
		}
	};

	const checkAuthentication = async () => {
		try {
			const response = await fetch("http://localhost:8080/check-auth", {
				// Updated URL
				method: "GET",
				credentials: "include",
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

	useEffect(() => {
		checkAuthentication();
	}, []);

	return (
		<AuthContext.Provider
			value={{
				user,
				login,
				logout,
				register,
				loading,
			}}
		>
			{children}
		</AuthContext.Provider>
	);
};
