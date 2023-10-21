import React, { useState } from "react";
import { useAuth } from "../../context/AuthContext";
import "./Auth.css";

function Logout() {
	const { logout } = useAuth();
	const [isConfirmOpen, setConfirmOpen] = useState(true);

	const handleLogout = () => {
		if (isConfirmOpen) {
			setConfirmOpen(false);
			logout();
		} else {
			setConfirmOpen(true);
		}
	};

	const handleCancel = () => {
		setConfirmOpen(false);
	};

	return (
		<div className="logout-container">
			{isConfirmOpen && (
				<div className="confirmation-box">
					<p>Are you sure you want to logout?</p>
					<button onClick={handleLogout} className="confirm-logout">
						Yes
					</button>
					<button onClick={handleCancel} className="cancel-logout">
						No
					</button>
				</div>
			)}
		</div>
	);
}

export default Logout;
