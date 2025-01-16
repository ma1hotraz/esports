// login.js
const BASE_URL = "your_base_url_here"; // Set your BASE_URL for API requests

// Function to handle login submission
// async function handleSubmit(event) {
//   event.preventDefault();

//   const email = document.getElementById("email").value;
//   const password = document.getElementById("password").value;
//   const errorDiv = document.getElementById("error-message");

//   try {
//     const response = await fetch(BASE_URL + "/api/login", {
//       method: "POST",
//       headers: {
//         "Content-Type": "application/json",
//       },
//       body: JSON.stringify({ email, password }),
//     });

//     if (response.ok) {
//       const data = await response.json();

//       // Storing user data in localStorage
//       localStorage.setItem("isLoggedIn", true);
//       localStorage.setItem("isAdmin", data.is_admin);
//       localStorage.setItem("email", data.email);
//       localStorage.setItem("userId", data.userId);
//       localStorage.setItem("user", JSON.stringify(data.user));

//       // Redirect based on user role
//       if (!data.is_admin) {
//         window.location.href = window.location.origin + "/cod";
//       } else {
//         window.location.href = window.location.origin + "/admin/user-manage";
//       }
//     } else {
//       throw new Error("Invalid email or password.");
//     }
//   } catch (error) {
//     console.error("Login failed:", error);
//     errorDiv.textContent = "Invalid email or password."; // Set error message
//     errorDiv.style.display = "block"; // Show error message
//   }
// }

// Function to render the login form
export function renderLoginForm() {
  const appElement = document.getElementById("app");
  appElement.innerHTML = `
        <div style="display: flex; justify-content: center; align-items: center; min-height: 100vh; background-color: #1675e0;">
            <div class="bg-white p-4 rounded shadow" style="width: 400px;">
                <form id="login-form" class="d-flex flex-column align-items-center">
                    <h1 class="mb-4">Login</h1>
                    <div class="mb-4">
                        <label for="email" class="form-label">Email:</label>
                     

                        <input type="email" id="email" name="email" class="form-control" style="width: 350px;" required />
                    </div>
                    <div class="mb-3">
                        <label for="password" class="form-label">Password:</label>
                        <input type="password" id="password" name="password" class="form-control" style="width: 350px;" required />
                    </div>
                    <div id="error-message" class="alert alert-danger" role="alert" style="display: none;"></div>
                    <button type="submit" class="btn btn-primary w-100">Login</button>
                    <p class="mt-3">Don't have an account? <a href="/register">Sign up</a></p>
                    <p class="mt-2 mb-4"><a href="/user/forgot-password">Forgot your password?</a></p>
                </form>
            </div>
        </div>
    `;

  // Add event listener for the form submit
  const form = document.getElementById("login-form");
  form.addEventListener("submit", handleSubmit);
}

// Initialize the app by rendering the login form
window.onload = function () {
  renderLoginForm();
};
