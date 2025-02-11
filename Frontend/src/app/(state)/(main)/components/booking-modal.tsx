"use client";

import { useState, useEffect } from "react";
import { IoMdClose } from "react-icons/io";
import { useForm } from "react-hook-form";
import { Calendar } from "@/components/ui/calendar";
import { DateRange } from "react-day-picker";

interface BookingModalProps {
  isOpen: boolean;
  onClose: () => void;
  selectedRoom: { number: string; type: string } | null;
  onSubmit: (data: any) => void;
}

export const BookingModal = ({
  isOpen,
  onClose,
  selectedRoom,
  onSubmit,
}: BookingModalProps) => {
  const [date, setDate] = useState<DateRange | undefined>();
  const {
    register,
    handleSubmit,
    setValue,
    formState: { errors },
  } = useForm();

  const handleDateChange = (selectedDate: DateRange | undefined) => {
    setDate(selectedDate);
    setValue("checkInDate", selectedDate?.from || "");
    setValue("checkOutDate", selectedDate?.to || "");
  };

  if (!isOpen) return null;

  return (
    <div className="fixed inset-0 bg-black bg-opacity-50 flex items-center justify-center p-4 transition-opacity duration-200">
      <div className="bg-white rounded-lg w-full max-w-4xl p-6 transform transition-transform duration-200">
        <div className="flex justify-between items-center mb-6">
          <h2 className="text-xl font-semibold">
            Book Room {selectedRoom?.number}
          </h2>
          <button
            onClick={onClose}
            className="text-gray-500 hover:text-gray-700"
          >
            <IoMdClose size={24} />
          </button>
        </div>

        <form onSubmit={handleSubmit(onSubmit)} className="space-y-4">
          <div>
            <label className="block text-sm font-medium text-gray-700 mb-1">
              Guest Name
            </label>
            <input
              {...register("guestName", {
                required: "Guest name is required",
              })}
              className="w-full px-3 py-2 border rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
            />
            {typeof errors?.guestName && (
              <span className="text-red-500">
                {String(errors?.guestName?.message)}
              </span>
            )}
          </div>

          <div>
            <label className="block text-sm font-medium text-gray-700 mb-1">
              Check-in & Check-out Date
            </label>
            <Calendar
              initialFocus
              mode="range"
              selected={date}
              onSelect={handleDateChange}
              numberOfMonths={2}
            />
          </div>

          <input type="hidden" {...register("checkInDate")} />
          <input type="hidden" {...register("checkOutDate")} />

          <div className="flex space-x-4">
            <button
              type="submit"
              className="flex-1 bg-blue-500 text-white px-4 py-2 rounded-lg hover:bg-blue-600 transition-colors"
            >
              Confirm Booking
            </button>
            <button
              type="button"
              onClick={onClose}
              className="flex-1 bg-gray-200 text-gray-800 px-4 py-2 rounded-lg hover:bg-gray-300 transition-colors"
            >
              Cancel
            </button>
          </div>
        </form>
      </div>
    </div>
  );
};
