import { useEffect, useState } from "react";
import { Line } from "react-chartjs-2";

export default function Home() {
  const [data, setData] = useState([]);

  useEffect(() => {
    const fetchReadings = async () => {
      const res = await fetch("http://localhost:3001/api/readings");
      const json = await res.json();
      setData(json);
    };

    const interval = setInterval(fetchReadings, 5000);
    return () => clearInterval(interval);
  }, []);

  const chartData = {
    labels: data.map((d) => new Date(d.timestamp * 1000).toLocaleTimeString()),
    datasets: [
      {
        label: "Temperature",
        data: data.map((d) => d.temperature),
        borderColor: "red",
        fill: false,
      },
      {
        label: "Humidity",
        data: data.map((d) => d.humidity),
        borderColor: "blue",
        fill: false,
      },
    ],
  };

  return <Line data={chartData} />;
}
