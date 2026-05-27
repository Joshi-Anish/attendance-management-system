import Header from "./components/Header";
import Footer from "./components/Footer";
import Dashboard from "./pages/Dashboard";
import Students from "./pages/Students";
import "./App.css";

function App() {
  return (
    <>
      <Header />
      <Dashboard />
      <Students />
      <Footer />
    </>
  );
}

export default App;