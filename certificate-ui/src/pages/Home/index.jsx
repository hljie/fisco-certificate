import useCounterStore from '~/stores/counter'

const Home = () => {
  const counter = useCounterStore((state) => state.counter)
  const increase = useCounterStore((state) => state.increase)

  return (
    <div>
      <div>Home Page</div>
      <button onClick={() => increase(1)}> counter: {counter} </button>
    </div>
  )
}

export default Home
