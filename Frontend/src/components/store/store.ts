import { configureStore } from "@reduxjs/toolkit";
import authReducer, { authSliceType } from "../features/authSlice";

export type RootState = {
  auth: authSliceType;
};
export const store = configureStore({
  reducer: {
    auth: authReducer,
  },
});
