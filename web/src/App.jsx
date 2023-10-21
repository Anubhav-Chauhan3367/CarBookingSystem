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
import Cars from "./components/Cars/Cars";
import CarDetail from "./components/Cars/CarDetails";
import CarForm from "./components/Cars/CarForm";
import Bookings from "./components/Bookings";
import BookingDetails from "./components/BookingDetails";
import Navbar from "./components/Layout/Navbar";
import Login from "./components/Auth/Login";
import Register from "./components/Auth/Register";
import Logout from "./components/Auth/Logout";

import "./styles/main.css";

function App() {
	return (
		<Router>
			<Navbar />
			<Routes>
				<Route path="/" element={<Home />} />
				<Route path="/about" element={<About />} />
				<Route path="/contact" element={<Contact />} />
				<Route path="/cars" element={<Cars />} />
				<Route path="/cars/:carId" element={<CarDetail />} />
				<Route path="/cars/add" element={<CarForm />} />
				<Route path="/bookings" element={<Bookings />} />
				<Route path="/bookings/:carId" element={<BookingDetails />} />

				{/*Login and Signup routes */}
				<Route path="/login" element={<Login />} />
				<Route path="/signup" element={<Register />} />
				<Route path="/logout" element={<Logout />} />
				<Route path="*" element={<Navigate to="/" />} />
			</Routes>
		</Router>
	);
}

export default App;
