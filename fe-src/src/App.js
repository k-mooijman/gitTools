import './App.css';
import { BrowserRouter, Routes, Route } from "react-router-dom";
import Layout from "./pages/Layout";
import Home from "./pages/Home";
import Blogs from "./pages/Blogs";
import Contact from "./pages/Contact";
import NoPage from "./pages/NoPage";
import Test01 from "./pages/test/test01";
import Test02 from "./pages/test/Test02";
import useTest from "./work/test";


function App() {
    useTest();
    return (
        <BrowserRouter>
            <Routes>
                <Route path="/" element={<Layout />}>
                    <Route index element={<Home />} />
                    <Route path="blogs" element={<Blogs />} />
                    <Route path="contact" element={<Contact />} />
                    <Route path="*" element={<NoPage />} />
                </Route>
                <Route path="Test01" element={<Test01 />} />
                <Route path="Test02" element={<Test02 />} />

            </Routes>
        </BrowserRouter>
    );
}

export default App;
