import { Navigate, useOutlet } from "react-router-dom";
import { useAuth } from "../hooks/useAuth";

export const HomeLayout = () => {
    const { user } = useAuth();
    const outlet = useOutlet();

    if (user) {
        return <Navigate to="/cod" replace />;
    }

    return (
        <div>
            {/* <AppBar
                pages={[
                    { label: "Home", path: "/" },
                    { label: "Login", path: "/login" },
                ]}
            /> */}
            {outlet}
        </div>
    );
};
