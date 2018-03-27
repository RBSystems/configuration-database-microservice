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

ALTER TABLE DeviceTypes ADD COnstraint unique_type_name
UNIQUE (typeName);

ALTER TABLE Devices ADD deviceType varchar(255);
UPDATE Devices
JOIN DeviceTypes on Devices.typeID = DeviceTypes.typeID
SET Devices.deviceType = DeviceTypes.typeName;

ALTER TABLE Devices ADD Constraint devices_fk_device_type
FOREIGN KEY (deviceType)
REFERENCES DeviceTypes(typeName)
ON DELETE CASCADE
ON UPDATE CASCADE;

ALTER TABLE Devices DROP FOREIGN KEY Devices_ibfk_5;
ALTER TABLE Devices DROP COLUMN typeID;

ALTER TABLE DeviceClasses ADD Constraint unique_class_name
UNIQUE (name);


ALTER TABLE Devices ADD deviceClass varchar(255);
UPDATE Devices
JOIN DeviceClasses on Devices.classID = DeviceClasses.deviceClassID
SET Devices.deviceClass = DeviceClasses.name;

ALTER TABLE Devices ADD Constraint devices_fk_device_class
FOREIGN KEY (deviceClass)
REFERENCES DeviceClasses(name)
ON DELETE CASCADE
ON UPDATE CASCADE;

ALTER TABLE Devices DROP FOREIGN KEY Devices_ibfk_2;
ALTER TABLE Devices DROP FOREIGN KEY Devices_ibfk_6;
ALTER TABLE Devices DROP COLUMN classID;




ALTER TABLE DeviceRoleDefinition ADD Constraint unique_role_name
UNIQUE (name);

/*
 * There's a duplicate port that needs to get updated, then deleted
 */

UPDATE PortConfiguration 
SET portID = 6
WHERE portID = 38;

UPDATE DeviceTypePorts 
SET portID = 6
WHERE portID = 38;

DELETE FROM PortConfiguration 
WHERE portID = 38;

/*
 * Done with that cleanup
 */ 

ALTER TABLE Ports ADD Constraint unique_port_name
UNIQUE (name);

ALTER TABLE DeviceRole ADD roleName varchar(256);
ALTER TABLE DeviceRole ADD device varchar(767);

UPDATE DeviceRole
JOIN Devices on Devices.deviceID = DeviceRole.deviceID
JOIN DeviceRoleDefinition on DeviceRoleDefinition.deviceRoleDefinitionID = DeviceRole.deviceRoleDefinitionID
SET DeviceRole.roleName = DeviceRoleDefinition.name, 
DeviceRole.device = Devices.FullName;

ALTER TABLE DeviceRole ADD Constraint device_role_fk_device_name
FOREIGN KEY (device)
REFERENCES Devices(fullName)
ON DELETE CASCADE
ON UPDATE CASCADE;

ALTER TABLE DeviceRole ADD Constraint device_role_fk_role_name
FOREIGN KEY (roleName)
REFERENCES DeviceRoleDefinition(name)
ON DELETE CASCADE
ON UPDATE CASCADE;

ALTER TABLE DeviceRole DROP FOREIGN KEY DeviceRole_ibfk_1;
ALTER TABLE DeviceRole DROP FOREIGN KEY DeviceRole_ibfk_2;

ALTER TABLE DeviceRole DROP COLUMN deviceID;
ALTER TABLE DeviceRole DROP COLUMN deviceRoleDefinitionID;

ALTER TABLE PortConfiguration ADD sourceDevice varchar(767);
ALTER TABLE PortConfiguration ADD destinationDevice varchar(767);
ALTER TABLE PortConfiguration ADD hostDevice varchar(767) NOT NULL;
ALTER TABLE PortConfiguration ADD port varchar(767) NOT NULL;

ALTER TABLE PortConfiguration ADD Constraint unique_host_port
UNIQUE (hostDevice, port);


ALTER TABLE PortConfiguration ADD Constraint port_config_fk_port
FOREIGN KEY (port)
REFERENCES Ports(name)
ON DELETE CASCADE
ON UPDATE CASCADE; ALTER TABLE PortConfiguration ADD Constraint port_config_fk_host
FOREIGN KEY (hostDevice)
REFERENCES Devices(fullName)
ON DELETE CASCADE
ON UPDATE CASCADE;

ALTER TABLE PortConfiguration ADD Constraint port_config_fk_dest
FOREIGN KEY (destinationDevice)
REFERENCES Devices(fullName)
ON DELETE CASCADE
ON UPDATE CASCADE;

ALTER TABLE PortConfiguration ADD Constraint port_config_fk_src
FOREIGN KEY (sourceDevice)
REFERENCES Devices(fullName)
ON DELETE CASCADE
ON UPDATE CASCADE;


Update PortConfiguration pc 
JOIN Devices src on src.deviceID = pc.sourceDeviceID
SET 
pc.sourceDevice = src.fullName;

Update PortConfiguration pc 
JOIN Devices dst on dst.deviceID = pc.destinationDeviceID
SET 
pc.destinationDevice = dst.fullName;

Update PortConfiguration pc 
JOIN Devices hst on hst.deviceID = pc.hostDeviceID
SET 
pc.hostDevice = hst.fullName;

Update PortConfiguration pc 
JOIN Ports pts on pts.portID = pc.portID
SET
pc.port = pts.name;


ALTER TABLE PortConfiguration DROP FOREIGN KEY PortConfiguration_ibfk_1;
ALTER TABLE PortConfiguration DROP FOREIGN KEY PortConfiguration_ibfk_2;
ALTER TABLE PortConfiguration DROP FOREIGN KEY PortConfiguration_ibfk_3;

ALTER TABLE PortConfiguration DROP COLUMN portID;
ALTER TABLE PortConfiguration DROP COLUMN hostDeviceID;
ALTER TABLE PortConfiguration DROP COLUMN destinationDeviceID;
ALTER TABLE PortConfiguration DROP COLUMN sourceDeviceID;

/*
 * Now we just gotta do power states
 */ 
