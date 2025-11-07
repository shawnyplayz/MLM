import { createSlice } from '@reduxjs/toolkit';

const initialState = {
  distributors: [],
  currentDistributor: null,
  loading: false,
  error: null,
};

const distributorSlice = createSlice({
  name: 'distributors',
  initialState,
  reducers: {
    setDistributors: (state, action) => {
      state.distributors = action.payload;
    },
    setCurrentDistributor: (state, action) => {
      state.currentDistributor = action.payload;
    },
    clearError: (state) => {
      state.error = null;
    },
  },
});

export const { setDistributors, setCurrentDistributor, clearError } = distributorSlice.actions;
export default distributorSlice.reducer;
