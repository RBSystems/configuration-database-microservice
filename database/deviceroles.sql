RENAME TABLE `configuration`.DeviceTypes TO `configuration`.DeviceClasses ;

ALTER TABLE `configuration`.Devices CHANGE typeID classID int(11) NULL ;

ALTER TABLE `configuration`.Devices ADD typeID INT NULL ;

ALTER TABLE `configuration`.DeviceClasses CHANGE deviceTypeID deviceClassID int(11) NOT NULL auto_increment ;

CREATE TABLE `configuration`.DeviceTypes {
    deviceTypeID int NOT NULL AUTO_INCREMENT,
    typeName varchar(255) NOT NULL,
    typeDescription varchar(1024),
    typeDisplayName varchar(255),
    PRIMARY KEY (deviceTypeID)
}

CREATE TABLE `configuration`.DeviceTypeCommandMapping (
    deviceTypeCommandMappingID int NOT NULL AUTO_INCREMENT,
    deviceTypeID int NOT NULL,
    commandID int,
    microserviceID int,
    endpointID int,
    PRIMARY KEY (deviceTypeCommandMappingID),
    KEY `devTypComDev_ind` (`deviceTypeID`),
    KEY `devTypComCom_ind` (`commandID`),
    KEY `devTypComMS_ind` (`microserviceID`),
    KEY `devTypComEnd_ind` (`endpointID`),
    CONSTRAINT `DeviceTypeCommands_ibfk_1` FOREIGN KEY (`deviceTypeID`) REFERENCES `DeviceTypes` (`deviceTypeID`),
    CONSTRAINT `DeviceTypeCommands_ibfk_2` FOREIGN KEY (`commandID`) REFERENCES `Commands` (`commandID`),
    CONSTRAINT `DeviceTypeCommands_ibfk_3` FOREIGN KEY (`endpointID`) REFERENCES `Endpoints` (`endpointID`),
    CONSTRAINT `DeviceTypeCommands_ibfk_4` FOREIGN KEY (`microserviceID`) REFERENCES `Microservices` (`microserviceID`)
)

CREATE TABLE `configuration`.DeviceTypePorts (
    deviceTypePortID int NOT NULL AUTO_INCREMENT,
    deviceTypeID int NOT NULL,
    portID int NOT NULL,
    description varchar(1024) NOT NULL,
    friendlyName varchar(255) NOT NULL,
    sourceDestinationMirror BIT NOT NULL,
    PRIMARY KEY (deviceTypePortID),
    KEY `devTypPorTyp_ind` (`deviceTypeID`),
    KEY `devTypPorPor_ind` (`portID`),
    CONSTRAINT `deviceTypePorts_ibfk_1` FOREIGN KEY (`deviceTypeID`) REFERENCES `DeviceTypes` (`deviceTypeID`),
    CONSTRAINT `deviceTypePorts_ibfk_2` FOREIGN KEY (`portID`) REFERENCES `Ports` (`portID`)
)

ALTER TABLE `configuration`.Devices
ADD FOREIGN KEY (typeID) REFERENCES `DeviceTypes` (`deviceTypeID`);

ALTER TABLE `configuration`.Devices
ADD FOREIGN KEY (classID) REFERENCES `DeviceClasses` (`deviceClassID`);

-- We don't rework the Port Configuration Table - we just use the information in the deviceTypePorts to generate the information necessary to denote what ports there are (and thus more easliy generate the values from the tool)

--I'll probably write a stored procedure to add the new port configurations (with the destination Device/HostDevice/SourceDevice

DROP TABLE `configuration`.DeviceCommands

