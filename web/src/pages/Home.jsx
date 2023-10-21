import React from "react";
import CarCarousel from "../components/Carousel";
import "./Home.css";

function Home() {
	return (
		<div className="home-container" id="home">
			<div className="welcome">
				<h1>Car Rental</h1>
				<h2>
					We are a car rental company that offers a wide range of
					vehicles for rent.
				</h2>
			</div>
			<div>
				<CarCarousel />
			</div>
		</div>
	);
}

export default Home;
