import {useState,useEffect} from 'react';


function GeoLocation(props) {
    const [lat, setLat] = useState(null);
    const [lng, setLng] = useState(null);
    const [status, setStatus] = useState(null);

    function getLocation() {
        if (!navigator.geolocation) {
            setStatus('Geolocation is not supported by your browser');
          } 
        else {
            setStatus('Locating...');
            navigator.geolocation.getCurrentPosition((position) => {
              setStatus(null);
              setLat(position.coords.latitude);
              setLng(position.coords.longitude);
            }, () => {
              setStatus('Unable to retrieve your location');
            });
          }
        }

  return (
    <div>
        <button onClick={getLocation}>Get Location</button>
        <p>{status}</p>
        {lat && <p>Latitude: {lat}</p>}
        {lng && <p>Longitude: {lng}</p>}
    </div>
  );
 }
    
 export default GeoLocation;