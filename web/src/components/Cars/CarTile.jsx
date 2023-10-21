import React from "react";
import { Link } from "react-router-dom";
import "./CarTile.css";

function CarTile({ car }) {
	return (
		<Link to={`/cars/${car.id}`} className="car-tile-link">
			<div className="car-tile">
				<img src={car.imageurl} alt={`${car.brand} ${car.model}`} />
				<div className="car-details">
					<h2>
						{car.brand} {car.model}
					</h2>
					<p>{car.description}</p>
				</div>
			</div>
		</Link>
	);
}

export default CarTile;
