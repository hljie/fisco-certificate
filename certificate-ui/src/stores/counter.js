import { create } from 'zustand'



const useCounterStore = create()((set) => ({
  counter: 0,
  increase: (by) => set((state) => ({ counter: state.counter + by })),
}))

export default useCounterStore
