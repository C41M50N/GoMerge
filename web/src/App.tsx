import { Component, createResource } from 'solid-js';

const App: Component = () => {

  const fetchData = async () => {
    const res = await fetch("http://localhost:8080/counter");
    return (await res.json()).count;
  }

  const [getData, { refetch }] = createResource(fetchData);

  return (
    <div class="text-center self-center">
      <p class="text-4xl text-green-700 py-20">Hello tailwind!</p>
      <p class="text-3xl text-blue-800">{getData()}</p>
      <button class="text-2xl" onClick={refetch}>Get Data</button>
    </div>
  );
};

export default App;
