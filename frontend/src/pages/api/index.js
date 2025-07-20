export default async function handler(req, res) {
  const result = await fetch("http://processor-service:8080/readings");
  const json = await result.json();
  res.status(200).json(json);
}
