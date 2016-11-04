CREATE TABLE `DeviceTypes` (
  `deviceTypeID` int NOT NULL AUTO_INCREMENT,
  `name` varchar(256),
  `description` text,
  PRIMARY KEY (deviceTypeID)
);

CREATE TABLE `DeviceRoleDefinition` (
  `deviceRoleDefinitionID` int NOT NULL AUTO_INCREMENT,
  `name` varchar(256),
  `description` text,
  PRIMARY KEY (deviceRoleDefinitionID)
);

CREATE TABLE `Commands` (
  `commandID` int NOT NULL AUTO_INCREMENT,
  `name` varchar(256),
  `description` text,
  PRIMARY KEY (commandID)
);

CREATE TABLE `Microservices` (
  `microserviceID` int NOT NULL AUTO_INCREMENT,
  `name` varchar(256),
  `address` text,
  `description` text,
  PRIMARY KEY (microserviceID)
);

CREATE TABLE `Endpoints` (
  `endpointID` int NOT NULL AUTO_INCREMENT,
  `name` varchar(256),
  `path` text,
  `description` text,
  PRIMARY KEY (endpointID)
);

CREATE TABLE `PowerState` (
  `powerStateID` int NOT NULL AUTO_INCREMENT,
  `name` varchar(256),
  `description` text,
  PRIMARY KEY (powerStateID)
);

CREATE TABLE `Buildings` (
  `buildingID` int NOT NULL AUTO_INCREMENT,
  `name` text,
  `shortName` varchar(256),
  `description` text,
  PRIMARY KEY (buildingID)
);

CREATE TABLE `Devices` (
  `deviceID` int NOT NULL AUTO_INCREMENT,
  `name` varchar(256),
  `address` varchar(256),
  `buildingID` int,
  `roomID` int,
  `typeID` int,
  `powerID` int,
  `responding` boolean,
  PRIMARY KEY (deviceID),
  INDEX devbld_ind (buildingID),
  INDEX devrm_ind (roomID),
  INDEX devty_ind (typeID),
  INDEX devpw_ind (powerID),
  FOREIGN KEY (buildingID)
    REFERENCES Buildings(buildingID),
  FOREIGN KEY (typeID)
    REFERENCES DeviceTypes(deviceTypeID),
  FOREIGN KEY (powerID)
    REFERENCES PowerState(powerStateID)
);

CREATE TABLE `Rooms` (
  `roomID` int NOT NULL AUTO_INCREMENT,
  `name` varchar(256),
  `buildingID` int,
  `description` varchar(256),
  `currentVideoOutputID` int,
  `currentAudioOutputID` int,
  `currentVideoInputID` int,
  `currentAudioInputID` int,
  PRIMARY KEY (roomID),
  INDEX rmBld_ind(buildingID),
  INDEX rmCurVOut_ind(currentVideoOutputID),
  INDEX rmCurVIn_ind(currentVideoInputID),
  INDEX rmCurAOut_ind(currentAudioOutputID),
  INDEX rmCurAIn_ind(currentAudioInputID),
  FOREIGN KEY (buildingID)
    REFERENCES Buildings(buildingID),
  FOREIGN KEY (currentVideoOutputID)
    REFERENCES Devices(deviceID),
  FOREIGN KEY (currentVideoInputID)
    REFERENCES Devices(deviceID),
  FOREIGN KEY (currentAudioOutputID)
    REFERENCES Devices(deviceID),
  FOREIGN KEY (currentAudioInputID)
    REFERENCES Devices(deviceID)
);

CREATE TABLE `DeviceRole` (
  `deviceRoleID` int NOT NULL AUTO_INCREMENT,
  `deviceID` int,
  `deviceRoleDefinitionID` int,
  PRIMARY KEY (deviceRoleID),
  INDEX devRolDevID_ind(deviceID),
  INDEX devRolDevRolDef_ind(deviceRoleDefinitionID),
  FOREIGN KEY (deviceID)
    REFERENCES Devices(deviceID),
  FOREIGN KEY (DeviceRoleDefinitionID)
    REFERENCES DeviceRoleDefinition(DeviceRoleDefinitionID)
);

