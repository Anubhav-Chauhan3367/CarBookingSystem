import React, { useEffect, useState } from "react";
import CarTile from "./CarTile.jsx";
import "./Cars.css";

function Cars() {
	const [cars, setCars] = useState([]);

	useEffect(() => {
		async function fetchCars() {
			try {
				const response = await fetch(
					"http://localhost:8080/api/allcars",
					{
						method: "GET",
						headers: {
							"Content-Type": "application/json",
						},
						credentials: "include",
					}
				);
				if (!response.ok) {
					throw new Error("Failed to fetch cars");
				}
				const data = await response.json();
				setCars(data);
				console.log("Cars fetched successfully:", data);
			} catch (error) {
				console.error("Error fetching cars:", error);
			}
		}

		fetchCars();
	}, []);

	return (
		<div className="cars-grid">
			{cars.map((car) => (
				<CarTile key={car.id} car={car} />
			))}
		</div>
	);
}

export default Cars;
