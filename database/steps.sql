INSERT INTO Commands (name, description) VALUES
('PowerOn', 'Pull out of standby.'),
('PowerOff', 'Put into standby.'),
('ChangeInput', 'Change the input to the supplied port (hdmi 1, hdmi 2) etc. etc.)'),
('SetVolume', 'Change the volume to supplied value'),
('VolumeDown', 'Tick volume down.'),
('VolumeUp', 'Tick volume up.'),
('BlankScreen', 'Blanks the screen'),
('UnblankScreen', 'Unblanks the screen');

INSERT INTO Endpoints (name, path, description) VALUES
('PowerOn', '/:address/power/on' ,'Standard PowerOn endpoint.'),
('PowerOff', '/:address/power/off' ,'Standard PowerOff endpoint.'),
('ChangeInput','/:address/input/:port', 'Standard ChangeInput endpoint.'),
('SetVolume', '/:address/volume/:level', 'Standard SetVolume endpoint.'),
('VolumeUp', '/:address/volume/Up', 'Standard SetVolume endpoint.'),
('VolumeDown', '/:address/volume/Down', 'Standard SetVolume endpoint.'),
('BlankScreen', '/:address/display/blank', 'Standard BlankScreen endpoint.'),
('UnblankScreen','/:address/display/unblank', 'Standard UnblankScreen endpoint.');

INSERT INTO Microservices (Address, name) VALUES
('localhost:8005', 'pjlink-microservice'),
('localhost:8007', 'sony-control-microservice');

INSERT INTO Buildings (name, shortName) VALUES
('Information Technology Buildilng', 'ITB');

INSERT INTO Rooms (buildingID, name, description) VALUES
((SELECT buildingID FROM Buildings WHERE shortName = 'ITB'), '1110', 'Nyle\'s office');

INSERT INTO DeviceTypes (name, description) VALUES
('tv', 'A TV'),
('comptuer', 'A Comptuer input'),
('roku', 'A Roku device'),
('iptv', 'An IPTV settop box');

INSERT INTO Devices (name, address, input, output, typeID, roomID, buildingID) VALUES
('Nyle\'s Computer', NULL, 1, 0, 2, 1, 1),
('Nyle\'s Roku', NULL, 1,0,2,1,1),
('Nyle\'s IPTV', NULL, 1,0,3,1,1),
('Nyle\'s TV', '10.66.9.6', 0, 1, 1, 1, 1);

INSERT INTO DeviceRoleDefinition (name, description) VALUES
('AudioOut', 'Device that outputs audio (speakers, tv, etc.)'),
('VideoOut', 'Device that displays video (projector, tv, etc.)');

INSERT INTO DeviceRole (deviceID, deviceRoleDefinitionID) VALUES
(1,1),
(1,2);

INSERT INTO Ports(port) VALUES
('Hdmi1'),
('Hdmi2'),
('Hdmi3'),
('Hdmi4');

INSERT INTO PortConfiguration (destinationDeviceID, portID, sourceDeviceID) VALUES
(1,1,2),
(1,2,4),
(1,3,3);

INSERT INTO DeviceCommands  (commandID, deviceID, enabled, endpointID, microserviceID) VALUES
(1, 1, 1, 1, 2),
(2, 1, 1, 2, 2),
(3, 1, 1, 3, 2),
(4, 1, 1, 4, 2),
(5, 1, 1, 5, 2),
(6, 1, 1, 6, 2),
(7, 1, 1, 8, 2),
(8, 1, 1, 7, 2);

-------------------------------
--Where we are
-------------------------------

-- Get all commands/microservice/endpoints for output devices in a given room. 
SELECT
Devices.name as deviceName,
Devices.address as deviceAddress,
Commands.name as commandName,
Endpoints.name as endpointName,
Endpoints.path as endpointPath,
Microservices.address as microserviceAddress
 FROM Devices JOIN DeviceCommands on Devices.deviceID = DeviceCommands.deviceID JOIN Commands on DeviceCommands.commandID = Commands.commandID JOIN Endpoints on DeviceCommands.endpointID = Endpoints.endpointID JOIN Microservices ON DeviceCommands.microserviceID = Microservices.microserviceID
 WHERE Devices.RoomID = 1 AND Devices.output = 1
