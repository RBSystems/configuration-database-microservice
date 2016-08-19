-- phpMyAdmin SQL Dump
-- version 4.6.3
-- https://www.phpmyadmin.net/
--
-- Host: db
-- Generation Time: Aug 18, 2016 at 11:03 PM
-- Server version: 10.1.14-MariaDB-1~jessie
-- PHP Version: 5.6.21

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Database: `configuration`
--

-- --------------------------------------------------------

--
-- Table structure for table `AudioDevice`
--

CREATE TABLE `AudioDevice` (
  `audioDeviceID` int(11) NOT NULL,
  `deviceID` int(11) DEFAULT NULL,
  `deviceRoleDefinitionID` int(11) DEFAULT NULL,
  `muted` tinyint(1) DEFAULT NULL,
  `volume` int(11) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

-- --------------------------------------------------------

--
-- Table structure for table `Buildings`
--

CREATE TABLE `Buildings` (
  `buildingID` int(11) NOT NULL,
  `name` text,
  `shortName` varchar(256) DEFAULT NULL,
  `description` text
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

-- --------------------------------------------------------

--
-- Table structure for table `Commands`
--

CREATE TABLE `Commands` (
  `commandID` int(11) NOT NULL,
  `name` varchar(256) DEFAULT NULL,
  `description` text
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

-- --------------------------------------------------------

--
-- Table structure for table `DeviceCommands`
--

CREATE TABLE `DeviceCommands` (
  `deviceCommandID` int(11) NOT NULL,
  `deviceID` int(11) DEFAULT NULL,
  `commandID` int(11) DEFAULT NULL,
  `microserviceID` int(11) DEFAULT NULL,
  `endpointID` int(11) DEFAULT NULL,
  `enabled` tinyint(1) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

-- --------------------------------------------------------

--
-- Table structure for table `DeviceRole`
--

CREATE TABLE `DeviceRole` (
  `deviceRoleID` int(11) NOT NULL,
  `deviceID` int(11) DEFAULT NULL,
  `deviceRoleDefinitionID` int(11) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

-- --------------------------------------------------------

--
-- Table structure for table `DeviceRoleDefinition`
--

CREATE TABLE `DeviceRoleDefinition` (
  `deviceRoleDefinitionID` int(11) NOT NULL,
  `name` varchar(256) DEFAULT NULL,
  `description` text
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

-- --------------------------------------------------------

--
-- Table structure for table `Devices`
--

CREATE TABLE `Devices` (
  `deviceID` int(11) NOT NULL,
  `name` varchar(256) DEFAULT NULL,
  `address` varchar(256) DEFAULT NULL,
  `buildingID` int(11) DEFAULT NULL,
  `roomID` int(11) DEFAULT NULL,
  `typeID` int(11) DEFAULT NULL,
  `powerID` int(11) DEFAULT NULL,
  `responding` tinyint(1) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

-- --------------------------------------------------------

--
-- Table structure for table `DeviceTypes`
--

CREATE TABLE `DeviceTypes` (
  `deviceTypeID` int(11) NOT NULL,
  `name` varchar(256) DEFAULT NULL,
  `description` text
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

-- --------------------------------------------------------

--
-- Table structure for table `Displays`
--

CREATE TABLE `Displays` (
  `displayID` int(11) NOT NULL,
  `deviceID` int(11) DEFAULT NULL,
  `deviceRoleDefinitionID` int(11) DEFAULT NULL,
  `blanked` tinyint(1) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

-- --------------------------------------------------------

--
-- Table structure for table `Endpoints`
--

CREATE TABLE `Endpoints` (
  `endpointID` int(11) NOT NULL,
  `name` varchar(256) DEFAULT NULL,
  `path` text,
  `description` text
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

-- --------------------------------------------------------

--
-- Table structure for table `Microservices`
--

CREATE TABLE `Microservices` (
  `microserviceID` int(11) NOT NULL,
  `name` varchar(256) DEFAULT NULL,
  `address` text,
  `description` text
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

-- --------------------------------------------------------

--
-- Table structure for table `PowerState`
--

CREATE TABLE `PowerState` (
  `powerStateID` int(11) NOT NULL,
  `name` varchar(256) DEFAULT NULL,
  `description` text
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

-- --------------------------------------------------------

--
-- Table structure for table `Rooms`
--

CREATE TABLE `Rooms` (
  `roomID` int(11) NOT NULL,
  `name` varchar(256) DEFAULT NULL,
  `buildingID` int(11) DEFAULT NULL,
  `description` varchar(256) DEFAULT NULL,
  `currentVideoOutputID` int(11) DEFAULT NULL,
  `currentAudioOutputID` int(11) DEFAULT NULL,
  `currentVideoInputID` int(11) DEFAULT NULL,
  `currentAudioInputID` int(11) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

--
-- Indexes for dumped tables
--

--
-- Indexes for table `AudioDevice`
--
ALTER TABLE `AudioDevice`
  ADD PRIMARY KEY (`audioDeviceID`),
  ADD KEY `audDev_ind` (`deviceID`),
  ADD KEY `audDevRol_ind` (`deviceRoleDefinitionID`);

--
-- Indexes for table `Buildings`
--
ALTER TABLE `Buildings`
  ADD PRIMARY KEY (`buildingID`);

--
-- Indexes for table `Commands`
--
ALTER TABLE `Commands`
  ADD PRIMARY KEY (`commandID`);

--
-- Indexes for table `DeviceCommands`
--
ALTER TABLE `DeviceCommands`
  ADD PRIMARY KEY (`deviceCommandID`),
  ADD KEY `devComDev_ind` (`deviceID`),
  ADD KEY `devComCom_ind` (`commandID`),
  ADD KEY `devComMS_ind` (`microserviceID`),
  ADD KEY `devComEnd_ind` (`endpointID`);

--
-- Indexes for table `DeviceRole`
--
ALTER TABLE `DeviceRole`
  ADD PRIMARY KEY (`deviceRoleID`),
  ADD KEY `devRolDevID_ind` (`deviceID`),
  ADD KEY `devRolDevRolDef_ind` (`deviceRoleDefinitionID`);

--
-- Indexes for table `DeviceRoleDefinition`
--
ALTER TABLE `DeviceRoleDefinition`
  ADD PRIMARY KEY (`deviceRoleDefinitionID`);

--
-- Indexes for table `Devices`
--
ALTER TABLE `Devices`
  ADD PRIMARY KEY (`deviceID`),
  ADD KEY `devbld_ind` (`buildingID`),
  ADD KEY `devrm_ind` (`roomID`),
  ADD KEY `devty_ind` (`typeID`),
  ADD KEY `devpw_ind` (`powerID`);

--
-- Indexes for table `DeviceTypes`
--
ALTER TABLE `DeviceTypes`
  ADD PRIMARY KEY (`deviceTypeID`);

--
-- Indexes for table `Displays`
--
ALTER TABLE `Displays`
  ADD PRIMARY KEY (`displayID`),
  ADD KEY `dispDev_ind` (`deviceID`),
  ADD KEY `dispDevRol_ind` (`deviceRoleDefinitionID`);

--
-- Indexes for table `Endpoints`
--
ALTER TABLE `Endpoints`
  ADD PRIMARY KEY (`endpointID`);

--
-- Indexes for table `Microservices`
--
ALTER TABLE `Microservices`
  ADD PRIMARY KEY (`microserviceID`);

--
-- Indexes for table `PowerState`
--
ALTER TABLE `PowerState`
  ADD PRIMARY KEY (`powerStateID`);

--
-- Indexes for table `Rooms`
--
ALTER TABLE `Rooms`
  ADD PRIMARY KEY (`roomID`),
  ADD KEY `rmBld_ind` (`buildingID`),
  ADD KEY `rmCurVOut_ind` (`currentVideoOutputID`),
  ADD KEY `rmCurVIn_ind` (`currentVideoInputID`),
  ADD KEY `rmCurAOut_ind` (`currentAudioOutputID`),
  ADD KEY `rmCurAIn_ind` (`currentAudioInputID`);

--
-- AUTO_INCREMENT for dumped tables
--

--
-- AUTO_INCREMENT for table `AudioDevice`
--
ALTER TABLE `AudioDevice`
  MODIFY `audioDeviceID` int(11) NOT NULL AUTO_INCREMENT;
--
-- AUTO_INCREMENT for table `Buildings`
--
ALTER TABLE `Buildings`
  MODIFY `buildingID` int(11) NOT NULL AUTO_INCREMENT;
--
-- AUTO_INCREMENT for table `Commands`
--
ALTER TABLE `Commands`
  MODIFY `commandID` int(11) NOT NULL AUTO_INCREMENT;
--
-- AUTO_INCREMENT for table `DeviceCommands`
--
ALTER TABLE `DeviceCommands`
  MODIFY `deviceCommandID` int(11) NOT NULL AUTO_INCREMENT;
--
-- AUTO_INCREMENT for table `DeviceRole`
--
ALTER TABLE `DeviceRole`
  MODIFY `deviceRoleID` int(11) NOT NULL AUTO_INCREMENT;
--
-- AUTO_INCREMENT for table `DeviceRoleDefinition`
--
ALTER TABLE `DeviceRoleDefinition`
  MODIFY `deviceRoleDefinitionID` int(11) NOT NULL AUTO_INCREMENT;
--
-- AUTO_INCREMENT for table `Devices`
--
ALTER TABLE `Devices`
  MODIFY `deviceID` int(11) NOT NULL AUTO_INCREMENT;
--
-- AUTO_INCREMENT for table `DeviceTypes`
--
ALTER TABLE `DeviceTypes`
  MODIFY `deviceTypeID` int(11) NOT NULL AUTO_INCREMENT;
--
-- AUTO_INCREMENT for table `Displays`
--
ALTER TABLE `Displays`
  MODIFY `displayID` int(11) NOT NULL AUTO_INCREMENT;
--
-- AUTO_INCREMENT for table `Endpoints`
--
ALTER TABLE `Endpoints`
  MODIFY `endpointID` int(11) NOT NULL AUTO_INCREMENT;
--
-- AUTO_INCREMENT for table `Microservices`
--
ALTER TABLE `Microservices`
  MODIFY `microserviceID` int(11) NOT NULL AUTO_INCREMENT;
--
-- AUTO_INCREMENT for table `PowerState`
--
ALTER TABLE `PowerState`
  MODIFY `powerStateID` int(11) NOT NULL AUTO_INCREMENT;
--
-- AUTO_INCREMENT for table `Rooms`
--
ALTER TABLE `Rooms`
  MODIFY `roomID` int(11) NOT NULL AUTO_INCREMENT;
--
-- Constraints for dumped tables
--

--
-- Constraints for table `AudioDevice`
--
ALTER TABLE `AudioDevice`
  ADD CONSTRAINT `AudioDevice_ibfk_1` FOREIGN KEY (`deviceID`) REFERENCES `Devices` (`deviceID`),
  ADD CONSTRAINT `AudioDevice_ibfk_2` FOREIGN KEY (`deviceRoleDefinitionID`) REFERENCES `DeviceRoleDefinition` (`deviceRoleDefinitionID`);

--
-- Constraints for table `DeviceCommands`
--
ALTER TABLE `DeviceCommands`
  ADD CONSTRAINT `DeviceCommands_ibfk_1` FOREIGN KEY (`deviceID`) REFERENCES `Devices` (`deviceID`),
  ADD CONSTRAINT `DeviceCommands_ibfk_2` FOREIGN KEY (`commandID`) REFERENCES `Commands` (`commandID`),
  ADD CONSTRAINT `DeviceCommands_ibfk_3` FOREIGN KEY (`endpointID`) REFERENCES `Endpoints` (`endpointID`),
  ADD CONSTRAINT `DeviceCommands_ibfk_4` FOREIGN KEY (`microserviceID`) REFERENCES `Microservices` (`microserviceID`);

--
-- Constraints for table `DeviceRole`
--
ALTER TABLE `DeviceRole`
  ADD CONSTRAINT `DeviceRole_ibfk_1` FOREIGN KEY (`deviceID`) REFERENCES `Devices` (`deviceID`),
  ADD CONSTRAINT `DeviceRole_ibfk_2` FOREIGN KEY (`deviceRoleDefinitionID`) REFERENCES `DeviceRoleDefinition` (`deviceRoleDefinitionID`);

--
-- Constraints for table `Devices`
--
ALTER TABLE `Devices`
  ADD CONSTRAINT `Devices_ibfk_1` FOREIGN KEY (`buildingID`) REFERENCES `Buildings` (`buildingID`),
  ADD CONSTRAINT `Devices_ibfk_2` FOREIGN KEY (`typeID`) REFERENCES `DeviceTypes` (`deviceTypeID`),
  ADD CONSTRAINT `Devices_ibfk_3` FOREIGN KEY (`powerID`) REFERENCES `PowerState` (`powerStateID`),
  ADD CONSTRAINT `Devices_ibfk_4` FOREIGN KEY (`roomID`) REFERENCES `Rooms` (`roomID`);

--
-- Constraints for table `Displays`
--
ALTER TABLE `Displays`
  ADD CONSTRAINT `Displays_ibfk_1` FOREIGN KEY (`deviceID`) REFERENCES `Devices` (`deviceID`),
  ADD CONSTRAINT `Displays_ibfk_2` FOREIGN KEY (`deviceRoleDefinitionID`) REFERENCES `DeviceRoleDefinition` (`deviceRoleDefinitionID`);

--
-- Constraints for table `Rooms`
--
ALTER TABLE `Rooms`
  ADD CONSTRAINT `Rooms_ibfk_1` FOREIGN KEY (`buildingID`) REFERENCES `Buildings` (`buildingID`),
  ADD CONSTRAINT `Rooms_ibfk_2` FOREIGN KEY (`currentVideoOutputID`) REFERENCES `Devices` (`deviceID`),
  ADD CONSTRAINT `Rooms_ibfk_3` FOREIGN KEY (`currentVideoInputID`) REFERENCES `Devices` (`deviceID`),
  ADD CONSTRAINT `Rooms_ibfk_4` FOREIGN KEY (`currentAudioOutputID`) REFERENCES `Devices` (`deviceID`),
  ADD CONSTRAINT `Rooms_ibfk_5` FOREIGN KEY (`currentAudioInputID`) REFERENCES `Devices` (`deviceID`);

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
