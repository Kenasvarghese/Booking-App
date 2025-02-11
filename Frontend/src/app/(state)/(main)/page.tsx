"use client";

import React, { useState } from "react";
import { useForm } from "react-hook-form";
import { Calendar } from "@/components/ui/calendar";
import { DateRange } from "react-day-picker";
import { addDays } from "date-fns";
import { BookingModal } from "./components/booking-modal";
import SuccessModal from "@/components/ui/SuccessModal";

const properties = [
  {
    id: 1,
    name: "Dragster Homes",
    image: "https://images.unsplash.com/photo-1566073771259-6a8506099945",
  },
  {
    id: 2,
    name: "Poolvilla",
    image: "https://images.unsplash.com/photo-1582719508461-905c673771fd",
  },
  {
    id: 3,
    name: "Tent Stay",
    image: "https://images.unsplash.com/photo-1542314831-068cd1dbfeeb",
  },
];

const rooms = [
  { id: 1, number: "101", status: "available", type: "Deluxe" },
  { id: 2, number: "102", status: "booked", type: "Suite" },
  { id: 3, number: "103", status: "pending", type: "Standard" },
  { id: 4, number: "104", status: "available", type: "Deluxe" },
  { id: 5, number: "105", status: "booked", type: "Suite" },
  { id: 6, number: "106", status: "pending", type: "Standard" },
];

/**
 * A dummy function that returns a room status based on the selected date.
 * In a real application, you might fetch this data from an API.
 */
const getRoomStatusForDate = (room: any, date: Date | undefined) => {
  if (!date) return room.status;
  const day = date.getDate();
  if (day % 3 === 0) return "booked";
  if (day % 2 === 0) return "pending";
  return "available";
};

const Dashboard = () => {
  const [selectedProperty, setSelectedProperty] = useState(properties[0]);
  const [isDropdownOpen, setIsDropdownOpen] = useState(false);
  const [selectedRoom, setSelectedRoom] = useState<any>(null);
  const [isModalOpen, setIsModalOpen] = useState(false);
  const [dateRange, setDateRange] = useState<DateRange | undefined>();
  const [selectedDate, setSelectedDate] = useState<Date | undefined>();
  const [date, setDate] = React.useState<DateRange | undefined>({
    from: new Date(2022, 0, 20),
    to: addDays(new Date(2022, 0, 20), 20),
  });
  const [monthlyRevenue, setMonthlyRevenue] = useState(0);
  const [bookingCount, setBookingCount] = useState(0);

  const [isSuccessModalOpen, setIsSuccessModalOpen] = useState(false);

  const {
    register,
    handleSubmit,
    setValue,
    formState: { errors },
  } = useForm();

  const handlePropertySelect = (property: {
    id: number;
    name: string;
    image: string;
  }) => {
    setSelectedProperty(property);
    setIsDropdownOpen(false);
  };

  const handleDateChange = (date: DateRange | undefined) => {
    setDateRange(date);
    setValue("checkInDate", date?.from || "");
    setValue("checkOutDate", date?.to || "");
  };
  const onSubmit = (data: any) => {
    console.log("Form Data:", data);
    setIsModalOpen(false);
  };
  const handleRoomClick = (room: any) => {
    setSelectedRoom(room);
    setIsModalOpen(true);
  };

  const handleBookingSubmit = (data: any) => {
    console.log("Form Data:", data);
    setIsModalOpen(false);
    setIsSuccessModalOpen(true);

    setMonthlyRevenue((prevRevenue) => prevRevenue + 100); // Assuming $100 per booking
    setBookingCount((prevCount) => prevCount + 1);
  };

  return (
    <div className="container mx-auto px-4  space-y-4">
      <div className="container mx-auto px-4 py-2">
        <div className="relative">
          <button
            className="flex items-center space-x-3 bg-white border rounded-lg px-4 py-2 w-full md:w-72 focus:outline-none focus:ring-2 focus:ring-blue-500"
            onClick={() => setIsDropdownOpen(!isDropdownOpen)}
          >
            <img
              src={selectedProperty.image}
              alt={selectedProperty.name}
              className="w-8 h-8 rounded-full object-cover"
            />
            <span className="flex-1 text-left">{selectedProperty.name}</span>
          </button>
          <div
            className={`absolute top-full left-0 w-full md:w-72 mt-2 bg-white rounded-lg shadow-lg z-10 transition-all duration-200 transform ${
              isDropdownOpen
                ? "opacity-100 translate-y-0"
                : "opacity-0 -translate-y-2 pointer-events-none"
            }`}
          >
            {properties.map((property) => (
              <button
                key={property.id}
                className="flex items-center space-x-3 w-full px-4 py-3 hover:bg-gray-50 transition-colors"
                onClick={() => handlePropertySelect(property)}
              >
                <img
                  src={property.image}
                  alt={property.name}
                  className="w-8 h-8 rounded-full object-cover"
                />
                <span>{property.name}</span>
              </button>
            ))}
          </div>
        </div>
      </div>
      <div className="flex gap-4  ">
        <div className="bg-white p-6 rounded-lg shadow-md w-full">
          <h2 className="text-xl font-semibold mb-2">Monthly Revenue</h2>
          <p className="text-3xl font-bold">
            ${monthlyRevenue.toLocaleString()}
          </p>
        </div>
        <div className="bg-white p-6 rounded-lg shadow-md w-full">
          <h2 className="text-xl font-semibold mb-2">Total Bookings</h2>
          <p className="text-3xl font-bold">{bookingCount}</p>
        </div>
      </div>

      <div className=" mx-auto max-w-4xl">
        <Calendar
          mode="single"
          selected={selectedDate}
          onSelect={setSelectedDate}
          className="rounded-md border"
        />
      </div>
      <div className="grid grid-cols-2 md:grid-cols-3 lg:grid-cols-4 gap-4">
        {rooms.map((room) => {
          const status = getRoomStatusForDate(room, selectedDate);
          const bgClass =
            status === "available"
              ? "bg-green-100 hover:bg-green-200"
              : status === "booked"
              ? "bg-gray-200 hover:bg-gray-300"
              : status === "pending"
              ? "bg-amber-100 hover:bg-amber-200"
              : "";
          return (
            <div
              key={room.id}
              onClick={() => handleRoomClick(room)}
              className={`p-4 rounded-lg cursor-pointer transform transition-transform duration-150 hover:scale-105 ${bgClass}`}
            >
              <div className="flex flex-col space-y-2">
                <span className="text-lg font-semibold">
                  Room {room.number}
                </span>
                <span className="text-sm text-gray-600">{room.type}</span>
                <span className="text-sm capitalize">{status}</span>
              </div>
            </div>
          );
        })}
      </div>

      {/* Modal for booking a room */}
      {isModalOpen && (
        <BookingModal
          isOpen={isModalOpen}
          onClose={() => setIsModalOpen(false)}
          selectedRoom={selectedRoom}
          onSubmit={handleBookingSubmit}
        />
      )}
      <SuccessModal
        isOpen={isSuccessModalOpen}
        onClose={() => setIsSuccessModalOpen(false)}
        message="Your booking has been confirmed!"
      />
    </div>
  );
};

export default Dashboard;
