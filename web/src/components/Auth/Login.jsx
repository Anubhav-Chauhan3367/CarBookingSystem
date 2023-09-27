import React, { useState } from "react";
import { useAuth } from "../../context/AuthContext";

import "./Auth.css";

const Login = () => {
	const { login } = useAuth();
	const [formData, setFormData] = useState({
		username: "",
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
			await login(formData.username, formData.password);
		} catch (error) {
			setError("Login failed. Please check your credentials.");
		}
	};

	return (
		<div className="container background-image">
			<div className="form-container">
				<h2>Login</h2>
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
					<button type="submit">Login</button>
				</form>
			</div>
		</div>
	);
};

export default Login;
