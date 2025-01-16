// js/routes.js
// import { LoginPage } from "./components/loginPage.js";
import { SignUpPage } from "./components/signUpPage.js";
// import { ExpiredPage } from "./components/expiredPage.js";
// import { ForgotPasswordPage } from "./components/forgotPasswordPage.js";
// import { ResetPasswordPage } from "./components/resetPasswordPage.js";
// import { COD, CSGO, LOL, VAL, DOTA, HALO } from "./components/lines.js";
// import {
//   CODSlips,
//   CSGOSlips,
//   HALOSlips,
//   VALSlips,
//   LOLSlips,
// } from "./components/slips.js";
// import {
//   CODPairs,
//   CSGOPairs,
//   HALOPairs,
//   VALPairs,
//   LOLPairs,
// } from "./components/pairs.js";
// import {
//   UserManage,
//   InviteCode,
//   AddUser,
//   EditUser,
//   AddInviteCode,
//   AdminMessageForm,
// } from "./components/admin.js";
// import { Layout } from "./components/layout.js";

// Authentication and Role Permission Checks
const isAuthenticated = () => localStorage.getItem("isLoggedIn") === "true";
const isAdmin = () => localStorage.getItem("isAdmin") === "true";

// Define routes and their corresponding components
export const routes = {
  "/": () => {
    if (isAuthenticated()) {
      return `<div>Welcome to the Dashboard</div>`; // Customize this part as per your layout
    } else {
      return `<div>Please log in to continue.</div>`;
    }
  },
  //   "/login": () => LoginPage(),
  "/register": () => SignUpPage(),
  //   "/user/expired": () => ExpiredPage(),
  //   "/user/forgot-password": () => ForgotPasswordPage(),
  //   "/user/change-password": () => ResetPasswordPage(),
  //   "/cod": () => COD(),
  //   "/cod/slips": () => CODSlips(),
  //   "/cod/pairs": () => CODPairs(),
  //   "/csgo": () => CSGO(),
  //   "/csgo/slips": () => CSGOSlips(),
  //   "/csgo/pairs": () => CSGOPairs(),
  //   "/lol": () => LOL(),
  //   "/lol/slips": () => LOLSlips(),
  //   "/lol/pairs": () => LOLPairs(),
  //   "/val": () => VAL(),
  //   "/val/slips": () => VALSlips(),
  //   "/val/pairs": () => VALPairs(),
  //   "/dota": () => DOTA(),
  //   "/halo": () => HALO(),
  //   "/halo/slips": () => HALOSlips(),
  //   "/halo/pairs": () => HALOPairs(),
  //   "/admin/user-manage": isAdmin()
  //     ? UserManage()
  //     : () => `<h1>Access Denied</h1>`,
  //   "/admin/invite-code": isAdmin()
  //     ? InviteCode()
  //     : () => `<h1>Access Denied</h1>`,
  //   "/admin/user-manage/add-user": isAdmin()
  //     ? AddUser()
  //     : () => `<h1>Access Denied</h1>`,
  //   "/admin/user-manage/edit-user/:userId": isAdmin()
  //     ? EditUser()
  //     : () => `<h1>Access Denied</h1>`,
  //   "/admin/invite-code/add-code": isAdmin()
  //     ? AddInviteCode()
  //     : () => `<h1>Access Denied</h1>`,
  //   "/admin/message": isAdmin()
  //     ? AdminMessageForm()
  //     : () => `<h1>Access Denied</h1>`,
  "*": () => `<h1>404 Page Not Found</h1>`, // Handle 404 routes
};

// Navigation handler
export const navigateTo = (path) => {
  window.history.pushState({}, path, window.location.origin + path);
  renderRoute(path);
};

// Render the current route
export const renderRoute = (route) => {
  const routeFunction = routes[route] || routes["*"]; // Fallback to 404 if the route doesn't exist
  document.getElementById("app").innerHTML = routeFunction();
};
