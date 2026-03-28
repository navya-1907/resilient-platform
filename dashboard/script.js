async function fetchData() {
  try {
    const res = await fetch("http://localhost:8080/containers");
    const text = await res.text();
    document.getElementById("output").innerText = text;
  } catch (err) {
    document.getElementById("output").innerText = "Backend not reachable";
  }
}

setInterval(fetchData, 2000);