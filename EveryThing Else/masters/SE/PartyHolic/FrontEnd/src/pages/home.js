import React from "react";
import { Link } from "react-router-dom";
import { useState } from "react";
import ZipCode from "../components/Zipcode";
import Geolocation from "../components/Geolocation";
import Navbar from "../components/Navbar";
import "../index.css";
import LandingStyle from "../components/LandingStyle";

function Home(props) {
  const [zipCodeOpen, setZipCodeOpen] = useState(false);
  const [geoLocOpen, setGeoLocOpen] = useState(true);
  const [lat, setLat] = useState(null);
  const [lng, setLng] = useState(null);
  const [status, setStatus] = useState(null);

  function openZipCodeModal() {
    setZipCodeOpen(true);
  }

  function closeZipCodeModal() {
    setZipCodeOpen(false);
    alert("Entered zip code is xx");
  }

  function openGeoLocModal() {
    setGeoLocOpen(true);
  }

  function closeGeoLocModal() {
    setGeoLocOpen(false);
    alert("Received geo location");
  }

  return (
    <div>
      <div>
        <main>
          <LandingStyle/>
          <button onClick={openZipCodeModal}>Press for zipcode</button>
          {zipCodeOpen && <ZipCode onEnter={closeZipCodeModal} />}
          {geoLocOpen && <Geolocation onEnter={closeGeoLocModal} />}
        </main>
      </div>
    </div>
  );
}

export default Home;
