import { Route, Routes } from "react-router-dom";
import { Sidebar } from "./components/sidebar";
import { ConfigPage } from "./pages/config.page";


export default function App() {

  return (
    <div className="flex h-screen">
      <Sidebar />

      <Routes>
        <Route path="/" element={<ConfigPage />} />
      </Routes>

    </div>
  )
}