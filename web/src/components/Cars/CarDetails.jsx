import React, { useEffect, useState } from "react";
import { useParams } from "react-router-dom";
import { Link } from "react-router-dom";
import "./CarDetails.css";

function CarDetail(props) {
	const { carId } = useParams();
	const [car, setCar] = useState(null);
	const [buttonIsActive, setButtonIsActive] = useState(true);
	if (props.buttonIsActive === false) {
		setButtonIsActive(false);
	}
	useEffect(() => {
		async function fetchCarDetails() {
			try {
				const response = await fetch(
					`http://localhost:8080/api/cars/${carId}`,
					{
						method: "GET",
						headers: {
							"Content-Type": "application/json",
						},
						credentials: "include",
					}
				);
				if (!response.ok) {
					throw new Error("Failed to fetch car details");
				}
				const data = await response.json();
				setCar(data);
				console.log("Car details fetched successfully:", data);
			} catch (error) {
				console.error("Error fetching car details:", error);
			}
		}

		fetchCarDetails();
	}, [carId]);

	if (!car) {
		return <div>Loading...</div>;
	}

	return (
		<div className="car-detail">
			<img src={car.imageurl} alt={`${car.brand} ${car.model}`} />
			<div className="car-details">
				<h2>
					{car.brand} {car.model}
				</h2>
				<p>{car.description}</p>
				{buttonIsActive && (
					<>
						<Link
							to={`/bookings/${car.id}`}
							className="book-button"
						>
							Available Slots
						</Link>
					</>
				)}
			</div>
		</div>
	);
}

export default CarDetail;
