"use client";
import React, { useState } from "react";
import PropertiesTab from "./components/PropertiesTab";
import RoomTab from "./components/RoomTab";

const PropertyManagement: React.FC = () => {
  const [activeTab, setActiveTab] = useState<"property" | "room">("property");

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
        {activeTab === "property" && <PropertiesTab />}
        {activeTab === "room" && <RoomTab />}
      </div>
    </div>
  );
};

export default PropertyManagement;
