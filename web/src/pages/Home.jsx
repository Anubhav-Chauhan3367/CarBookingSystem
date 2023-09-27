import React, { useRef, useState } from "react";
import CarCarousel from "../components/Carousel";
import "./Home.css";

function Home() {
	const carouselRef = useRef(null);
	const [showCarousel, setShowCarousel] = useState(false);

	const handleArrowClick = () => {
		if (carouselRef.current) {
			carouselRef.current.scrollIntoView({ behavior: "smooth" });
			setShowCarousel(true); // Update the state to show the carousel
		}
	};

	return (
		<div className="container">
			<div
				className={`home-container ${showCarousel ? "hidden" : ""}`}
				id="home"
			>
				<div className="welcome">
					<h1>Car Rental</h1>
					<h2>
						We are a car rental company that offers a wide range of
						vehicles for rent.
					</h2>
				</div>
			</div>
			<div className="arrow">
				<a href="#carousel" onClick={handleArrowClick}>
					<i className="fas fa-chevron-down"></i>
				</a>
			</div>
			<div
				className={`carousel-container ${
					showCarousel ? "show-carousel" : ""
				}`}
				ref={carouselRef}
				id="carousel"
			>
				<CarCarousel />
			</div>
		</div>
	);
}

export default Home;
