import { configureStore } from '@reduxjs/toolkit';
import authReducer from '../features/auth/authSlice';
import distributorReducer from '../features/distributors/distributorSlice';
import dashboardReducer from '../features/dashboard/dashboardSlice';

export const store = configureStore({
  reducer: {
    auth: authReducer,
    distributors: distributorReducer,
    dashboard: dashboardReducer,
  },
});
