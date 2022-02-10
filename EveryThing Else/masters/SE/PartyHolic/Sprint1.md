# Software Engineering Spring 2022

## Project - PartyHolic

## Sprint 1

### Link to the video presentation for sprint 1 - 
https://uflorida-my.sharepoint.com/:v:/g/personal/himakireetikonda_ufl_edu/EQt7EE4c5KtKkxYoQejIzXYBjg4Cof_Wl5CN1lE0-Z2FpQ?e=UJppv0

### Deliverable -
User should be able to see list of parties going on at the moment near their location 

### Sprint planning- 
23 January 2022 
### Sprint start - 
24 January 2022 
### Sprint end - 
3 February 2022 



## What we accomplished (7/7 closed stories in sprint 1)

#### PBI 1- As a DevOps engineer, I want to set up version control
#### PBI 2- As a frontend engineer, I want to setup React Project and file structure
#### PBI 3 - As a partygoer, I want to see list of parties going around me right now
##### Tasks:
- [x] Browser should ask for location permission from the user
- [x] If user denies it, app will ask for zip code
- [x] After this step, user should see the list of parties going on around them
- [x] each party list item will have - title, background picture, host name, distance, no of people attending
### PBI 4 - As a frontend engineer, I want to integrate the party list with backend apis 
##### Tasks:
- [x] save user location in the client side application
- [x] request the list of parties going on in that location - api call 1
- [x] receive the list of parties from front end 
- [x] display it in react components created in PBI 3
- [x] host frontend from backend
### PBI 5 - As a backend engineer, I want to design database schema for party information data
##### Tasks:
- [x] Assuming backend is set up,
- [x] create schema and empty table creation on server init
- [x] create methods for populating new entries for party list
- [x] send party list for given user and given geolocation - for 10 mile radius
### PBI 6 - As a backend engineer, I want ot set up backend structure
##### Tasks:
- [x] create a dummy web app in go and sqlite
### PBI 7 - As a backend developer, I want to create a party list api 
##### Tasks:
- [x] get request from user with user id and geolocation 
- [x] send json response with party id, party title, host, distance, no of attendees, image link and all else mentions in frontend list
- [x] create a folder for image resources







