import { routes } from "./routes.js";

// Initial state (simulating Redux state)
let state = {
  collapsed: false,
  isAuthenticated: localStorage.getItem("isLoggedIn") === "true",
  userId: localStorage.getItem("userId"),
  lines: { lol: [] }, // Example: use a default empty array
  usedList: [],
  pairs: [],
  slips: [],
};

// Simulate Redux selectors
const selectUsed = () => state.usedList;
const selectUserId = () => state.userId;
const selectPairsOfEsport1 = (league) => state.pairs[league] || [];
const selectSlipsOfEsport1 = (league) => state.slips[league] || [];

// Function to get the league from the URL (simulating `getLeague` function)
function getLeague(pathname) {
  const parts = (pathname || "").split("/");
  return parts.length < 2 ? "" : parts[1];
}

// Simulate dispatching actions (simplified version)
function dispatch(action) {
  if (action.type === "fetchLines") {
    state.lines.lol = ["some data"]; // Example
  }
  if (action.type === "fetchpairs") {
    state.pairs = { lol: ["pair1", "pair2"] }; // Example
  }
  if (action.type === "fetchAllUsed") {
    state.usedList = ["usedItem1", "usedItem2"]; // Example
  }
}

// Rendering Components (Simplified)
function renderNavbar() {
  const navbarElement = document.getElementById("navbar");
  navbarElement.innerHTML = state.collapsed
    ? "<div>Collapsed Navbar</div>"
    : "<div>Expanded Navbar</div>";
}

function renderMainHeader() {
  const headerElement = document.getElementById("main-header");
  headerElement.innerHTML = `<div>Main Header with data: ${JSON.stringify(
    state.pairs
  )}</div>`;
}

function renderNotificationBar() {
  const notificationElement = document.getElementById("notification-bar");
  notificationElement.innerHTML = `<div>Notification Bar</div>`;
}

function renderContent() {
  const contentElement = document.getElementById("content");
  const league = getLeague(window.location.pathname);
  const pairs = selectPairsOfEsport1(league);
  const slips = selectSlipsOfEsport1(league);

  contentElement.innerHTML = `<div>Content for league: ${league}</div>
                                <div>Pairs: ${JSON.stringify(pairs)}</div>
                                <div>Slips: ${JSON.stringify(slips)}</div>`;
}

function renderAutoLogout() {
  const autoLogoutElement = document.getElementById("auto-logout");
  autoLogoutElement.innerHTML = "<div>Auto Logout Timer</div>";
}

// Handle the collapsing of the sidebar
function toggleCollapse() {
  state.collapsed = !state.collapsed;
  renderNavbar();
}

// Handle routing & page changes
function renderRoute(route) {
  const routeFunction = routes[route] || routes["*"]; // Fallback to 404 if route doesn't exist
  const appElement = document.getElementById("app");
  appElement.innerHTML = ""; // Clear existing content

  if (!state.isAuthenticated && route !== "/login" && route !== "/register") {
    // Redirect to login page if the user is not authenticated and trying to access a protected route
    navigateTo("/login");
    return;
  }
  if (state.isAuthenticated) {
    appElement.innerHTML += `
            <div id="navbar"></div>
            <div id="main-header"></div>
            <div id="content"></div>
            <div id="notification-bar"></div>
            <div id="auto-logout"></div>
        `;

    renderNavbar();
    renderMainHeader();
    renderContent();
    renderNotificationBar();
    renderAutoLogout();
  }
}

function renderLoginPage() {
  const appElement = document.getElementById("app");
  appElement.innerHTML = renderLoginForm();
}

// Handle initial route rendering
const handleRouteChange = () => {
  const path = window.location.pathname;
  renderRoute(path);
};

// Listen for changes in the URL (for browser navigation)
window.addEventListener("popstate", handleRouteChange);

// Handle navigation function to change the browser's URL and trigger re-render
window.navigateTo = (path) => {
  window.history.pushState({}, path, window.location.origin + path);
  handleRouteChange();
};

// Example navigation in the app (this can be in your HTML)
const setupNavigation = () => {
  document.getElementById("app").innerHTML += `
        <nav>
            <a href="#" onclick="navigateTo('/login'); return false;">Login</a> |
            <a href="#" onclick="navigateTo('/register'); return false;">Register</a> |
            <a href="#" onclick="navigateTo('/admin'); return false;">Admin</a>
        </nav>
    `;
};

// Set up navigation once the page loads
setupNavigation();

// Simulate componentDidMount (like `useEffect`)
function componentDidMount() {
  if (!state.lines.lol.length) {
    dispatch({ type: "fetchLines" });
    dispatch({ type: "fetchpairs" });
  }

  if (!state.usedList.length && state.userId) {
    dispatch({ type: "fetchAllUsed", userId: state.userId });
  }

  handleRouteChange();
}

// Simulate page load
window.onload = function () {
  componentDidMount();
};
