import { useEffect, useState } from 'react';
import { Line } from 'react-chartjs-2';
import {
  Chart as ChartJS,
  CategoryScale,
  LinearScale,
  PointElement,
  LineElement,
  Legend,
  Tooltip,
} from 'chart.js';

ChartJS.register(CategoryScale, LinearScale, PointElement, LineElement, Legend, Tooltip);

function App() {
  const [readings, setReadings] = useState([]);
  const [metrics, setMetrics] = useState([]);

  useEffect(() => {
    const fetchReadings = async () => {
      const res = await fetch('http://localhost:8080/readings');
      const json = await res.json();
      setReadings(json);
    };

    const fetchMetrics = async () => {
      const res = await fetch('http://localhost:8080/metrics');
      const json = await res.json();
      setMetrics(json);
    };

    const interval = setInterval(() => {
      fetchReadings();
      fetchMetrics();
    }, 3000);

    return () => clearInterval(interval);
  }, []);

  // --- Chart 1: Temperature & Humidity ---
  const telemetryChart = {
    labels: readings.map(r => new Date(r.timestamp * 1000).toLocaleTimeString()),
    datasets: [
      {
        label: 'Temperature (°C)',
        data: readings.map(r => r.temperature),
        borderColor: 'red',
        fill: false,
      },
      {
        label: 'Humidity (%)',
        data: readings.map(r => r.humidity),
        borderColor: 'blue',
        fill: false,
      },
    ],
  };

  // --- Chart 2: Latency of JSON vs Protobuf ---
  const jsonPoints = metrics.filter(m => m.format === 'json');
  const protoPoints = metrics.filter(m => m.format === 'protobuf');

  const labels = jsonPoints.map((_, i) => `#${i + 1}`);

  const latencyChart = {
    labels,
    datasets: [
      {
        label: 'JSON Latency (ms)',
        data: jsonPoints.map(m => m.latency.toFixed(2)),
        borderColor: 'green',
        fill: false,
      },
      {
        label: 'Protobuf Latency (ms)',
        data: protoPoints.map(m => m.latency.toFixed(2)),
        borderColor: 'purple',
        fill: false,
      },
    ],
  };

  return (
    <div style={{ padding: 40 }}>
      <h1>IoT Telemetry Dashboard</h1>

      <h2>Temperature & Humidity</h2>
      <Line data={telemetryChart} />

      <h2 style={{ marginTop: 60 }}>Message Format Latency (ms)</h2>
      <Line data={latencyChart} />
    </div>
  );
}

export default App;
