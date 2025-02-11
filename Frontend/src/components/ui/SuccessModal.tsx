import { CheckCircle2 } from "lucide-react";
import React from "react";
import { Button } from "./button";

interface SuccessModalProps {
  isOpen: boolean;
  onClose: () => void;
  message?: string;
}

const SuccessModal: React.FC<SuccessModalProps> = ({
  isOpen,
  onClose,
  message = "Booking successful!",
}) => {
  if (!isOpen) return null;

  return (
    <div className="fixed inset-0 bg-black bg-opacity-50 flex items-center justify-center p-4 z-50 w-full h-full">
      <div className="bg-white rounded-lg p-10 max-w-2xl w-full flex flex-col h-1/2 mx-4">
        <div className="flex justify-end">
          <button
            onClick={onClose}
            className="text-gray-500 hover:text-gray-700 transition-colors"
          >
            ✕
          </button>
        </div>
        <div className="flex flex-col items-center justify-center flex-grow">
          <div className="mb-4">
            <CheckCircle2 className="text-green-500 w-24 h-24" />
          </div>
          <h3 className="text-xl font-semibold mb-2">Success!</h3>
          <p className="text-gray-600 mb-6">{message}</p>
          <Button
            onClick={onClose}
            className="bg-green-500 text-white px-10 py-2 rounded-md hover:bg-green-600 transition-colors"
            size={"xl"}
          >
            OK
          </Button>
        </div>
      </div>
    </div>
  );
};

export default SuccessModal;
