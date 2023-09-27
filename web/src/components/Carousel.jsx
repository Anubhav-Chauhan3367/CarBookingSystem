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
					}
				);
				if (!response.ok) {
					throw new Error("Failed to fetch cars");
				}
				const data = await response.json();
				setCars(data);
			} catch (error) {
				console.error("Error fetching cars:", error);
			}
		}

		fetchCars();
	}, []);

	return (
		<div className="carousel-container">
			<Carousel
				showArrows={true}
				showThumbs={false}
				showStatus={false}
				infiniteLoop={true}
				autoPlay={true}
				interval={5000}
			>
				{cars.map((car) => (
					<div key={car.id}>
						<img
							src={car.imageUrl}
							alt={`${car.brand} ${car.model}`}
						/>
						<h2>
							{car.brand} {car.model}
						</h2>
						<p>{car.description}</p>
					</div>
				))}
			</Carousel>
		</div>
	);
}

export default CarCarousel;
