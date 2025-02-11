"use client";

export type RoomType = {
  id: number;
  number: string;
  status: string;
  type: string;
};

const RoomCard = ({
  room,
  handleRoomClick,
}: {
  room: RoomType;
  handleRoomClick: (room: RoomType) => void;
}) => {
  return (
    <div
      key={room.id}
      onClick={() => handleRoomClick(room)}
      className={`p-4 rounded-lg cursor-pointer transform transition-transform duration-150 hover:scale-105
            ${
              room.status === "available"
                ? "bg-green-100 hover:bg-green-200"
                : ""
            }
            ${room.status === "booked" ? "bg-gray-200 hover:bg-gray-300" : ""}
            ${
              room.status === "pending" ? "bg-amber-100 hover:bg-amber-200" : ""
            }`}
    >
      <div className="flex flex-col space-y-2">
        <span className="text-lg font-semibold">Room {room.number}</span>
        <span className="text-sm text-gray-600">{room.type}</span>
        <span className="text-sm capitalize">{room.status}</span>
      </div>
    </div>
  );
};

export default RoomCard;
