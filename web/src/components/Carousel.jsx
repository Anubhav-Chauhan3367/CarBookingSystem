// CarCarousel.js

import React, { useEffect, useState } from "react";
import { Carousel } from "react-responsive-carousel";
import "react-responsive-carousel/lib/styles/carousel.min.css";
import "./Carousel.css";

function CarCarousel() {
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

	const handleCarClick = (car) => {
		console.log("Car clicked:", car);
	};

	return (
		<div>
			<Carousel
				showArrows={false}
				showThumbs={true}
				showStatus={false}
				infiniteLoop={true}
				autoPlay={true}
				interval={4000}
			>
				{cars.map((car) => (
					<div
						className="carousel-slide"
						key={car.id}
						onClick={() => handleCarClick(car)}
					>
						<h2>
							{car.brand} {car.model}
						</h2>
						<img
							src={car.imageurl}
							alt={`${car.brand} ${car.model}`}
						/>
					</div>
				))}
			</Carousel>
		</div>
	);
}

export default CarCarousel;
