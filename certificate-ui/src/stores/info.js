import { create } from 'zustand'



const useInfoStore = create()((set) => ({
  info: {
    nickname: '呢称',
  },
  setInfo: (info) => set({ info }),
}))

export default useInfoStore
