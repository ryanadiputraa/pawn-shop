import { BrowserRouter as Router, Route } from "react-router-dom";
import Login from "./components/Login/Login";

export default function App() {
    return (
        <Router>
            <Route path="/" exact component={Login} />
        </Router>
    );
}
