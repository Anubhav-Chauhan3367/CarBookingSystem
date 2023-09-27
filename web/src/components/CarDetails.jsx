import React, { useEffect } from "react";
import { useParams } from "react-router-dom";

function CarDetails() {
	const { carId } = useParams();
	// Use state to store car data fetched from the backend

	useEffect(() => {
		// Fetch car data for the specified carId from the backend
	}, [carId]);

	return <div>{/* Render car details */}</div>;
}

export default CarDetails;
