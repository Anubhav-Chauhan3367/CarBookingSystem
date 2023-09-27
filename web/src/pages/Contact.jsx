import React from "react";
import "./Contact.css";

function Contact() {
	return (
		<div className="container">
			<div className="contact-container">
				<h1>Contact Us</h1>
				<p>
					If you have any questions or inquiries, please don't
					hesitate to contact us using the information below.
				</p>
				<div className="contact-info">
					<div className="contact-item">
						<h2>Email:</h2>
						<p>contact@example.com</p>
					</div>
					<div className="contact-item">
						<h2>Phone:</h2>
						<p>+1 (123) 456-7890</p>
					</div>
					<div className="contact-item">
						<h2>Address:</h2>
						<p>123 Street Name, City, Country</p>
					</div>
				</div>
			</div>
		</div>
	);
}

export default Contact;
