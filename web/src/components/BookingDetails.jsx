import React, { useState, useEffect } from "react";
import { useParams } from "react-router-dom";
import DatePicker from "react-datepicker";
import "react-datepicker/dist/react-datepicker.css";
import "./BookingDetails.css";
import "./Cars/CarDetails.css";
import { DateTime } from "luxon";
import BookingForm from "./BookingForm";

function BookingDetails() {
	const { carId } = useParams();
	const [requiredCar, setRequiredCar] = useState({}); // [carId]
	const [selectedDate, setSelectedDate] = useState(null);
	const [selectedSlot, setSelectedSlot] = useState(null);
	const [availability, setAvailability] = useState({});
	console.log(selectedSlot, "selectedSlot");
	useEffect(() => {
		async function fetchCarAvailabilityDetails() {
			try {
				const response = await fetch(`http://localhost:8080/`, {
					method: "GET",
					headers: {
						"Content-Type": "application/json",
					},
					credentials: "include",
				});

				if (!response.ok) {
					throw new Error("Failed to fetch car availability details");
				}

				const data = await response.json();
				const required_Car = data.find((obj) => obj.car.id == carId);
				// console.log(data, "data", required_Car, "requiredCar");

				if (required_Car) {
					const availabilityData = required_Car.availability;
					const formattedAvailability =
						populateAvailableSlots(availabilityData);
					setAvailability(formattedAvailability);
					setRequiredCar(required_Car.car);
				}
			} catch (error) {
				console.error(
					"Error fetching car availability details:",
					error
				);
			}
		}

		fetchCarAvailabilityDetails();
	}, [carId]);

	// console.log(availability, "availability");
	function populateAvailableSlots(timeIntervals) {
		const availableSlots = {};

		timeIntervals.forEach((interval) => {
			const start = DateTime.fromISO(interval.start_time);
			const end = DateTime.fromISO(interval.end_time);

			let current = start.startOf("hour");
			while (current <= end) {
				const date = current.toFormat("yyyy-LL-dd");
				const time = current.toFormat("HH:mm:ss");

				if (!availableSlots[date]) {
					availableSlots[date] = [];
				}

				availableSlots[date].push(time);

				current = current.plus({ hours: 1 });
			}
		});

		return availableSlots;
	}

	const isSlotAvailable = (date, timeSlot) => {
		// Parse the selected date in the "Asia/Kolkata" time zone
		const kolkataDate = DateTime.fromISO(date.toISOString(), {
			zone: "Asia/Kolkata",
		});

		// Format the date without converting to UTC
		const formattedDate = kolkataDate.toISODate();

		if (
			availability[formattedDate] &&
			availability[formattedDate].includes(timeSlot)
		) {
			return true;
		}
		return false;
	};

	return (
		<div className="booking-details-container">
			<h2>Booking Details</h2>
			<div className="calendar-container">
				<DatePicker
					selected={selectedDate}
					onChange={(date) => setSelectedDate(date)}
					inline
					minDate={new Date()}
					maxDate={
						new Date(new Date().setMonth(new Date().getMonth() + 1))
					}
				/>
				<div className="time-slots-container">
					{selectedDate &&
						Array.from({ length: 24 }, (_, index) => {
							const hour = index.toString().padStart(2, "0");
							const timeSlot = `${hour}:00:00`;
							const isAvailable = isSlotAvailable(
								selectedDate,
								timeSlot
							);
							// console.log(timeSlot, " ", hour, " ", isAvailable);
							return (
								<div
									key={timeSlot}
									className={`time-slot ${
										isAvailable
											? "available"
											: "unavailable"
									}`}
									onClick={() => setSelectedSlot(timeSlot)}
								>
									{timeSlot}
								</div>
							);
						})}
				</div>
			</div>
			<div className="car-detail">
				<img
					src={requiredCar.imageurl}
					alt={`${requiredCar.brand} ${requiredCar.model}`}
				/>
				<div className="car-details">
					<h2>
						{requiredCar.brand} {requiredCar.model}
					</h2>
					<p>{requiredCar.description}</p>
				</div>
			</div>
			<BookingForm
				carId={carId}
				carName={`${requiredCar.brand} ${requiredCar.model}`}
				availability={availability}
			/>
		</div>
	);
}

export default BookingDetails;