INSERT INTO `configuration`.DeviceTypePorts
(deviceTypePortID, deviceTypeID, portID, description, friendlyName, sourceDestinationMirror)
VALUES(1, 3, 36, 'HDMI 1', 'HDMI 1', 1);
INSERT INTO `configuration`.DeviceTypePorts
(deviceTypePortID, deviceTypeID, portID, description, friendlyName, sourceDestinationMirror)
VALUES(2, 3, 37, 'HDMI 2', 'HDMI 2', 1);
INSERT INTO `configuration`.DeviceTypePorts
(deviceTypePortID, deviceTypeID, portID, description, friendlyName, sourceDestinationMirror)
VALUES(3, 3, 39, 'HDMI 3', 'HDMI 3', 1);
INSERT INTO `configuration`.DeviceTypePorts
(deviceTypePortID, deviceTypeID, portID, description, friendlyName, sourceDestinationMirror)
VALUES(4, 3, 40, 'HDMI 4', 'HDMI 4', 1);
INSERT INTO `configuration`.DeviceTypePorts
(deviceTypePortID, deviceTypeID, portID, description, friendlyName, sourceDestinationMirror)
VALUES(5, 2, 13, 'PJLINK Network 1', 'Network 1', 1);
INSERT INTO `configuration`.DeviceTypePorts
(deviceTypePortID, deviceTypeID, portID, description, friendlyName, sourceDestinationMirror)
VALUES(6, 2, 14, 'PJLINK Digital 1', 'Digital 1', 1);
INSERT INTO `configuration`.DeviceTypePorts
(deviceTypePortID, deviceTypeID, portID, description, friendlyName, sourceDestinationMirror)
VALUES(7, 2, 14, 'PJLINK Digital 3', 'Digital 3', 1);
INSERT INTO `configuration`.DeviceTypePorts
(deviceTypePortID, deviceTypeID, portID, description, friendlyName, sourceDestinationMirror)
VALUES(8, 2, 38, 'PJLINK Digital 2', 'Digital 2', 1);
INSERT INTO `configuration`.DeviceTypePorts
(deviceTypePortID, deviceTypeID, portID, description, friendlyName, sourceDestinationMirror)
VALUES(9, 4, 16, 'PulseEight 1 - 1', '1 - 1', 0);
INSERT INTO `configuration`.DeviceTypePorts
(deviceTypePortID, deviceTypeID, portID, description, friendlyName, sourceDestinationMirror)
VALUES(10, 4, 17, 'Pulse Eight 1 - 2 ', '1 - 2', 0);
INSERT INTO `configuration`.DeviceTypePorts
(deviceTypePortID, deviceTypeID, portID, description, friendlyName, sourceDestinationMirror)
VALUES(11, 4, 18, 'Pulse Eight 1 - 3', '1 - 3', 0);
INSERT INTO `configuration`.DeviceTypePorts
(deviceTypePortID, deviceTypeID, portID, description, friendlyName, sourceDestinationMirror)
VALUES(12, 4, 19, 'Pulse Eight 1 - 4 ', '1 - 4', 0);
INSERT INTO `configuration`.DeviceTypePorts
(deviceTypePortID, deviceTypeID, portID, description, friendlyName, sourceDestinationMirror)
VALUES(13, 4, 20, 'Pulse Eight 2 - 1', '2 - 1 ', 0);
INSERT INTO `configuration`.DeviceTypePorts
(deviceTypePortID, deviceTypeID, portID, description, friendlyName, sourceDestinationMirror)
VALUES(14, 4, 21, 'Pulse Eight 2 - 2 ', '2 - 2 ', 0);
INSERT INTO `configuration`.DeviceTypePorts
(deviceTypePortID, deviceTypeID, portID, description, friendlyName, sourceDestinationMirror)
VALUES(15, 4, 22, 'Pulse Eight 2 - 3', '2 - 3', 0);
INSERT INTO `configuration`.DeviceTypePorts
(deviceTypePortID, deviceTypeID, portID, description, friendlyName, sourceDestinationMirror)
VALUES(16, 4, 23, 'Pulse Eight 2 - 4  ', '2 - 4 ', 0);
INSERT INTO `configuration`.DeviceTypePorts
(deviceTypePortID, deviceTypeID, portID, description, friendlyName, sourceDestinationMirror)
VALUES(17, 4, 24, 'Pulse Eight 3 - 1', '3 - 1', 0);
INSERT INTO `configuration`.DeviceTypePorts
(deviceTypePortID, deviceTypeID, portID, description, friendlyName, sourceDestinationMirror)
VALUES(18, 4, 25, 'Pulse Eight 3 - 2 ', '3 - 2', 0);
INSERT INTO `configuration`.DeviceTypePorts
(deviceTypePortID, deviceTypeID, portID, description, friendlyName, sourceDestinationMirror)
VALUES(19, 4, 26, 'Pulse Eight 3 - 3', '3 - 3', 0);
INSERT INTO `configuration`.DeviceTypePorts
(deviceTypePortID, deviceTypeID, portID, description, friendlyName, sourceDestinationMirror)
VALUES(20, 4, 27, 'Pulse Eight 3 - 4', '3 - 4', 0);
INSERT INTO `configuration`.DeviceTypePorts
(deviceTypePortID, deviceTypeID, portID, description, friendlyName, sourceDestinationMirror)
VALUES(21, 4, 28, 'Pulse Eight 4 - 1', '4 - 1', 0);
INSERT INTO `configuration`.DeviceTypePorts
(deviceTypePortID, deviceTypeID, portID, description, friendlyName, sourceDestinationMirror)
VALUES(22, 4, 29, 'Pulse Eight 4 - 2', '4 - 2', 0);
INSERT INTO `configuration`.DeviceTypePorts
(deviceTypePortID, deviceTypeID, portID, description, friendlyName, sourceDestinationMirror)
VALUES(23, 4, 30, 'Pulse Eight 4 - 3', '4 - 3', 0);
INSERT INTO `configuration`.DeviceTypePorts
(deviceTypePortID, deviceTypeID, portID, description, friendlyName, sourceDestinationMirror)
VALUES(24, 4, 31, 'Pulse Eight 4 - 4', '4 - 4', 0);
INSERT INTO `configuration`.DeviceTypePorts
(deviceTypePortID, deviceTypeID, portID, description, friendlyName, sourceDestinationMirror)
VALUES(25, 5, 41, 'Blu Channel 1', 'Channel 1', 0);
INSERT INTO `configuration`.DeviceTypePorts
(deviceTypePortID, deviceTypeID, portID, description, friendlyName, sourceDestinationMirror)
VALUES(26, 5, 42, 'Blu Channel 2', 'Channel 2', 0);
INSERT INTO `configuration`.DeviceTypePorts
(deviceTypePortID, deviceTypeID, portID, description, friendlyName, sourceDestinationMirror)
VALUES(27, 5, 43, 'Blu Channel 3', 'Channel 3', 0);
INSERT INTO `configuration`.DeviceTypePorts
(deviceTypePortID, deviceTypeID, portID, description, friendlyName, sourceDestinationMirror)
VALUES(28, 5, 44, 'Blu Channel 4', 'Channel 4 ', 0);
INSERT INTO `configuration`.DeviceTypePorts
(deviceTypePortID, deviceTypeID, portID, description, friendlyName, sourceDestinationMirror)
VALUES(29, 5, 45, 'Blu Channel 5', 'Channel 5 ', 0);
INSERT INTO `configuration`.DeviceTypePorts
(deviceTypePortID, deviceTypeID, portID, description, friendlyName, sourceDestinationMirror)
VALUES(30, 5, 46, 'Blu Channel', 'Channel 6', 0);
INSERT INTO `configuration`.DeviceTypePorts
(deviceTypePortID, deviceTypeID, portID, description, friendlyName, sourceDestinationMirror)
VALUES(31, 5, 47, 'Blu Channel', 'Channel 7 ', 0);
INSERT INTO `configuration`.DeviceTypePorts
(deviceTypePortID, deviceTypeID, portID, description, friendlyName, sourceDestinationMirror)
VALUES(32, 5, 48, 'Blu Channel', 'Channel 8', 0);
INSERT INTO `configuration`.DeviceTypePorts
(deviceTypePortID, deviceTypeID, portID, description, friendlyName, sourceDestinationMirror)
VALUES(33, 5, 49, 'Blu Channel', 'Channel 9', 0);
INSERT INTO `configuration`.DeviceTypePorts
(deviceTypePortID, deviceTypeID, portID, description, friendlyName, sourceDestinationMirror)
VALUES(34, 5, 50, 'Blu Channel', 'Channel 10', 0);
INSERT INTO `configuration`.DeviceTypePorts
(deviceTypePortID, deviceTypeID, portID, description, friendlyName, sourceDestinationMirror)
VALUES(35, 5, 51, 'Blu Channel', 'Channel 11', 0);
INSERT INTO `configuration`.DeviceTypePorts
(deviceTypePortID, deviceTypeID, portID, description, friendlyName, sourceDestinationMirror)
VALUES(37, 6, 53, 'Corresponds to all Shure channels', 'All Channels', 0);
INSERT INTO `configuration`.DeviceTypePorts
(deviceTypePortID, deviceTypeID, portID, description, friendlyName, sourceDestinationMirror)
VALUES(38, 6, 54, 'Shure Receiver Channel 1', 'Channel 1', 0);
INSERT INTO `configuration`.DeviceTypePorts
(deviceTypePortID, deviceTypeID, portID, description, friendlyName, sourceDestinationMirror)
VALUES(39, 6, 55, 'Shure Reciever Channel 2', 'Channel 2', 0);
INSERT INTO `configuration`.DeviceTypePorts
(deviceTypePortID, deviceTypeID, portID, description, friendlyName, sourceDestinationMirror)
VALUES(40, 6, 56, 'Shure Reciever Channel 3', 'Channel 3', 0);
INSERT INTO `configuration`.DeviceTypePorts
(deviceTypePortID, deviceTypeID, portID, description, friendlyName, sourceDestinationMirror)
VALUES(41, 6, 57, 'Shure Receiver Channel 4', 'Channel 4', 0);
INSERT INTO `configuration`.DeviceTypePorts
(deviceTypePortID, deviceTypeID, portID, description, friendlyName, sourceDestinationMirror)
VALUES(36, 5, 52, 'Blu Channel', 'Channel 12', 0);


