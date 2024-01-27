import { BrowserRouter, Routes, Route } from "react-router-dom";
import { Home } from "../pages/home";

const App = () => {
  return (
    <BrowserRouter>
      <div className="m-8">
        <Routes>
          <Route path="/" element={<Home />} />
        </Routes>
      </div>
    </BrowserRouter>
  );
};

export default App;
