// js/components/signUpPage.js
// import { BASE_URL } from "../constants.js";

export const SignUpPage = () => {
  let name = "";
  let email = "";
  let password = "";
  let code = "";
  let error = "";

  const handleSubmit = async (event) => {
    event.preventDefault();
    // try {
    //   const response = await axios.post(BASE_URL + "/api/register", {
    //     name,
    //     email,
    //     password,
    //     invite_code: code,
    //     is_admin: false.toString(),
    //   });
    //   if (response.status === 200) {
    //     window.location.href = window.location.origin + "/login";
    //   }
    // } catch (error) {
    //   error = error.response.data.error;
    //   renderForm(error);
    // }
  };

  const renderForm = (errorMessage = "") => {
    return `
      <div class="d-flex justify-content-center align-items-center" style="min-height: 100vh; background-color: #1675e0;">
        <form onsubmit="handleSubmit(event)" style="display: flex; flex-direction: column; align-items: center; padding: 20px; border: 1px solid #ddd; border-radius: 4px; margin: 0 auto; width: 400px; background-color: #fff; box-shadow: 0px 0px 10px rgba(0, 0, 0, 0.1);">
          <h1>Sign Up</h1>
          ${
            errorMessage
              ? `<div class="alert alert-danger" role="alert">${errorMessage}</div>`
              : ""
          }
          <div class="mb-3">
            <label for="name" class="form-label">Name:</label>
            <input type="text" id="name" name="name" class="form-control" style="width: 300px" value="${name}" oninput="name = this.value" required />
          </div>
          <div class="mb-3">
            <label for="email" class="form-label">Email:</label>
            <input type="email" id="email" name="email" class="form-control" style="width: 300px" value="${email}" oninput="email = this.value" required />
          </div>
          <div class="mb-3">
            <label for="password" class="form-label">Password:</label>
            <input type="password" id="password" name="password" class="form-control" style="width: 300px" value="${password}" oninput="password = this.value" required />
          </div>
          <div class="mb-3">
            <label for="code" class="form-label">Invite code:</label>
            <input type="text" id="code" name="code" class="form-control" style="width: 300px" value="${code}" oninput="code = this.value" required />
          </div>
          <button type="submit" class="btn btn-primary" style="width: 300px">Sign Up</button>
          <p style="margin-top: 10px;">Already have an account? <a href="/login">Login</a></p>
        </form>
      </div>
    `;
  };

  return renderForm(error);
};
