async function fetchData() {
  try {
    const response = await fetch("http://127.0.0.1:8080/containers");
    
    if (!response.ok) {
      throw new Error("Backend error");
    }

    const data = await response.text();
    document.getElementById("output").innerText = data;

  } catch (error) {
    document.getElementById("output").innerText = "Backend not reachable";
    console.error(error);
  }
}

fetchData();
setInterval(fetchData, 2000);