import React, { useState, ChangeEvent } from "react";

export const Form: React.FC = () => {
    const [phoneNumber, setPhoneNumber] = useState("");
    const [message, setMessage] = useState("");

    const handlePhoneNumberChange = (e: ChangeEvent<HTMLInputElement>) => {
        setPhoneNumber(e.target.value);
    };

    const handleMessageChange = (e: ChangeEvent<HTMLTextAreaElement>) => {
        setMessage(e.target.value);
    };

    const handleSubmit = async (event: React.FormEvent<HTMLFormElement>) => {
        event.preventDefault();
        const formData = { phoneNumber, message }; // Construct the form data object here
        const url = "http://localhost:3000/submit";

        try {
            const response = await fetch(url, {
                method: "POST",
                headers: {
                    "Content-Type": "application/json",
                },
                body: JSON.stringify(formData),
            });

            if (!response.ok) {
                throw new Error("Network response was not ok");
            }

            const data = await response.json();
            console.log(data);
            setPhoneNumber("");
            setMessage("");
            // Provide feedback to the user
        } catch (error) {
            console.error("There was an error submitting the form", error);
            // Provide feedback to the user
        }
    };

    return (
        <div className="p-4 border rounded-3 bg-body-tertiary text-start">
            <form onSubmit={handleSubmit}>
                <div className="mb-3">
                    <input
                        type="tel"
                        className="form-control"
                        id="phoneNumber"
                        placeholder="Phone Number"
                        value={phoneNumber}
                        onChange={handlePhoneNumberChange}
                        name="phoneNumber" />
                </div>
                <div className="mb-3">
                    <textarea
                        className="form-control"
                        placeholder="Enter your message here"
                        id="message"
                        value={message}
                        onChange={handleMessageChange}
                        name="message" />
                </div>
                <button type="submit" className="btn btn-primary">Submit</button>
            </form>
        </div>
    );
};

