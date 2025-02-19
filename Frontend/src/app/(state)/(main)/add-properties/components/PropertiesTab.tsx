import React, { ChangeEvent } from "react";
import { useForm } from "react-hook-form";
import ImageUploader from "./ImageUpload";
import { PropertyFormData } from "../types";
import { propertyAmenities, propertyTypes } from "@/components/consts/consts";

const PropertiesTab: React.FC = () => {
  const {
    register,
    handleSubmit,
    setValue,
    watch,
    reset,
    formState: { errors },
  } = useForm<PropertyFormData>({
    defaultValues: {
      propertyName: "",
      address: "",
      coordinates: { lat: "", lng: "" },
      propertyType: "",
      contact: { name: "", phone: "", email: "" },
      description: "",
      amenities: [],
      images: [],
    },
  });

  const images = watch("images");

  const onSubmit = (data: PropertyFormData) => {
    console.log("Property data submitted:", data);
  };

  const handleImageUpload = (e: ChangeEvent<HTMLInputElement>) => {
    if (!e.target.files) return;
    const files = Array.from(e.target.files);
    const imageUrls = files.map((file) => URL.createObjectURL(file));
    //todo:upload images to cloud as blob
    // Update images in the form state
    const currentImages = watch("images") || [];
    setValue("images", [...currentImages, ...imageUrls]);
  };

  const removeImage = (index: number) => {
    const currentImages = watch("images") || [];
    const updatedImages = currentImages.filter((_, i) => i !== index);
    setValue("images", updatedImages);
  };

  const resetForm = () => {
    reset();
  };

  return (
    <form onSubmit={handleSubmit(onSubmit)} className="space-y-6">
      <div className="grid grid-cols-1 md:grid-cols-2 gap-6">
        <div>
          <label className="block text-sm font-medium text-gray-700">
            Property Name *
          </label>
          <input
            type="text"
            {...register("propertyName", {
              required: "Property name is required",
            })}
            className={`p-2 mt-1 block w-full rounded-md border-gray-300 shadow-sm focus:border-blue-500 focus:ring-blue-500 ${
              errors.propertyName ? "border-red-500" : ""
            }`}
          />
          {errors.propertyName && (
            <p className="mt-1 text-sm text-red-500">
              {errors.propertyName.message}
            </p>
          )}
        </div>

        <div>
          <label className="block text-sm font-medium text-gray-700">
            Address *
          </label>
          <input
            type="text"
            {...register("address", { required: "Address is required" })}
            className={`p-2 mt-1 block w-full rounded-md border-gray-300 shadow-sm focus:border-blue-500 focus:ring-blue-500 ${
              errors.address ? "border-red-500" : ""
            }`}
          />
          {errors.address && (
            <p className="mt-1 text-sm text-red-500">
              {errors.address.message}
            </p>
          )}
        </div>
      </div>

      <div>
        <label className="block text-sm font-medium text-gray-700">
          Property Type
        </label>
        <select
          {...register("propertyType")}
          className="p-2 mt-1 block w-full rounded-md border-gray-300 shadow-sm focus:border-blue-500 focus:ring-blue-500"
        >
          <option value="">Select Type</option>
          {propertyTypes.map((type) => (
            <option key={type} value={type}>
              {type}
            </option>
          ))}
        </select>
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
                value={amenity}
                {...register("amenities")}
                className="rounded border-gray-300 text-blue-600 focus:ring-blue-500"
              />
              <span className="ml-2 text-sm text-gray-600">{amenity}</span>
            </label>
          ))}
        </div>
      </div>

      <ImageUploader
        images={images || []}
        onUpload={handleImageUpload}
        onRemove={removeImage}
        label="Property Images"
      />

      <div className="flex justify-end space-x-4">
        <button
          type="button"
          onClick={resetForm}
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
  );
};

export default PropertiesTab;
