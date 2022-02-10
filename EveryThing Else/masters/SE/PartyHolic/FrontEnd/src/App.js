import React, { useState, useEffect, useContext } from "react";
import "./App.css";
import Navbar from "./components/Navbar";
import { BrowserRouter as Router, Route, Routes } from "react-router-dom";
import Home from "./pages/home";
import About from "./pages/about";
import Contact from "./pages/contact";
import SignUp from "./pages/signup";
import SignIn from "./pages/signin";
import PartyList from "./pages/PartyList";
import { EventsProvider } from "./context/EventsContext";
import EventsContext from "./context/EventsContext";
import PartyDetail from "./pages/PartyDetail";


function App(props) {
  // const {getCoordinates} = useContext(EventsContext)

  const Submithandler = (e) => {
    // e.preventDefault();
    // console.log(coordinates);
    // axios
    //   .post("https://my-json-server.typicode.com/", coordinates)
    //   .then((resonse) => {
    //     console.log(resonse);
    //   })
    //   .catch((error) => {
    //     console.log(error);
    //   });
  };

  return (
    <EventsProvider>
      <Router>
        <Navbar />
        <Routes>
          <Route path="/" exact element={<Home />} />
          <Route path="/about" element={<About />} />
          <Route path="/partylist" element={<PartyList />} />
          <Route path="/contact-us" element={<Contact />} />
          <Route path="/sign-up" element={<SignUp />} />
          <Route path="/sign-in" element={<SignIn />} />
          <Route path="/partyDetail" element={<PartyDetail/>} />
        </Routes>
      </Router>
    </EventsProvider>
  );
}

export default App;
