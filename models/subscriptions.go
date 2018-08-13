package models

//    this.belongsTo(models['target'], { as: 'source', foreignKey: 'sourceId' });
// this.belongsTo(models['target'], { as: 'target', foreignKey: 'targetId' });
// this.belongsTo(models['targetDestinations'], { as: 'destination', foreignKey: 'destinationId' });
// this.belongsTo(models['event'], { as: 'event', foreignKey: 'eventId' });
// this.belongsTo(models['endpoint'], { as: 'endpoint', foreignKey: 'endpointId' });
// this.belongsTo(models['channel'], { as: 'channel', foreignKey: 'channelId' });
// this.belongsTo(models['redoxSource'], { as: 'redoxSource', foreignKey: 'targetId' });

type Subscriptions struct {
	ID            uint  `gorm:"column:id"`
	SourceID      uint  `gorm:"column:sourceId"`
	TargetID      uint  `gorm:"column:targetId"`
	DestinationID uint  `gorm:"column:destinationId"`
	EventID       uint  `gorm:"column:eventId"`
	DataModelID   uint  `gorm:"column:dataModelId"`
	EndpointID    int64 `gorm:"column:endpointId"`

	Source            Targets            `gorm:"foreignkey:SourceID"`
	Target            Targets            `gorm:"foreignkey:TargetID"`
	DataModel         DataModels         `gorm:"foreignkey:DataModelID"`
	Event             Events             `gorm:"foreignkey:EventID"`
	TargetDestination TargetDestinations `gorm:"foreignkey:DestinationID"`
	Endpoint          Endpoints          `gorm:"foreignkey:EndpointID"`
}
