import { FiHome, FiLogOut, FiSettings } from "react-icons/fi";

const Header = () => (
  <header className="bg-white shadow-md py-4 px-6 sticky mb-6 w-full top-0 z-50">
    <div className="max-w-7xl mx-auto flex justify-between items-center">
      <div className="flex items-center space-x-2">
        <FiHome className="h-6 w-6 text-blue-600" />
        <h1 className="text-xl font-semibold text-gray-800">
          Booking App
        </h1>
      </div>
      <div className="flex items-center space-x-4">
        <button className="p-2 hover:bg-gray-100 rounded-full">
          <FiSettings className="h-5 w-5 text-gray-600" />
        </button>
        <button className="p-2 hover:bg-gray-100 rounded-full">
          <FiLogOut className="h-5 w-5 text-gray-600" />
        </button>
      </div>
    </div>
  </header>
);
export default Header;
