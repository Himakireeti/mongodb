import React from 'react';
import EventsContext from "../context/EventsContext";
import { useContext } from "react";
import {useParams} from "react-router-dom"


function PartyDetail(props) {
    const {partyId} = useParams()
    const { eventsList } = useContext(EventsContext);
    const partyData =eventsList.parties.map((partyData));
    const thisParty = partyData.find(data => data.id === partyId)
    
    return (
        <div
        style={{
          display: 'flex',
          justifyContent: 'center',
          alignItems: 'center',
          height: '90vh'
        }}
      >
        <h1>PartyDetail</h1>
        <h2>Title: {thisParty.title}</h2>
        <h2>Name: {thisParty.host_name}</h2>
      </div>
    
    )
  }
  

export default PartyDetail;