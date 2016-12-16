# ************************************************************
# Sequel Pro SQL dump
# Version 4541
#
# http://www.sequelpro.com/
# https://github.com/sequelpro/sequelpro
#
# Host: byuoitav-configuration-database.cserrr6xcwqt.us-west-2.rds.amazonaws.com (MySQL 5.5.5-10.0.24-MariaDB)
# Database: configuration
# Generation Time: 2016-12-16 20:51:44 +0000
# ************************************************************


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;


# Dump of table AudioDevices
# ------------------------------------------------------------

DROP TABLE IF EXISTS `AudioDevices`;

CREATE TABLE `AudioDevices` (
  `audioDeviceID` int(11) NOT NULL AUTO_INCREMENT,
  `deviceID` int(11) DEFAULT NULL,
  `deviceRoleDefinitionID` int(11) DEFAULT NULL,
  `muted` tinyint(1) DEFAULT NULL,
  `volume` int(11) DEFAULT NULL,
  PRIMARY KEY (`audioDeviceID`),
  KEY `audDev_ind` (`deviceID`),
  KEY `audDevRol_ind` (`deviceRoleDefinitionID`),
  CONSTRAINT `AudioDevices_ibfk_1` FOREIGN KEY (`deviceID`) REFERENCES `Devices` (`deviceID`),
  CONSTRAINT `AudioDevices_ibfk_2` FOREIGN KEY (`deviceRoleDefinitionID`) REFERENCES `DeviceRoleDefinition` (`deviceRoleDefinitionID`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

LOCK TABLES `AudioDevices` WRITE;
/*!40000 ALTER TABLE `AudioDevices` DISABLE KEYS */;

INSERT INTO `AudioDevices` (`audioDeviceID`, `deviceID`, `deviceRoleDefinitionID`, `muted`, `volume`)
VALUES
	(1,1,1,0,0),
	(2,7,1,0,50);

/*!40000 ALTER TABLE `AudioDevices` ENABLE KEYS */;
UNLOCK TABLES;


# Dump of table Buildings
# ------------------------------------------------------------

DROP TABLE IF EXISTS `Buildings`;

CREATE TABLE `Buildings` (
  `buildingID` int(11) NOT NULL AUTO_INCREMENT,
  `name` text,
  `shortName` varchar(256) DEFAULT NULL,
  `description` text,
  PRIMARY KEY (`buildingID`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

LOCK TABLES `Buildings` WRITE;
/*!40000 ALTER TABLE `Buildings` DISABLE KEYS */;

INSERT INTO `Buildings` (`buildingID`, `name`, `shortName`, `description`)
VALUES
	(1,'Information Technology Building','ITB','Generic description to make up for (temporarily) lazy code');

/*!40000 ALTER TABLE `Buildings` ENABLE KEYS */;
UNLOCK TABLES;


# Dump of table Commands
# ------------------------------------------------------------

DROP TABLE IF EXISTS `Commands`;

CREATE TABLE `Commands` (
  `commandID` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(256) DEFAULT NULL,
  `description` text,
  `Priority` int(11) DEFAULT NULL,
  PRIMARY KEY (`commandID`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

LOCK TABLES `Commands` WRITE;
/*!40000 ALTER TABLE `Commands` DISABLE KEYS */;

INSERT INTO `Commands` (`commandID`, `name`, `description`, `Priority`)
VALUES
	(1,'PowerOn','Pull out of standby',1),
	(2,'Standby','Put into standby',1),
	(3,'ChangeInput','Change the input to the supplied port (hdmi 1, hdmi 2) etc. etc.)',10),
	(4,'SetVolume','Change the volume to supplied value',10),
	(5,'BlankScreen','Blanks the screen',10),
	(6,'UnblankScreen','Unblanks the screen',7),
	(7,'VolumeDown','Tick volume down',10),
	(8,'VolumeUp','Tick volume up',10),
	(9,'Mute','Mute\n',10),
	(10,'UnMute','UnMute',10);

/*!40000 ALTER TABLE `Commands` ENABLE KEYS */;
UNLOCK TABLES;


# Dump of table DeviceCommands
# ------------------------------------------------------------

DROP TABLE IF EXISTS `DeviceCommands`;

CREATE TABLE `DeviceCommands` (
  `deviceCommandID` int(11) NOT NULL AUTO_INCREMENT,
  `deviceID` int(11) DEFAULT NULL,
  `commandID` int(11) DEFAULT NULL,
  `microserviceID` int(11) DEFAULT NULL,
  `endpointID` int(11) DEFAULT NULL,
  `enabled` tinyint(1) DEFAULT NULL,
  PRIMARY KEY (`deviceCommandID`),
  KEY `devComDev_ind` (`deviceID`),
  KEY `devComCom_ind` (`commandID`),
  KEY `devComMS_ind` (`microserviceID`),
  KEY `devComEnd_ind` (`endpointID`),
  CONSTRAINT `DeviceCommands_ibfk_1` FOREIGN KEY (`deviceID`) REFERENCES `Devices` (`deviceID`),
  CONSTRAINT `DeviceCommands_ibfk_2` FOREIGN KEY (`commandID`) REFERENCES `Commands` (`commandID`),
  CONSTRAINT `DeviceCommands_ibfk_3` FOREIGN KEY (`endpointID`) REFERENCES `Endpoints` (`endpointID`),
  CONSTRAINT `DeviceCommands_ibfk_4` FOREIGN KEY (`microserviceID`) REFERENCES `Microservices` (`microserviceID`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

LOCK TABLES `DeviceCommands` WRITE;
/*!40000 ALTER TABLE `DeviceCommands` DISABLE KEYS */;

INSERT INTO `DeviceCommands` (`deviceCommandID`, `deviceID`, `commandID`, `microserviceID`, `endpointID`, `enabled`)
VALUES
	(1,1,1,2,1,1),
	(2,1,2,2,2,1),
	(3,1,3,2,3,1),
	(4,1,4,2,4,1),
	(5,1,5,2,5,1),
	(6,1,6,2,6,1),
	(7,1,7,2,8,1),
	(8,1,8,2,7,1),
	(10,6,1,1,1,1),
	(11,6,2,1,2,1),
	(12,7,1,2,1,1),
	(13,7,2,2,2,1),
	(14,7,3,2,3,1),
	(15,7,4,2,4,1),
	(16,7,5,2,5,1),
	(17,7,6,2,6,1),
	(18,7,7,2,7,1),
	(19,7,8,2,8,1),
	(20,7,9,2,9,1),
	(21,7,10,2,10,1),
	(22,6,3,1,3,1),
	(23,6,5,1,5,1),
	(24,6,6,1,6,1);

/*!40000 ALTER TABLE `DeviceCommands` ENABLE KEYS */;
UNLOCK TABLES;


# Dump of table DevicePowerStates
# ------------------------------------------------------------

DROP TABLE IF EXISTS `DevicePowerStates`;

CREATE TABLE `DevicePowerStates` (
  `devicePowerStateID` int(11) NOT NULL AUTO_INCREMENT,
  `deviceID` int(11) DEFAULT NULL,
  `powerStateID` int(11) DEFAULT NULL,
  PRIMARY KEY (`devicePowerStateID`),
  KEY `deviceID` (`deviceID`),
  KEY `powerStateID` (`powerStateID`),
  CONSTRAINT `DevicePowerStates_ibfk_1` FOREIGN KEY (`deviceID`) REFERENCES `Devices` (`deviceID`),
  CONSTRAINT `DevicePowerStates_ibfk_2` FOREIGN KEY (`powerStateID`) REFERENCES `PowerStates` (`powerStateID`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

LOCK TABLES `DevicePowerStates` WRITE;
/*!40000 ALTER TABLE `DevicePowerStates` DISABLE KEYS */;

INSERT INTO `DevicePowerStates` (`devicePowerStateID`, `deviceID`, `powerStateID`)
VALUES
	(1,1,1),
	(2,1,2),
	(3,6,1),
	(4,6,2),
	(5,7,1),
	(6,7,2);

/*!40000 ALTER TABLE `DevicePowerStates` ENABLE KEYS */;
UNLOCK TABLES;


# Dump of table DeviceRole
# ------------------------------------------------------------

DROP TABLE IF EXISTS `DeviceRole`;

CREATE TABLE `DeviceRole` (
  `deviceRoleID` int(11) NOT NULL AUTO_INCREMENT,
  `deviceID` int(11) DEFAULT NULL,
  `deviceRoleDefinitionID` int(11) DEFAULT NULL,
  PRIMARY KEY (`deviceRoleID`),
  KEY `devRolDevID_ind` (`deviceID`),
  KEY `devRolDevRolDef_ind` (`deviceRoleDefinitionID`),
  CONSTRAINT `DeviceRole_ibfk_1` FOREIGN KEY (`deviceID`) REFERENCES `Devices` (`deviceID`),
  CONSTRAINT `DeviceRole_ibfk_2` FOREIGN KEY (`deviceRoleDefinitionID`) REFERENCES `DeviceRoleDefinition` (`deviceRoleDefinitionID`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

LOCK TABLES `DeviceRole` WRITE;
/*!40000 ALTER TABLE `DeviceRole` DISABLE KEYS */;

INSERT INTO `DeviceRole` (`deviceRoleID`, `deviceID`, `deviceRoleDefinitionID`)
VALUES
	(1,1,1),
	(2,1,2),
	(3,2,3),
	(4,2,4),
	(5,3,3),
	(6,3,4),
	(7,4,3),
	(8,4,4),
	(9,5,6),
	(11,5,5),
	(12,6,2),
	(13,7,2),
	(14,8,3),
	(15,10,3),
	(16,8,4),
	(17,10,4),
	(18,7,1),
	(19,11,5);

/*!40000 ALTER TABLE `DeviceRole` ENABLE KEYS */;
UNLOCK TABLES;


# Dump of table DeviceRoleDefinition
# ------------------------------------------------------------

DROP TABLE IF EXISTS `DeviceRoleDefinition`;

CREATE TABLE `DeviceRoleDefinition` (
  `deviceRoleDefinitionID` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(256) DEFAULT NULL,
  `description` text,
  PRIMARY KEY (`deviceRoleDefinitionID`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

LOCK TABLES `DeviceRoleDefinition` WRITE;
/*!40000 ALTER TABLE `DeviceRoleDefinition` DISABLE KEYS */;

INSERT INTO `DeviceRoleDefinition` (`deviceRoleDefinitionID`, `name`, `description`)
VALUES
	(1,'AudioOut','Device that outputs audio (speakers, tv, etc.)'),
	(2,'VideoOut','Device that displays video (projector, tv, etc.)'),
	(3,'AudioIn','Device that provides Audio input (computer, 3.5mm jack, hdmi, etc.)'),
	(4,'VideoIn','Device that provides Video input (computer, HDMI, VGA, etc.)'),
	(5,'ControlProcessor','A device that controls other devices in the room'),
	(6,'Touchpanel','The touch interface for controlling devices in a room');

/*!40000 ALTER TABLE `DeviceRoleDefinition` ENABLE KEYS */;
UNLOCK TABLES;


# Dump of table Devices
# ------------------------------------------------------------

DROP TABLE IF EXISTS `Devices`;

CREATE TABLE `Devices` (
  `deviceID` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(256) DEFAULT NULL,
  `address` varchar(256) DEFAULT NULL,
  `input` tinyint(1) DEFAULT NULL,
  `output` tinyint(1) DEFAULT NULL,
  `buildingID` int(11) DEFAULT NULL,
  `roomID` int(11) DEFAULT NULL,
  `typeID` int(11) DEFAULT NULL,
  `powerID` int(11) DEFAULT NULL,
  `responding` tinyint(1) DEFAULT NULL,
  PRIMARY KEY (`deviceID`),
  KEY `devbld_ind` (`buildingID`),
  KEY `devrm_ind` (`roomID`),
  KEY `devty_ind` (`typeID`),
  KEY `devpw_ind` (`powerID`),
  CONSTRAINT `Devices_ibfk_1` FOREIGN KEY (`buildingID`) REFERENCES `Buildings` (`buildingID`),
  CONSTRAINT `Devices_ibfk_2` FOREIGN KEY (`typeID`) REFERENCES `DeviceTypes` (`deviceTypeID`),
  CONSTRAINT `Devices_ibfk_3` FOREIGN KEY (`powerID`) REFERENCES `PowerStates` (`powerStateID`),
  CONSTRAINT `Devices_ibfk_4` FOREIGN KEY (`roomID`) REFERENCES `Rooms` (`roomID`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

LOCK TABLES `Devices` WRITE;
/*!40000 ALTER TABLE `Devices` DISABLE KEYS */;

INSERT INTO `Devices` (`deviceID`, `name`, `address`, `input`, `output`, `buildingID`, `roomID`, `typeID`, `powerID`, `responding`)
VALUES
	(1,'D1','10.66.9.6',0,1,1,1,1,1,1),
	(2,'PC1','10.0.0.0',1,0,1,1,2,1,1),
	(3,'Roku','10.0.0.0',1,0,1,1,2,1,1),
	(4,'IPTV','10.0.0.0',1,0,1,1,3,1,1),
	(5,'CP1','10.5.34.102',0,0,1,2,5,1,1),
	(6,'D2','10.5.34.95',0,1,1,2,6,1,1),
	(7,'D1','10.5.34.146',0,1,1,2,1,1,1),
	(8,'AppleTV','10.0.0.0',1,0,1,2,7,1,1),
	(10,'HDMIIn','10.0.0.0',1,0,1,2,8,1,1),
	(11,'CP2','10.5.34.94',0,0,1,2,5,1,1);

/*!40000 ALTER TABLE `Devices` ENABLE KEYS */;
UNLOCK TABLES;


# Dump of table DeviceTypes
# ------------------------------------------------------------

DROP TABLE IF EXISTS `DeviceTypes`;

CREATE TABLE `DeviceTypes` (
  `deviceTypeID` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(256) DEFAULT NULL,
  `description` text,
  PRIMARY KEY (`deviceTypeID`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

LOCK TABLES `DeviceTypes` WRITE;
/*!40000 ALTER TABLE `DeviceTypes` DISABLE KEYS */;

INSERT INTO `DeviceTypes` (`deviceTypeID`, `name`, `description`)
VALUES
	(1,'tv','A TV'),
	(2,'computer','A computer input'),
	(3,'roku','A Roku device'),
	(4,'iptv','An IPTV settop box'),
	(5,'pi','A Raspberry Pi touchpanel'),
	(6,'projector','A projector\n'),
	(7,'appletv','An AppleTV\n'),
	(8,'hdmiin','An HDMI input cable\n');

/*!40000 ALTER TABLE `DeviceTypes` ENABLE KEYS */;
UNLOCK TABLES;


# Dump of table Displays
# ------------------------------------------------------------

DROP TABLE IF EXISTS `Displays`;

CREATE TABLE `Displays` (
  `displayID` int(11) NOT NULL AUTO_INCREMENT,
  `deviceID` int(11) DEFAULT NULL,
  `deviceRoleDefinitionID` int(11) DEFAULT NULL,
  `blanked` tinyint(1) DEFAULT NULL,
  PRIMARY KEY (`displayID`),
  KEY `dispDev_ind` (`deviceID`),
  KEY `dispDevRol_ind` (`deviceRoleDefinitionID`),
  CONSTRAINT `Displays_ibfk_1` FOREIGN KEY (`deviceID`) REFERENCES `Devices` (`deviceID`),
  CONSTRAINT `Displays_ibfk_2` FOREIGN KEY (`deviceRoleDefinitionID`) REFERENCES `DeviceRoleDefinition` (`deviceRoleDefinitionID`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

LOCK TABLES `Displays` WRITE;
/*!40000 ALTER TABLE `Displays` DISABLE KEYS */;

INSERT INTO `Displays` (`displayID`, `deviceID`, `deviceRoleDefinitionID`, `blanked`)
VALUES
	(1,1,2,0);

/*!40000 ALTER TABLE `Displays` ENABLE KEYS */;
UNLOCK TABLES;


# Dump of table Endpoints
# ------------------------------------------------------------

DROP TABLE IF EXISTS `Endpoints`;

CREATE TABLE `Endpoints` (
  `endpointID` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(256) DEFAULT NULL,
  `path` text,
  `description` text,
  PRIMARY KEY (`endpointID`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

LOCK TABLES `Endpoints` WRITE;
/*!40000 ALTER TABLE `Endpoints` DISABLE KEYS */;

INSERT INTO `Endpoints` (`endpointID`, `name`, `path`, `description`)
VALUES
	(1,'PowerOn','/:address/power/on','Standard PowerOn endpoint.'),
	(2,'Standby','/:address/power/standby','Standard standby endpoint.'),
	(3,'ChangeInput','/:address/input/:port','Standard ChangeInput endpoint.'),
	(4,'SetVolume','/:address/volume/set/:difference','SetVolume endpoint for devices with only Volume Up/Volume Down Controls. '),
	(5,'BlankScreen','/:address/display/blank','Standard BlankScreen endpoint.'),
	(6,'UnblankScreen','/:address/display/unblank','Standard UnblankScreen endpoint.'),
	(7,'VolumeUp','/:address/volume/up','Standard SetVolume endpoint.'),
	(8,'VolumeDown','/:address/volume/down','Standard SetVolume endpoint.'),
	(9,'Mute','/:address/volume/mute','Standard Mute endpoint'),
	(10,'UnMute','/:address/volume/unmute','Standard UnMute endpoint');

/*!40000 ALTER TABLE `Endpoints` ENABLE KEYS */;
UNLOCK TABLES;


# Dump of table Microservices
# ------------------------------------------------------------

DROP TABLE IF EXISTS `Microservices`;

CREATE TABLE `Microservices` (
  `microserviceID` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(256) DEFAULT NULL,
  `address` text,
  `description` text,
  PRIMARY KEY (`microserviceID`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

LOCK TABLES `Microservices` WRITE;
/*!40000 ALTER TABLE `Microservices` DISABLE KEYS */;

INSERT INTO `Microservices` (`microserviceID`, `name`, `address`, `description`)
VALUES
	(1,'pjlink-microservice','http://localhost:8005',NULL),
	(2,'sony-control-microservice','http://localhost:8007',NULL);

/*!40000 ALTER TABLE `Microservices` ENABLE KEYS */;
UNLOCK TABLES;


# Dump of table PortConfiguration
# ------------------------------------------------------------

DROP TABLE IF EXISTS `PortConfiguration`;

CREATE TABLE `PortConfiguration` (
  `portConfigurationID` int(11) NOT NULL AUTO_INCREMENT,
  `destinationDeviceID` int(11) DEFAULT NULL,
  `portID` int(11) DEFAULT NULL,
  `sourceDeviceID` int(11) DEFAULT NULL,
  PRIMARY KEY (`portConfigurationID`),
  KEY `destinationDeviceID` (`destinationDeviceID`),
  KEY `portID` (`portID`),
  KEY `sourceDeviceID` (`sourceDeviceID`),
  CONSTRAINT `PortConfiguration_ibfk_1` FOREIGN KEY (`destinationDeviceID`) REFERENCES `Devices` (`deviceID`),
  CONSTRAINT `PortConfiguration_ibfk_2` FOREIGN KEY (`portID`) REFERENCES `Ports` (`portID`),
  CONSTRAINT `PortConfiguration_ibfk_3` FOREIGN KEY (`sourceDeviceID`) REFERENCES `Devices` (`deviceID`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

LOCK TABLES `PortConfiguration` WRITE;
/*!40000 ALTER TABLE `PortConfiguration` DISABLE KEYS */;

INSERT INTO `PortConfiguration` (`portConfigurationID`, `destinationDeviceID`, `portID`, `sourceDeviceID`)
VALUES
	(1,1,1,2),
	(2,1,2,4),
	(3,1,3,3),
	(4,7,1,8),
	(6,7,3,10),
	(7,6,6,8),
	(9,6,5,10);

/*!40000 ALTER TABLE `PortConfiguration` ENABLE KEYS */;
UNLOCK TABLES;


# Dump of table Ports
# ------------------------------------------------------------

DROP TABLE IF EXISTS `Ports`;

CREATE TABLE `Ports` (
  `portID` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(256) DEFAULT NULL,
  `description` text,
  PRIMARY KEY (`portID`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

LOCK TABLES `Ports` WRITE;
/*!40000 ALTER TABLE `Ports` DISABLE KEYS */;

INSERT INTO `Ports` (`portID`, `name`, `description`)
VALUES
	(1,'Hdmi1',NULL),
	(2,'Hdmi2',NULL),
	(3,'Hdmi3',NULL),
	(4,'Hdmi4',NULL),
	(5,'network6','Epson Network6 Port\n'),
	(6,'digital2','Epson Digital 2 port\n');

/*!40000 ALTER TABLE `Ports` ENABLE KEYS */;
UNLOCK TABLES;


# Dump of table PowerStates
# ------------------------------------------------------------

DROP TABLE IF EXISTS `PowerStates`;

CREATE TABLE `PowerStates` (
  `powerStateID` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(256) DEFAULT NULL,
  `description` text,
  PRIMARY KEY (`powerStateID`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

LOCK TABLES `PowerStates` WRITE;
/*!40000 ALTER TABLE `PowerStates` DISABLE KEYS */;

INSERT INTO `PowerStates` (`powerStateID`, `name`, `description`)
VALUES
	(1,'On',NULL),
	(2,'Standby',NULL),
	(3,'Off',NULL);

/*!40000 ALTER TABLE `PowerStates` ENABLE KEYS */;
UNLOCK TABLES;


# Dump of table RoomConfiguration
# ------------------------------------------------------------

DROP TABLE IF EXISTS `RoomConfiguration`;

CREATE TABLE `RoomConfiguration` (
  `roomConfigurationID` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(256) DEFAULT NULL,
  `description` text,
  `roomConfigurationKey` varchar(256) DEFAULT NULL,
  PRIMARY KEY (`roomConfigurationID`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

LOCK TABLES `RoomConfiguration` WRITE;
/*!40000 ALTER TABLE `RoomConfiguration` DISABLE KEYS */;

INSERT INTO `RoomConfiguration` (`roomConfigurationID`, `name`, `description`, `roomConfigurationKey`)
VALUES
	(1,'Default','The default room configuration','Default');

/*!40000 ALTER TABLE `RoomConfiguration` ENABLE KEYS */;
UNLOCK TABLES;


# Dump of table RoomConfigurationMapping
# ------------------------------------------------------------

DROP TABLE IF EXISTS `RoomConfigurationMapping`;

CREATE TABLE `RoomConfigurationMapping` (
  `roomConfigurationMappingID` int(11) NOT NULL AUTO_INCREMENT,
  `commandID` int(11) DEFAULT NULL,
  `roomConfigurationID` int(11) DEFAULT NULL,
  `commandCodeKey` varchar(256) DEFAULT NULL,
  `priority` int(11) DEFAULT NULL,
  PRIMARY KEY (`roomConfigurationMappingID`),
  KEY `commandID` (`commandID`),
  KEY `roomConfigurationID` (`roomConfigurationID`),
  CONSTRAINT `RoomConfigurationMapping_ibfk_1` FOREIGN KEY (`commandID`) REFERENCES `Commands` (`commandID`),
  CONSTRAINT `RoomConfigurationMapping_ibfk_2` FOREIGN KEY (`roomConfigurationID`) REFERENCES `RoomConfiguration` (`roomConfigurationID`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

LOCK TABLES `RoomConfigurationMapping` WRITE;
/*!40000 ALTER TABLE `RoomConfigurationMapping` DISABLE KEYS */;

INSERT INTO `RoomConfigurationMapping` (`roomConfigurationMappingID`, `commandID`, `roomConfigurationID`, `commandCodeKey`, `priority`)
VALUES
	(2,1,1,'PowerOnDefault',1),
	(3,2,1,'StandbyDefault',9999),
	(4,3,1,'ChangeInputDefault',1337);

/*!40000 ALTER TABLE `RoomConfigurationMapping` ENABLE KEYS */;
UNLOCK TABLES;


# Dump of table Rooms
# ------------------------------------------------------------

DROP TABLE IF EXISTS `Rooms`;

CREATE TABLE `Rooms` (
  `roomID` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(256) DEFAULT NULL,
  `buildingID` int(11) DEFAULT NULL,
  `description` varchar(256) DEFAULT NULL,
  `currentVideoOutputID` int(11) DEFAULT NULL,
  `currentAudioOutputID` int(11) DEFAULT NULL,
  `currentVideoInputID` int(11) DEFAULT NULL,
  `currentAudioInputID` int(11) DEFAULT NULL,
  `configurationID` int(11) DEFAULT NULL,
  PRIMARY KEY (`roomID`),
  KEY `rmBld_ind` (`buildingID`),
  KEY `rmCurVOut_ind` (`currentVideoOutputID`),
  KEY `rmCurVIn_ind` (`currentVideoInputID`),
  KEY `rmCurAOut_ind` (`currentAudioOutputID`),
  KEY `rmCurAIn_ind` (`currentAudioInputID`),
  CONSTRAINT `Rooms_ibfk_1` FOREIGN KEY (`buildingID`) REFERENCES `Buildings` (`buildingID`),
  CONSTRAINT `Rooms_ibfk_2` FOREIGN KEY (`currentVideoOutputID`) REFERENCES `Devices` (`deviceID`),
  CONSTRAINT `Rooms_ibfk_3` FOREIGN KEY (`currentVideoInputID`) REFERENCES `Devices` (`deviceID`),
  CONSTRAINT `Rooms_ibfk_4` FOREIGN KEY (`currentAudioOutputID`) REFERENCES `Devices` (`deviceID`),
  CONSTRAINT `Rooms_ibfk_5` FOREIGN KEY (`currentAudioInputID`) REFERENCES `Devices` (`deviceID`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

LOCK TABLES `Rooms` WRITE;
/*!40000 ALTER TABLE `Rooms` DISABLE KEYS */;

INSERT INTO `Rooms` (`roomID`, `name`, `buildingID`, `description`, `currentVideoOutputID`, `currentAudioOutputID`, `currentVideoInputID`, `currentAudioInputID`, `configurationID`)
VALUES
	(1,'1110',1,'Nyle\'s office',1,1,1,1,1),
	(2,'1001D',1,'Joe/Jesse Cube',1,1,1,1,1);

/*!40000 ALTER TABLE `Rooms` ENABLE KEYS */;
UNLOCK TABLES;


# Dump of table vConfigurationMapping
# ------------------------------------------------------------

DROP VIEW IF EXISTS `vConfigurationMapping`;

CREATE TABLE `vConfigurationMapping` (
   `ConfigurationID` INT(11) NOT NULL DEFAULT '0',
   `ConfigurationName` VARCHAR(256) NULL DEFAULT NULL,
   `CodeKey` VARCHAR(256) NULL DEFAULT NULL,
   `Priority` INT(11) NULL DEFAULT NULL,
   `ConfigurationKey` VARCHAR(256) NULL DEFAULT NULL
) ENGINE=MyISAM;





# Replace placeholder table for vConfigurationMapping with correct view syntax
# ------------------------------------------------------------

DROP TABLE `vConfigurationMapping`;

CREATE ALGORITHM=UNDEFINED DEFINER=`root`@`%` SQL SECURITY DEFINER VIEW `vConfigurationMapping`
AS SELECT
   `rc`.`roomConfigurationID` AS `ConfigurationID`,
   `rc`.`name` AS `ConfigurationName`,
   `rcm`.`commandCodeKey` AS `CodeKey`,
   `rcm`.`priority` AS `Priority`,
   `rc`.`roomConfigurationKey` AS `ConfigurationKey`
FROM (`RoomConfiguration` `rc` join `RoomConfigurationMapping` `rcm` on((`rc`.`roomConfigurationID` = `rcm`.`roomConfigurationID`)));

--
-- Dumping routines (PROCEDURE) for database 'configuration'
--
DELIMITER ;;

# Dump of PROCEDURE AddConfigurationMapping
# ------------------------------------------------------------

/*!50003 DROP PROCEDURE IF EXISTS `AddConfigurationMapping` */;;
/*!50003 SET SESSION SQL_MODE=""*/;;
/*!50003 CREATE*/ /*!50020 DEFINER=`root`@`%`*/ /*!50003 PROCEDURE `AddConfigurationMapping`(IN ConfigurationName VARCHAR(256), IN CommandKey VARCHAR(256), IN Priority INT)
BEGIN
	
	DECLARE rcid int;
	DECLARE cid int;
	
	SELECT roomConfigurationID into rcid 
	FROM RoomConfiguration WHERE RoomConfiguration.name = ConfigurationName;

	Insert into RoomConfigurationMapping 
	(roomConfigurationID, commandCodeKey,priority)
	VALUES 
	(rcid, CommandKey, Priority);

END */;;

/*!50003 SET SESSION SQL_MODE=@OLD_SQL_MODE */;;
DELIMITER ;

/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;
/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
