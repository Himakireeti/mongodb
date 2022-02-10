import React from 'react';
import { useEffect, useState } from 'react'
import  ListingParties  from '../components/ListingParties';

function PartyList() {
    
    return (
      <div className='category container align-center'>
              <div>
                  <ListingParties/> 
              </div>
      </div>
    )
  }
  

export default PartyList;
