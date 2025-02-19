export interface Coordinates {
  lat: string;
  lng: string;
}

export interface Contact {
  name: string;
  phone: string;
  email: string;
}

export interface PropertyFormData {
  propertyName: string;
  address: string;
  coordinates: Coordinates;
  propertyType: string;
  contact: Contact;
  description: string;
  amenities: string[];
  images: string[];
}

export interface RoomData {
  category: string;
  capacity: string;
  bedConfig: string;
  price: string;
  amenities: string[];
  images: string[];
}

export interface Errors {
  [key: string]: string;
}
