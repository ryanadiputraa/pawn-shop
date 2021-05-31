import { BrowserRouter as Router, Route } from "react-router-dom";
import Dashboard from "./components/Dashboard/Dashboard";
import Login from "./components/Login/Login";
import { ThemeProvider } from "@material-ui/core/styles";
import theme from "./theme";

export default function App() {
    return (
        <ThemeProvider theme={theme}>
            <Router>
                <Route path="/" exact component={Login} />
                <Route path="/dashboard" component={Dashboard} />
            </Router>
        </ThemeProvider>
    );
}