INSERT INTO `configuration`.DeviceTypeCommandMapping
(deviceTypeCommandMappingID, deviceTypeID, commandID, microserviceID, endpointID)
VALUES(1, 3, 2, 2, 2);
INSERT INTO `configuration`.DeviceTypeCommandMapping
(deviceTypeCommandMappingID, deviceTypeID, commandID, microserviceID, endpointID)
VALUES(2, 3, 1, 2, 1);
INSERT INTO `configuration`.DeviceTypeCommandMapping
(deviceTypeCommandMappingID, deviceTypeID, commandID, microserviceID, endpointID)
VALUES(3, 3, 16, 2, 14);
INSERT INTO `configuration`.DeviceTypeCommandMapping
(deviceTypeCommandMappingID, deviceTypeID, commandID, microserviceID, endpointID)
VALUES(4, 3, 18, 2, 15);
INSERT INTO `configuration`.DeviceTypeCommandMapping
(deviceTypeCommandMappingID, deviceTypeID, commandID, microserviceID, endpointID)
VALUES(5, 3, 19, 2, 16);
INSERT INTO `configuration`.DeviceTypeCommandMapping
(deviceTypeCommandMappingID, deviceTypeID, commandID, microserviceID, endpointID)
VALUES(6, 3, 20, 2, 18);
INSERT INTO `configuration`.DeviceTypeCommandMapping
(deviceTypeCommandMappingID, deviceTypeID, commandID, microserviceID, endpointID)
VALUES(7, 3, 3, 2, 3);
INSERT INTO `configuration`.DeviceTypeCommandMapping
(deviceTypeCommandMappingID, deviceTypeID, commandID, microserviceID, endpointID)
VALUES(8, 3, 4, 2, 4);
INSERT INTO `configuration`.DeviceTypeCommandMapping
(deviceTypeCommandMappingID, deviceTypeID, commandID, microserviceID, endpointID)
VALUES(9, 3, 5, 2, 5);
INSERT INTO `configuration`.DeviceTypeCommandMapping
(deviceTypeCommandMappingID, deviceTypeID, commandID, microserviceID, endpointID)
VALUES(10, 3, 6, 2, 6);
INSERT INTO `configuration`.DeviceTypeCommandMapping
(deviceTypeCommandMappingID, deviceTypeID, commandID, microserviceID, endpointID)
VALUES(11, 3, 9, 2, 9);
INSERT INTO `configuration`.DeviceTypeCommandMapping
(deviceTypeCommandMappingID, deviceTypeID, commandID, microserviceID, endpointID)
VALUES(12, 3, 10, 2, 10);
INSERT INTO `configuration`.DeviceTypeCommandMapping
(deviceTypeCommandMappingID, deviceTypeID, commandID, microserviceID, endpointID)
VALUES(13, 3, 17, 2, 17);
INSERT INTO `configuration`.DeviceTypeCommandMapping
(deviceTypeCommandMappingID, deviceTypeID, commandID, microserviceID, endpointID)
VALUES(14, 2, 1, 6, 1);
INSERT INTO `configuration`.DeviceTypeCommandMapping
(deviceTypeCommandMappingID, deviceTypeID, commandID, microserviceID, endpointID)
VALUES(15, 2, 2, 6, 2);
INSERT INTO `configuration`.DeviceTypeCommandMapping
(deviceTypeCommandMappingID, deviceTypeID, commandID, microserviceID, endpointID)
VALUES(16, 2, 3, 1, 3);
INSERT INTO `configuration`.DeviceTypeCommandMapping
(deviceTypeCommandMappingID, deviceTypeID, commandID, microserviceID, endpointID)
VALUES(17, 2, 4, 6, 4);
INSERT INTO `configuration`.DeviceTypeCommandMapping
(deviceTypeCommandMappingID, deviceTypeID, commandID, microserviceID, endpointID)
VALUES(18, 2, 5, 1, 5);
INSERT INTO `configuration`.DeviceTypeCommandMapping
(deviceTypeCommandMappingID, deviceTypeID, commandID, microserviceID, endpointID)
VALUES(19, 2, 6, 1, 6);
INSERT INTO `configuration`.DeviceTypeCommandMapping
(deviceTypeCommandMappingID, deviceTypeID, commandID, microserviceID, endpointID)
VALUES(20, 2, 9, 6, 9);
INSERT INTO `configuration`.DeviceTypeCommandMapping
(deviceTypeCommandMappingID, deviceTypeID, commandID, microserviceID, endpointID)
VALUES(21, 2, 10, 6, 10);
INSERT INTO `configuration`.DeviceTypeCommandMapping
(deviceTypeCommandMappingID, deviceTypeID, commandID, microserviceID, endpointID)
VALUES(22, 2, 16, 6, 14);
INSERT INTO `configuration`.DeviceTypeCommandMapping
(deviceTypeCommandMappingID, deviceTypeID, commandID, microserviceID, endpointID)
VALUES(23, 2, 17, 1, 17);
INSERT INTO `configuration`.DeviceTypeCommandMapping
(deviceTypeCommandMappingID, deviceTypeID, commandID, microserviceID, endpointID)
VALUES(24, 2, 18, 6, 15);
INSERT INTO `configuration`.DeviceTypeCommandMapping
(deviceTypeCommandMappingID, deviceTypeID, commandID, microserviceID, endpointID)
VALUES(25, 2, 19, 6, 16);
INSERT INTO `configuration`.DeviceTypeCommandMapping
(deviceTypeCommandMappingID, deviceTypeID, commandID, microserviceID, endpointID)
VALUES(26, 2, 20, 1, 18);
INSERT INTO `configuration`.DeviceTypeCommandMapping
(deviceTypeCommandMappingID, deviceTypeID, commandID, microserviceID, endpointID)
VALUES(27, 4, 3, 5, 13);
INSERT INTO `configuration`.DeviceTypeCommandMapping
(deviceTypeCommandMappingID, deviceTypeID, commandID, microserviceID, endpointID)
VALUES(28, 4, 17, 5, 20);
INSERT INTO `configuration`.DeviceTypeCommandMapping
(deviceTypeCommandMappingID, deviceTypeID, commandID, microserviceID, endpointID)
VALUES(29, 5, 4, 7, 25);
INSERT INTO `configuration`.DeviceTypeCommandMapping
(deviceTypeCommandMappingID, deviceTypeID, commandID, microserviceID, endpointID)
VALUES(30, 5, 9, 7, 26);
INSERT INTO `configuration`.DeviceTypeCommandMapping
(deviceTypeCommandMappingID, deviceTypeID, commandID, microserviceID, endpointID)
VALUES(31, 5, 10, 7, 27);
INSERT INTO `configuration`.DeviceTypeCommandMapping
(deviceTypeCommandMappingID, deviceTypeID, commandID, microserviceID, endpointID)
VALUES(32, 5, 22, 7, 23);
INSERT INTO `configuration`.DeviceTypeCommandMapping
(deviceTypeCommandMappingID, deviceTypeID, commandID, microserviceID, endpointID)
VALUES(33, 5, 21, 7, 24);
INSERT INTO `configuration`.DeviceTypeCommandMapping
(deviceTypeCommandMappingID, deviceTypeID, commandID, microserviceID, endpointID)
VALUES(34, 5, 1, 7, 1);
INSERT INTO `configuration`.DeviceTypeCommandMapping
(deviceTypeCommandMappingID, deviceTypeID, commandID, microserviceID, endpointID)
VALUES(35, 5, 2, 7, 2);

