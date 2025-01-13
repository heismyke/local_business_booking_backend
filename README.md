# **Local Business Booking App**  

## **Overview**  
In many neighborhoods, small businesses such as **barbershops, hair salons, beauty spas, and nail studios** provide essential services. However, customers often struggle to find and book appointments efficiently. Many of these businesses rely on walk-ins, phone calls, or social media for bookings, leading to:  

- 🚫 **Inconvenience** for customers who have to visit multiple locations or call ahead to check availability.  
- 💰 **Lost revenue** for businesses due to inefficient scheduling and missed appointments.  
- 🔍 **Lack of visibility** for local businesses, making it difficult for new customers to find them.  

## **Problem Statement**  
There is no **centralized booking platform** that allows users to:  
✅ Discover businesses on their street.  
✅ View available services.  
✅ Book appointments seamlessly.  

Small business owners **lack an easy-to-use system** to manage bookings, reduce no-shows, and improve customer engagement.  

---

## **Proposed Solution**  
This project aims to develop a **mobile application** that:  

1️⃣ **Business Discovery** – Allows users to search for businesses on a specific street using **Google Maps**.  
2️⃣ **Auto-fetch Business Details** – Retrieves business info (**name, phone, services**) from **Google Places API**.  
3️⃣ **Online Booking System** – Enables customers to **instantly book** services.  
4️⃣ **Business Dashboard** – Allows owners to **manage bookings, set working hours, and send reminders**.  
5️⃣ **Notifications & Reminders** – Sends **email/SMS alerts** to confirm appointments and reduce no-shows.  

---

## **Key Features & Benefits**  
✅ **Convenience** – Users can quickly find and book services nearby.  
✅ **Increased Visibility** – Small businesses get **more exposure** to new customers.  
✅ **Efficient Scheduling** – Automated bookings reduce **missed appointments**.  
✅ **Improved Customer Engagement** – Notifications & reminders enhance the user experience.  

---

## **Tech Stack**  
- **Backend:** Golang, PostgreSQL, SQLC  
- **Frontend:** React Native (Planned)  
- **APIs:** Google Maps API, Google Places API, Twilio (for SMS), SendGrid (for emails)  
- **Infrastructure:** Docker, GitHub Actions (CI/CD)  

---

## **Setup & Installation**  
### **1. Clone the Repository**  
```sh
git clone https://github.com/heismyke/local_business_booking_backend.git
cd local_business_booking_backend
```

### **2. Install Dependencies**  
```sh
go mod tidy
```

### **3. Setup Environment Variables**  
Create a `.env` file and configure:  
```env
DB_URL=postgres://username:password@localhost:5432/local_booking?sslmode=disable
GOOGLE_API_KEY=your_google_api_key
TWILIO_SID=your_twilio_sid
TWILIO_AUTH_TOKEN=your_twilio_auth_token
SENDGRID_API_KEY=your_sendgrid_api_key
```

### **4. Run the Application**  
```sh
go run main.go
```

---

## **Running Tests**  
```sh
make test
```

---

## **Contributing**  
We welcome contributions! Please follow these steps:  

1️⃣ **Fork** the repository.  
2️⃣ Create a new **feature branch**:  
   ```sh
   git checkout -b feature-name
   ```
3️⃣ **Commit** your changes:  
   ```sh
   git commit -m "Add new feature"
   ```
4️⃣ **Push** to GitHub:  
   ```sh
   git push origin feature-name
   ```
5️⃣ Open a **Pull Request** 🚀  

---

## **License**  
This project is licensed under the **MIT License**.  

---

## **Contact**  
👤 **Author:** [heismyke](https://github.com/heismyke)  
📧 **Email:** 81974207+heismyke@users.noreply.github.com  

---

### **🚀 Let's revolutionize local business bookings!**  

---

This README provides a **clear structure**, making it easy for contributors and users to understand your project. 🚀 Let me know if you need modifications!
