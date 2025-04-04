const BASE_URL = "http://localhost:3000";
const TESTING_TOKEN =
  "eyJhbGciOiJSUzI1NiIsImNhdCI6ImNsX0I3ZDRQRDExMUFBQSIsImtpZCI6Imluc18yc08zOWpQYUpJemlSU3BXZHpuWWN6eWg1WDgiLCJ0eXAiOiJKV1QifQ.eyJhenAiOiJodHRwOi8vbG9jYWxob3N0OjUxNzMiLCJleHAiOjE3NDM0OTk0MjQsImZ2YSI6WzEsLTFdLCJpYXQiOjE3NDM0OTkzNjQsImlzcyI6Imh0dHBzOi8vY3J1Y2lhbC1qYWd1YXItNDUuY2xlcmsuYWNjb3VudHMuZGV2IiwibmJmIjoxNzQzNDk5MzU0LCJzaWQiOiJzZXNzXzJ2N2J5MjZ0ekFSVDhXT3dPZGdEdmw5UElTUyIsInN1YiI6InVzZXJfMnY3Ynk4dGZ5Vm5HRk5kbmE1NEU2NEJkVGIwIn0.N6Wtm8yfFIVkoX88bQWoYlJELTtQ6aEarbw4zeXx6h37h0_BhJ9xCaHrCV_XLV1SBEgPPjzkMv_rZYsKSzuR-YfrE-RCFpUDeZGnO9AtwY7muVFE99jiMYzfKU7bH0F0wYnxeH1LhO6iD8z6IZvOIXTFMHX2E9Pk8iOJvWBxhlObuGMXhAvmJJDVxPsjmoeRWjv7VRjdRep4ZlGMrG-sQ-cglSkM2awtpI7sqPDXxQwhzpKbXzgUogV0OYlFckFnsM3cD7NmFYcgU6f1cOi_apFMnmJY_NrDIk-m0gPqgDN8tyN1LIVOD_uWcovPcsHJXluaCq_39aXj2RZbMSA_zQ";

async function testPublicRoute() {
  const response = await fetch(`${BASE_URL}/`);
  const data = await response.json();
  console.log("Public Route Response:", data);
}

async function testProtectedRoute() {
  const response = await fetch(`${BASE_URL}/protected`, {
    method: "GET",
    headers: {
      Authorization: `Bearer ${TESTING_TOKEN}`,
    },
  });

  const text = await response.text(); // Read raw response first
  console.log("Raw Response:", text);

  try {
    const data = JSON.parse(text); // Try parsing JSON
    console.log("Protected Route Response:", data);
  } catch (error) {
    console.error("Failed to parse JSON:", error);
  }
}

async function runTests() {
  console.log("Testing Public Route...");
  await testPublicRoute();

  console.log("Testing Protected Route...");
  await testProtectedRoute();
}

runTests().catch(console.error);
