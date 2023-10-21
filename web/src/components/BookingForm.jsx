import React, { useState, useEffect } from "react";
import DatePicker from "react-datepicker";
import "react-datepicker/dist/react-datepicker.css";
import "./BookingForm.css";
import { DateTime } from "luxon";
import { useAuth } from "../context/AuthContext";

function BookingForm(props) {
	const { user } = useAuth();

	const [car, setCar] = useState(null);
	const [error, setError] = useState(""); // Error message for booking form
	const [successMessage, setSuccessMessage] = useState("");
	const [startDate, setStartDate] = useState(null);
	const [endDate, setEndDate] = useState(null);
	const [startError, setStartError] = useState(""); // Error message for start date
	const [endError, setEndError] = useState(""); // Error message for end date

	useEffect(() => {
		async function fetchRequiredCar() {
			try {
				const response = await fetch(
					`http://localhost:8080/api/cars/${props.carId}`,
					{
						method: "GET",
						headers: {
							"Content-Type": "application/json",
						},
						credentials: "include",
					}
				);
				if (!response.ok) {
					throw new Error("Failed to fetch car");
				}
				const data = await response.json();
				setCar(data);
				console.log("Required Car fetched successfully:", car);
			} catch (error) {
				setCar(null);
				console.error("Error fetching car:", error);
			}
		}

		fetchRequiredCar();
	}, [props.carId]);

	const isSlotAvailable = (date, timeSlot) => {
		const kolkataDate = DateTime.fromISO(date.toISOString(), {
			zone: "Asia/Kolkata",
		});
		const formattedDate = kolkataDate.toISODate();

		if (
			props.availability[formattedDate] &&
			props.availability[formattedDate].includes(timeSlot)
		) {
			return true;
		}
		return false;
	};

	const handleStartDateChange = (date) => {
		setStartDate(date);

		// Extract hours, minutes, and seconds from the selected date
		const hours = date.getHours().toString().padStart(2, "0");
		const minutes = date.getMinutes().toString().padStart(2, "0");
		const seconds = date.getSeconds().toString().padStart(2, "0");

		// Format the time as "00:00:00"
		const selectedTime = `${hours}:${minutes}:${seconds}`;
		console.log("Selected Time:", selectedTime);

		// Check if the selected start date is available
		if (!isSlotAvailable(date, selectedTime)) {
			setStartError("Selected start date is not available.");
		} else if (endDate && startDate >= endDate) {
			setStartError("Start date should be lower than end date.");
		} else {
			setStartError("");
		}
	};

	const handleEndDateChange = (date) => {
		setEndDate(date);

		// Extract hours, minutes, and seconds from the selected date
		const hours = date.getHours().toString().padStart(2, "0");
		const minutes = date.getMinutes().toString().padStart(2, "0");
		const seconds = date.getSeconds().toString().padStart(2, "0");

		// Format the time as "00:00:00"
		const selectedTime = `${hours}:${minutes}:${seconds}`;
		console.log("Selected Time:", selectedTime);

		// Check if the selected start date is available
		if (!isSlotAvailable(date, selectedTime)) {
			setEndError("Selected end date is not available.");
		} else if (startDate && startDate >= endDate) {
			setEndError("End date should be greater than start date.");
		} else {
			setEndError("");
		}
	};

	const handleSubmit = async (event) => {
		event.preventDefault();
		if (!user) {
			setStartError("User data not available. Please log in.");
			return;
		}

		// Check if there are any errors before submitting the form
		if (startError || endError) {
			return;
		}

		// Create a booking object with the form data
		const bookingData = {
			user_id: +user.id,
			car_id: +props.carId,
			start_time: startDate.toISOString(),
			end_time: endDate.toISOString(),
		};

		try {
			// Send a request to your backend to create a booking using the bookingData
			const response = await fetch("http://localhost:8080/api/bookings", {
				method: "POST",
				headers: {
					"Content-Type": "application/json",
				},
				body: JSON.stringify(bookingData),
				credentials: "include",
			});

			if (!response.ok) {
				throw new Error("Failed to create booking");
			}

			// Reset the form fields
			setStartDate(null);
			setEndDate(null);
			setError("");
			setSuccessMessage("Booking successful!");
		} catch (error) {
			setError("Error creating booking: " + error.message);
			console.error("Error creating booking:", error);
		}
	};

	return (
		<div>
			<form className="booking-form" onSubmit={handleSubmit}>
				<div className="form-group">
					<label htmlFor="username">Username</label>
					<input
						type="text"
						id="username"
						value={user ? user.username : ""}
						disabled
					/>
				</div>
				<div className="form-group">
					<input
						type="number"
						id="userId"
						value={user ? user.id : ""}
						disabled
						hidden
					/>
				</div>
				<div className="form-group">
					<label htmlFor="carId">Car ID</label>
					<input
						type="text"
						id="carId"
						value={props.carId}
						disabled
						hidden
					/>
				</div>
				<div className="form-group">
					<label htmlFor="carName">Car Name</label>
					<input
						type="text"
						id="carName"
						value={car ? car.brand + " " + car.model : ""}
						disabled
					/>
				</div>
				<div className="form-group">
					<label htmlFor="startDate">
						<i className="fas fa-calendar-alt"></i> Start Date
					</label>
					<DatePicker
						selected={startDate}
						onChange={handleStartDateChange}
						showTimeSelect
						minDate={new Date()}
						maxDate={
							new Date(
								new Date().setMonth(new Date().getMonth() + 1)
							)
						}
						timeFormat="HH:mm"
						timeIntervals={60}
						timeCaption="Time"
						dateFormat="MMMM d, yyyy h:mm aa"
						id="startDate"
						required
					/>
					{startError && <div className="error">{startError}</div>}
				</div>
				<div className="form-group">
					<label htmlFor="endDate">
						<i className="fas fa-calendar-alt"></i> End Date
					</label>
					<DatePicker
						selected={endDate}
						onChange={handleEndDateChange}
						showTimeSelect
						minDate={new Date()}
						maxDate={
							new Date(
								new Date().setMonth(new Date().getMonth() + 1)
							)
						}
						timeFormat="HH:mm"
						timeIntervals={60}
						timeCaption="Time"
						dateFormat="MMMM d, yyyy h:mm aa"
						id="endDate"
						required
					/>
					{endError && <div className="error">{endError}</div>}
				</div>
				<button type="submit">Book Now</button>
				{error && <div className="error">{error}</div>}
				{successMessage && (
					<div className="success">{successMessage}</div>
				)}
			</form>
		</div>
	);
}

export default BookingForm;
