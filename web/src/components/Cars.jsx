import React, { useEffect } from "react";
// import { Link } from "react-router-dom";

function Cars() {
	// Use state to store car data fetched from the backend

	useEffect(() => {
		// Fetch car data from the backend
	}, []);

	return (
		<div>
			<h2>Available Cars</h2>
			{/* Render a list of cars with links to CarDetails */}
		</div>
	);
}

export default Cars;
