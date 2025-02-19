import React, { ChangeEvent } from "react";
import { useForm } from "react-hook-form";
import ImageUploader from "./ImageUpload";
import { RoomData } from "../types";
import { bedConfigs, roomAmenities } from "@/components/consts/consts";

const RoomTab: React.FC = () => {
  const {
    register,
    handleSubmit,
    setValue,
    watch,
    reset,
    formState: { errors },
  } = useForm<RoomData>({
    defaultValues: {
      category: "",
      capacity: "",
      bedConfig: "",
      price: "",
      amenities: [],
      images: [],
    },
  });

  const images = watch("images");

  const onSubmit = (data: RoomData) => {
    console.log("Room data submitted:", data);
  };

  const handleImageUpload = (e: ChangeEvent<HTMLInputElement>) => {
    if (!e.target.files) return;
    const files = Array.from(e.target.files);
    const imageUrls = files.map((file) => URL.createObjectURL(file));
    //todo:upload images to cloud as blob

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
            Room Category *
          </label>
          <input
            type="text"
            {...register("category", { required: "Room category is required" })}
            className={`mt-1 block w-full rounded-md border-gray-300 shadow-sm focus:border-blue-500 focus:ring-blue-500 ${
              errors.category ? "border-red-500" : ""
            }`}
          />
          {errors.category && (
            <p className="mt-1 text-sm text-red-500">
              {errors.category.message}
            </p>
          )}
        </div>

        <div>
          <label className="block text-sm font-medium text-gray-700">
            Bed Configuration
          </label>
          <select
            {...register("bedConfig")}
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
            {...register("price", { required: "Price is required" })}
            className={`mt-1 block w-full rounded-md border-gray-300 shadow-sm focus:border-blue-500 focus:ring-blue-500 ${
              errors.price ? "border-red-500" : ""
            }`}
          />
          {errors.price && (
            <p className="mt-1 text-sm text-red-500">{errors.price.message}</p>
          )}
        </div>

        <div>
          <label className="block text-sm font-medium text-gray-700">
            Maximum Capacity
          </label>
          <input
            type="number"
            {...register("capacity")}
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
        label="Room Images"
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
          Save Room
        </button>
      </div>
    </form>
  );
};

export default RoomTab;