INSERT INTO `configuration`.DeviceTypes
(deviceTypeID, typeName, typeDescription, typeDisplayName)
VALUES(2, 'SonyVPL', 'The Sony VPL Projector line', 'Sony VPL');
INSERT INTO `configuration`.DeviceTypes
(deviceTypeID, typeName, typeDescription, typeDisplayName)
VALUES(3, 'SonyXBR', 'The Sony XBR TV line.', 'Sony XBR');
INSERT INTO `configuration`.DeviceTypes
(deviceTypeID, typeName, typeDescription, typeDisplayName)
VALUES(4, 'PulseEight4x4', 'A Pulse Eight 4x6 Matrix', 'PulseEight 4x4');
INSERT INTO `configuration`.DeviceTypes
(deviceTypeID, typeName, typeDescription, typeDisplayName)
VALUES(5, 'Blu50', 'A London Blu 50 DSP', 'London Blu 50');
INSERT INTO `configuration`.DeviceTypes
(deviceTypeID, typeName, typeDescription, typeDisplayName)
VALUES(6, 'ShureULXD', 'A Shure ULXD device', 'Shure ULXD');
INSERT INTO `configuration`.DeviceTypes
(deviceTypeID, typeName, typeDescription, typeDisplayName)
VALUES(7, 'Pi3', 'A Raspberry Pi 3', 'Raspberry Pi 3');
INSERT INTO `configuration`.DeviceTypes
(deviceTypeID, typeName, typeDescription, typeDisplayName)
VALUES(8, 'AppleTV', 'An Apple TV', NULL);
INSERT INTO `configuration`.DeviceTypes
(deviceTypeID, typeName, typeDescription, typeDisplayName)
VALUES(9, 'ChromeCast', 'A ChromeCast', 'Chrome Cast');
