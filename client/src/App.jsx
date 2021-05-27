import { BrowserRouter as Router, Route } from "react-router-dom";
import Dashboard from "./components/Dashboard/Dashboard";
import Login from "./components/Login/Login";

export default function App() {
    return (
        <Router>
            <Route path="/" exact component={Login} />
            <Route path="/dashboard" component={Dashboard} />
        </Router>
    );
}
