import React, { ChangeEvent } from "react";
import { FiUpload, FiX } from "react-icons/fi";

interface ImageUploaderProps {
  images: string[];
  onUpload: (e: ChangeEvent<HTMLInputElement>) => void;
  onRemove: (index: number) => void;
  label: string;
}

const ImageUploader: React.FC<ImageUploaderProps> = ({
  images,
  onUpload,
  onRemove,
  label,
}) => (
  <div>
    <label className="block text-sm font-medium text-gray-700">{label}</label>
    <label className="relative cursor-pointer bg-white rounded-md font-medium hover:border-accent-foreground focus:border-none">
      <div className="mt-2 flex justify-center px-6 pt-5 pb-6 border-2 border-gray-300 border-dashed rounded-md">
        <div className="space-y-1 text-center">
          <FiUpload className="mx-auto h-12 w-12 text-gray-400" />
          <div className="flex text-sm text-gray-600">
            <span>Upload files</span>
            <input
              type="file"
              multiple
              onChange={onUpload}
              className="sr-only"
              accept="image/*"
            ></input>
          </div>
        </div>
      </div>
    </label>
    <div className="mt-4 grid grid-cols-2 md:grid-cols-4 gap-4">
      {images.map((image, index) => (
        <div key={index} className="relative">
          <img
            src={image}
            alt={`Image ${index + 1}`}
            className="h-24 w-full object-cover rounded-lg"
          />
          <button
            type="button"
            onClick={() => onRemove(index)}
            className="absolute top-0 right-0 -mt-2 -mr-2 bg-red-500 text-white rounded-full p-1"
          >
            <FiX className="h-4 w-4" />
          </button>
        </div>
      ))}
    </div>
  </div>
);

export default ImageUploader;
