import React, { useState } from "react";
import { useAuth } from "../../context/AuthContext";

import "./Auth.css";

const Register = () => {
	const { login } = useAuth();
	const [formData, setFormData] = useState({
		username: "",
		email: "",
		password: "",
	});
	const [error, setError] = useState(null);

	const handleInputChange = (e) => {
		const { name, value } = e.target;
		setFormData({
			...formData,
			[name]: value,
		});
	};

	const handleSubmit = async (e) => {
		e.preventDefault();
		try {
			// Here, you should send a registration request to your backend API
			// and handle user creation and login in your backend.
			// For simplicity, we'll assume a successful registration logs in the user.
			await login(formData.username, formData.password);
		} catch (error) {
			setError("Registration failed. Please try again.");
		}
	};

	return (
		<div className="container background-image">
			<div className="form-container">
				<h2>Register</h2>
				{error && <div className="error">{error}</div>}
				<form onSubmit={handleSubmit}>
					<div>
						<label htmlFor="username">Username</label>
						<input
							type="text"
							id="username"
							name="username"
							value={formData.username}
							onChange={handleInputChange}
							required
						/>
					</div>
					<div>
						<label htmlFor="email">Email</label>
						<input
							type="email"
							id="email"
							name="email"
							value={formData.email}
							onChange={handleInputChange}
							required
						/>
					</div>
					<div>
						<label htmlFor="password">Password</label>
						<input
							type="password"
							id="password"
							name="password"
							value={formData.password}
							onChange={handleInputChange}
							required
						/>
					</div>
					<button type="submit">Register</button>
				</form>
			</div>
		</div>
	);
};

export default Register;
