"use client";
import React, { useState, ChangeEvent, FormEvent } from "react";
import { FiUpload, FiX } from "react-icons/fi";

interface Coordinates {
  lat: string;
  lng: string;
}

interface Contact {
  name: string;
  phone: string;
  email: string;
}

interface PropertyFormData {
  propertyName: string;
  address: string;
  coordinates: Coordinates;
  propertyType: string;
  contact: Contact;
  description: string;
  amenities: string[];
  images: string[];
}

interface RoomData {
  category: string;
  capacity: string;
  bedConfig: string;
  price: string;
  amenities: string[];
  images: string[];
}

interface Errors {
  [key: string]: string;
}

type ImageType = "property" | "room";

const PropertyManagement: React.FC = () => {
  const [activeTab, setActiveTab] = useState<"property" | "room">("property");

  const [formData, setFormData] = useState<PropertyFormData>({
    propertyName: "",
    address: "",
    coordinates: { lat: "", lng: "" },
    propertyType: "",
    contact: {
      name: "",
      phone: "",
      email: "",
    },
    description: "",
    amenities: [],
    images: [],
  });

  const [roomData, setRoomData] = useState<RoomData>({
    category: "",
    capacity: "",
    bedConfig: "",
    price: "",
    amenities: [],
    images: [],
  });

  const [errors, setErrors] = useState<Errors>({});

  const propertyTypes = ["Hotel", "Resort", "Boutique", "Villa", "Apartment"];
  const propertyAmenities = [
    "Swimming Pool",
    "Gym",
    "Restaurant",
    "Parking",
    "WiFi",
  ];
  const roomAmenities = [
    "Air Conditioning",
    "TV",
    "Mini Refrigerator",
    "Attached Bathroom",
  ];
  const bedConfigs = ["Single", "Double", "Twin", "King"];

  const handlePropertySubmit = (e: FormEvent<HTMLFormElement>) => {
    e.preventDefault();
    const newErrors: Errors = {};

    if (!formData.propertyName.trim()) {
      newErrors.propertyName = "Property name is required";
    }

    if (!formData.address.trim()) {
      newErrors.address = "Address is required";
    }

    if (Object.keys(newErrors).length === 0) {
      console.log("Property data submitted:", formData);
      // Handle successful submission
      setErrors({});
    } else {
      setErrors(newErrors);
    }
  };

  const handleRoomSubmit = (e: FormEvent<HTMLFormElement>) => {
    e.preventDefault();
    const newErrors: Errors = {};

    if (!roomData.category.trim()) {
      newErrors.category = "Room category is required";
    }

    if (!roomData.price.trim()) {
      newErrors.price = "Price is required";
    }

    if (Object.keys(newErrors).length === 0) {
      console.log("Room data submitted:", roomData);
      // Handle successful submission
      setErrors({});
    } else {
      setErrors(newErrors);
    }
  };

  const handleImageUpload = (
    e: ChangeEvent<HTMLInputElement>,
    type: ImageType
  ) => {
    if (!e.target.files) return;
    const files = Array.from(e.target.files);
    const imageUrls = files.map((file) => URL.createObjectURL(file));

    if (type === "property") {
      setFormData((prev) => ({
        ...prev,
        images: [...prev.images, ...imageUrls],
      }));
    } else {
      setRoomData((prev) => ({
        ...prev,
        images: [...prev.images, ...imageUrls],
      }));
    }
  };

  const removeImage = (index: number, type: ImageType) => {
    if (type === "property") {
      setFormData((prev) => ({
        ...prev,
        images: prev.images.filter((_, i) => i !== index),
      }));
    } else {
      setRoomData((prev) => ({
        ...prev,
        images: prev.images.filter((_, i) => i !== index),
      }));
    }
  };

  return (
      <div className="max-w-7xl mx-auto bg-white rounded-lg shadow-lg">
        <div className="flex border-b">
          <button
            onClick={() => setActiveTab("property")}
            className={`px-6 py-4 ${
              activeTab === "property"
                ? "border-b-2 border-blue-500 text-blue-600"
                : "text-gray-600"
            }`}
          >
            Property Details
          </button>
          <button
            onClick={() => setActiveTab("room")}
            className={`px-6 py-4 ${
              activeTab === "room"
                ? "border-b-2 border-blue-500 text-blue-600"
                : "text-gray-600"
            }`}
          >
            Room Management
          </button>
        </div>

        <div className="p-6">
          {activeTab === "property" ? (
            <form onSubmit={handlePropertySubmit} className="space-y-6">
              <div className="grid grid-cols-1 md:grid-cols-2 gap-6">
                <div>
                  <label className="block text-sm font-medium text-gray-700">
                    Property Name *
                  </label>
                  <input
                    type="text"
                    value={formData.propertyName}
                    onChange={(e) =>
                      setFormData({ ...formData, propertyName: e.target.value })
                    }
                    className={`mt-1 block w-full rounded-md border-gray-300 shadow-sm focus:border-blue-500 focus:ring-blue-500 ${
                      errors.propertyName ? "border-red-500" : ""
                    }`}
                  />
                  {errors.propertyName && (
                    <p className="mt-1 text-sm text-red-500">
                      {errors.propertyName}
                    </p>
                  )}
                </div>

                <div>
                  <label className="block text-sm font-medium text-gray-700">
                    Property Type
                  </label>
                  <select
                    value={formData.propertyType}
                    onChange={(e) =>
                      setFormData({ ...formData, propertyType: e.target.value })
                    }
                    className="mt-1 block w-full rounded-md border-gray-300 shadow-sm focus:border-blue-500 focus:ring-blue-500"
                  >
                    <option value="">Select Type</option>
                    {propertyTypes.map((type) => (
                      <option key={type} value={type}>
                        {type}
                      </option>
                    ))}
                  </select>
                </div>
              </div>

              <div>
                <label className="block text-sm font-medium text-gray-700">
                  Property Amenities
                </label>
                <div className="mt-2 grid grid-cols-2 md:grid-cols-3 gap-4">
                  {propertyAmenities.map((amenity) => (
                    <label key={amenity} className="inline-flex items-center">
                      <input
                        type="checkbox"
                        checked={formData.amenities.includes(amenity)}
                        onChange={(e) => {
                          const newAmenities = e.target.checked
                            ? [...formData.amenities, amenity]
                            : formData.amenities.filter((a) => a !== amenity);
                          setFormData({ ...formData, amenities: newAmenities });
                        }}
                        className="rounded border-gray-300 text-blue-600 focus:ring-blue-500"
                      />
                      <span className="ml-2 text-sm text-gray-600">
                        {amenity}
                      </span>
                    </label>
                  ))}
                </div>
              </div>

              <div>
                <label className="block text-sm font-medium text-gray-700">
                  Property Images
                </label>
                <div className="mt-2 flex justify-center px-6 pt-5 pb-6 border-2 border-gray-300 border-dashed rounded-md">
                  <div className="space-y-1 text-center">
                    <FiUpload className="mx-auto h-12 w-12 text-gray-400" />
                    <div className="flex text-sm text-gray-600">
                      <label className="relative cursor-pointer bg-white rounded-md font-medium text-blue-600 hover:text-blue-500 focus-within:outline-none focus-within:ring-2 focus-within:ring-offset-2 focus-within:ring-blue-500">
                        <span>Upload files</span>
                        <input
                          type="file"
                          multiple
                          onChange={(e) => handleImageUpload(e, "property")}
                          className="sr-only"
                          accept="image/*"
                        />
                      </label>
                    </div>
                  </div>
                </div>

                <div className="mt-4 grid grid-cols-2 md:grid-cols-4 gap-4">
                  {formData.images.map((image, index) => (
                    <div key={index} className="relative">
                      <img
                        src={image}
                        alt={`Property ${index + 1}`}
                        className="h-24 w-full object-cover rounded-lg"
                      />
                      <button
                        type="button"
                        onClick={() => removeImage(index, "property")}
                        className="absolute top-0 right-0 -mt-2 -mr-2 bg-red-500 text-white rounded-full p-1"
                      >
                        <FiX className="h-4 w-4" />
                      </button>
                    </div>
                  ))}
                </div>
              </div>

              <div className="flex justify-end space-x-4">
                <button
                  type="button"
                  onClick={() =>
                    setFormData({
                      propertyName: "",
                      address: "",
                      coordinates: { lat: "", lng: "" },
                      propertyType: "",
                      contact: { name: "", phone: "", email: "" },
                      description: "",
                      amenities: [],
                      images: [],
                    })
                  }
                  className="px-4 py-2 border border-gray-300 rounded-md text-gray-700 hover:bg-gray-50"
                >
                  Reset
                </button>
                <button
                  type="submit"
                  className="px-4 py-2 bg-blue-600 text-white rounded-md hover:bg-blue-700"
                >
                  Save Property
                </button>
              </div>
            </form>
          ) : (
            <form onSubmit={handleRoomSubmit} className="space-y-6">
              <div className="grid grid-cols-1 md:grid-cols-2 gap-6">
                <div>
                  <label className="block text-sm font-medium text-gray-700">
                    Room Category *
                  </label>
                  <input
                    type="text"
                    value={roomData.category}
                    onChange={(e) =>
                      setRoomData({ ...roomData, category: e.target.value })
                    }
                    className={`mt-1 block w-full rounded-md border-gray-300 shadow-sm focus:border-blue-500 focus:ring-blue-500 ${
                      errors.category ? "border-red-500" : ""
                    }`}
                  />
                  {errors.category && (
                    <p className="mt-1 text-sm text-red-500">
                      {errors.category}
                    </p>
                  )}
                </div>

                <div>
                  <label className="block text-sm font-medium text-gray-700">
                    Bed Configuration
                  </label>
                  <select
                    value={roomData.bedConfig}
                    onChange={(e) =>
                      setRoomData({ ...roomData, bedConfig: e.target.value })
                    }
                    className="mt-1 block w-full rounded-md border-gray-300 shadow-sm focus:border-blue-500 focus:ring-blue-500"
                  >
                    <option value="">Select Configuration</option>
                    {bedConfigs.map((config) => (
                      <option key={config} value={config}>
                        {config}
                      </option>
                    ))}
                  </select>
                </div>

                <div>
                  <label className="block text-sm font-medium text-gray-700">
                    Price per Night *
                  </label>
                  <input
                    type="number"
                    value={roomData.price}
                    onChange={(e) =>
                      setRoomData({ ...roomData, price: e.target.value })
                    }
                    className={`mt-1 block w-full rounded-md border-gray-300 shadow-sm focus:border-blue-500 focus:ring-blue-500 ${
                      errors.price ? "border-red-500" : ""
                    }`}
                  />
                  {errors.price && (
                    <p className="mt-1 text-sm text-red-500">{errors.price}</p>
                  )}
                </div>

                <div>
                  <label className="block text-sm font-medium text-gray-700">
                    Maximum Capacity
                  </label>
                  <input
                    type="number"
                    value={roomData.capacity}
                    onChange={(e) =>
                      setRoomData({ ...roomData, capacity: e.target.value })
                    }
                    className="mt-1 block w-full rounded-md border-gray-300 shadow-sm focus:border-blue-500 focus:ring-blue-500"
                  />
                </div>
              </div>

              <div>
                <label className="block text-sm font-medium text-gray-700">
                  Room Amenities
                </label>
                <div className="mt-2 grid grid-cols-2 md:grid-cols-3 gap-4">
                  {roomAmenities.map((amenity) => (
                    <label key={amenity} className="inline-flex items-center">
                      <input
                        type="checkbox"
                        checked={roomData.amenities.includes(amenity)}
                        onChange={(e) => {
                          const newAmenities = e.target.checked
                            ? [...roomData.amenities, amenity]
                            : roomData.amenities.filter((a) => a !== amenity);
                          setRoomData({ ...roomData, amenities: newAmenities });
                        }}
                        className="rounded border-gray-300 text-blue-600 focus:ring-blue-500"
                      />
                      <span className="ml-2 text-sm text-gray-600">
                        {amenity}
                      </span>
                    </label>
                  ))}
                </div>
              </div>

              <div>
                <label className="block text-sm font-medium text-gray-700">
                  Room Images
                </label>
                <div className="mt-2 flex justify-center px-6 pt-5 pb-6 border-2 border-gray-300 border-dashed rounded-md">
                  <div className="space-y-1 text-center">
                    <FiUpload className="mx-auto h-12 w-12 text-gray-400" />
                    <div className="flex text-sm text-gray-600">
                      <label className="relative cursor-pointer bg-white rounded-md font-medium text-blue-600 hover:text-blue-500 focus-within:outline-none focus-within:ring-2 focus-within:ring-offset-2 focus-within:ring-blue-500">
                        <span>Upload files</span>
                        <input
                          type="file"
                          multiple
                          onChange={(e) => handleImageUpload(e, "room")}
                          className="sr-only"
                          accept="image/*"
                        />
                      </label>
                    </div>
                  </div>
                </div>

                <div className="mt-4 grid grid-cols-2 md:grid-cols-4 gap-4">
                  {roomData.images.map((image, index) => (
                    <div key={index} className="relative">
                      <img
                        src={image}
                        alt={`Room ${index + 1}`}
                        className="h-24 w-full object-cover rounded-lg"
                      />
                      <button
                        type="button"
                        onClick={() => removeImage(index, "room")}
                        className="absolute top-0 right-0 -mt-2 -mr-2 bg-red-500 text-white rounded-full p-1"
                      >
                        <FiX className="h-4 w-4" />
                      </button>
                    </div>
                  ))}
                </div>
              </div>

              <div className="flex justify-end space-x-4">
                <button
                  type="button"
                  onClick={() =>
                    setRoomData({
                      category: "",
                      capacity: "",
                      bedConfig: "",
                      price: "",
                      amenities: [],
                      images: [],
                    })
                  }
                  className="px-4 py-2 border border-gray-300 rounded-md text-gray-700 hover:bg-gray-50"
                >
                  Reset
                </button>
                <button
                  type="submit"
                  className="px-4 py-2 bg-blue-600 text-white rounded-md hover:bg-blue-700"
                >
                  Save Room
                </button>
              </div>
            </form>
          )}
        </div>
    </div>
  );
};

export default PropertyManagement;
