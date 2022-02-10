import React from "react";

function LandingStyle() {
  return (
    <>
      <div className="container">
        <img
          src="assets/127045.jpg"
          alt="rent"
          className="object-cover w-screen h-screen blur-xs"
        />
        <div className="cont-text">
          <span className="text1">Welcome</span>
          <span className="text2">into amazing experiences</span>
        </div>
        <button className="cont-btn1  bg-blue-500 hover:bg-blue-700 text-white font-bold py-2 px-4 rounded-full">
          Attend an Event
        </button>
        <button className="cont-btn2 bg-blue-500 hover:bg-blue-700 text-white font-bold py-2 px-4 rounded-full">
          Host an Event
        </button>
      </div>
    </>
  );
}

export default LandingStyle;
