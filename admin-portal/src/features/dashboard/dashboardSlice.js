import { createSlice } from '@reduxjs/toolkit';

const initialState = {
  stats: null,
  loading: false,
  error: null,
};

const dashboardSlice = createSlice({
  name: 'dashboard',
  initialState,
  reducers: {
    setStats: (state, action) => {
      state.stats = action.payload;
    },
    clearError: (state) => {
      state.error = null;
    },
  },
});

export const { setStats, clearError } = dashboardSlice.actions;
export default dashboardSlice.reducer;