CREATE TABLE `AudioDevice` (
  `audioDeviceID` int NOT NULL AUTO_INCREMENT,
  `deviceID` int,
  `deviceRoleDefinitionID` int,
  `muted` boolean,
  `volume` int,
  PRIMARY KEY (audioDeviceID),
  INDEX audDev_ind(deviceID),
  INDEX audDevRol_ind(deviceRoleDefinitionID),
  FOREIGN KEY (deviceID)
    REFERENCES Devices(deviceID),
  FOREIGN KEY (deviceRoleDefinitionID)
    REFERENCES DeviceRoleDefinition(deviceRoleDefinitionID)
);

CREATE TABLE `Displays` (
  `displayID` int NOT NULL AUTO_INCREMENT,
  `deviceID` int,
  `deviceRoleDefinitionID` int,
  `blanked` boolean,
  PRIMARY KEY (displayID),
  INDEX dispDev_ind(deviceID),
  INDEX dispDevRol_ind(deviceRoleDefinitionID),
  FOREIGN KEY (deviceID)
    REFERENCES Devices(deviceID),
  FOREIGN KEY (deviceRoleDefinitionID)
    REFERENCES DeviceRoleDefinition(deviceRoleDefinitionID)
);

CREATE TABLE `DeviceCommands` (
  `deviceCommandID` int NOT NULL AUTO_INCREMENT,
  `deviceID` int,
  `commandID` int,
  `microserviceID` int,
  `endpointID` int,
  `enabled` boolean,
  PRIMARY KEY (deviceCommandID),
  INDEX devComDev_ind(deviceID),
  INDEX devComCom_ind(commandID),
  INDEX devComMS_ind(microserviceID),
  INDEX devComEnd_ind(endpointID),
  FOREIGN KEY (deviceID)
    REFERENCES Devices(deviceID),
  FOREIGN KEY (commandID)
    REFERENCES Commands(commandID),
  FOREIGN KEY (endpointID)
    REFERENCES Endpoints(endpointID),
  FOREIGN KEY (microserviceID)
    REFERENCES Microservices(microserviceID)
);

CREATE TABLE `Ports` (
  `portID` int NOT NULL AUTO_INCREMENT,
  `port` varchar(256),
  `description` text,
  PRIMARY KEY (portID)
);

CREATE TABLE `RoomConfiguration` (
  `roomConfigurationID` int NOT NULL AUTO_INCREMENT,
  `name` varchar(256),
  `description` text,
  PRIMARY KEY (roomConfigurationID)
);

CREATE TABLE `RoomConfigurationMapping` (
  `roomConfigurationMappingID` int NOT NULL AUTO_INCREMENT,
  `commandID` varchar(256),
  `roomConfigurationID` text,
  `commandCodeKey` varchar(256),
  PRIMARY KEY (roomConfigurationMappingID),
  FOREIGN KEY (commandID)
    REFERENCES Commands(commandID),
  FOREIGN KEY (roomConfigurationID)
    REFERENCES RoomConfiguration(roomConfigurationID)
);

CREATE TABLE `PortConfiguration` (
  `portConfigurationID` int NOT NULL AUTO_INCREMENT,
  `destinationDeviceID` int,
  `portID` int,
  `sourceDeviceID` int,
  PRIMARY KEY (portConfigurationID),
  FOREIGN KEY (destinationDeviceID)
    REFERENCES Devices(DeviceID),
  FOREIGN KEY (portID)
    REFERENCES Ports(portID),
  FOREIGN KEY (sourceDeviceID)
    REFERENCES Devices(DeviceID)
);

CREATE TABLE `DevicePowerState` (
  `devicePowerStateID` int NOT NULL AUTO_INCREMENT,
  `deviceID` int,
  `powerStateID` int,
  PRIMARY KEY (devicePowerStateID),
  FOREIGN KEY (deviceID)
    REFERENCES Devices(deviceID),
  FOREIGN KEY (powerStateID)
    REFERENCES PowerState(powerStateID)
);

ALTER TABLE Devices
ADD FOREIGN KEY (roomID)
  REFERENCES Rooms(roomID);
