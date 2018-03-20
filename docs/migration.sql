/* add buildingName */
ALTER TABLE Rooms ADD buildingName varchar(256);

/* Make cross reference in rooms table */
UPDATE Rooms
JOIN Buildings
ON Rooms.buildingID = Buildings.buildingID
SET Rooms.buildingName = Buildings.shortName;

ALTER TABLE Devices ADD buildingName varchar(256);

UPDATE Devices
JOIN Buildings
ON Devices.buildingID = Buildings.buildingID
SET Devices.buildingName = Buildings.shortName;

ALTER TABLE Rooms DROP FOREIGN KEY Rooms_ibfk_1;
ALTER TABLE Devices DROP FOREIGN KEY Devices_ibfk_1;

ALTER TABLE Rooms DROP COLUMN buildingID;                                                                                                                                                                                                                                   
ALTER TABLE Devices DROP COLUMN buildingID;
/*
 * Now we can edit the buildings table
  */
  ALTER TABLE Buildings DROP PRIMARY KEY,
change buildingID buildingID int,
ADD PRIMARY KEY (shortName);

ALTER TABLE Buildings DROP COLUMN buildingID;
ALTER TABLE Rooms ADD Constraint rooms_fk_building_name
FOREIGN KEY (buildingName)
REFERENCES Buildings(shortName)
ON DELETE CASCADE
ON UPDATE CASCADE;

ALTER TABLE Devices ADD Constraint devices_fk_building_name
FOREIGN KEY (buildingName)
REFERENCES Buildings(shortName)
ON DELETE CASCADE
ON UPDATE CASCADE;

/*
 *  Now we need to be able to change the room name such that the primary key is the BUILDING-ROOM
 */
ALTER TABLE Rooms ADD RoomIdentifier varchar(512);

UPDATE Rooms SET RoomIdentifier = CONCAT(buildingName, "-", name);

/*
 *
 */
ALTER Table Devices ADD roomName varchar(256);
UPDATE Devices
JOIN Rooms
ON Devices.roomID = Rooms.roomID
SET Devices.roomName = Rooms.RoomIdentifier;

ALTER TABLE Devices DROP FOREIGN KEY Devices_ibfk_4;
ALTER TABLE Devices DROP COLUMN roomID;

ALTER TABLE Rooms DROP PRIMARY KEY,
change roomID roomID int;

ALTER TABLE Rooms DROP COLUMN roomID;
ALTER TABLE Rooms CHANGE `roomIdentifier` `roomID` varchar(512);

ALTER TABLE Rooms ADD PRIMARY KEY (roomID);

ALTER TABLE Devices ADD Constraint devices_fk_room_name
FOREIGN KEY (roomName)
REFERENCES Rooms(roomID)
ON DELETE CASCADE
ON UPDATE CASCADE;

/*
 * Triggers to autogenerate the ID based on the room name and building
 */
CREATE TRIGGER insert_trigger
BEFORE INSERT ON Rooms
FOR EACH ROW
SET new.roomID = CONCAT(buildingName, "-", name);

CREATE TRIGGER update_trigger
BEFORE UPDATE ON Rooms
FOR EACH ROW
SET new.roomID = CONCAT(buildingName, "-", name);


/*
* Now we're gonna add the device fullName - the trigger, and then the unique constraint 
*/
ALTER TABLE Devices ADD fullName varchar(767);
UPDATE Devices SET fullName = CONCAT(roomName, "-", name);

ALTER TABLE Devices ADD Constraint unique_devices_name
UNIQUE (fullName);

