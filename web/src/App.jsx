import React from "react";
import {
	BrowserRouter as Router,
	Route,
	Routes,
	Navigate,
} from "react-router-dom";
import Home from "./pages/Home";
import About from "./pages/About";
import Contact from "./pages/Contact";
import Cars from "./components/Cars";
import CarDetails from "./components/CarDetails";
import Bookings from "./components/Bookings";
import BookingDetails from "./components/BookingDetails";
import Navbar from "./components/Layout/Navbar";
import Login from "./components/Auth/Login"; // Import the Login component
import Register from "./components/Auth/Register"; // Import the Signup component

import "./styles/main.css";

function App() {
	return (
		<Router>
			<Navbar />
			<Routes>
				<Route exact path="/" element={<Home />} />
				<Route path="/about" element={<About />} />
				<Route path="/contact" element={<Contact />} />
				<Route exact path="/cars" element={<Cars />} />
				<Route path="/cars/:carId" element={<CarDetails />} />
				<Route exact path="/bookings" element={<Bookings />} />
				<Route
					path="/bookings/:bookingId"
					element={<BookingDetails />}
				/>
				{/* Add Login and Signup routes */}
				<Route path="/login" element={<Login />} />
				<Route path="/signup" element={<Register />} />
				<Route path="*" element={<Navigate to="/" />} />
			</Routes>
		</Router>
	);
}

export default App;
