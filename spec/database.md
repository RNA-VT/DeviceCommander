# Device Commander - Database Spec

## Technologies

- gorm
- SQLite

## Models

All models inherit `gorm.Model` :

``` golang
// gorm.Model definition
type Model struct {
  ID        uint           `gorm:"primaryKey"`
  CreatedAt time.Time
  UpdatedAt time.Time
  DeletedAt gorm.DeletedAt `gorm:"index"`
}
```

### Device

- Serial No String
- Name, String
- Board, String
- FirstRegistration, time.Time

### Configuration

Options:

1. Store arbitrary jsons as strings for all configs & query the sub json if we need to using sqlite's `json1`.
2. Model each config in go and let gorm make a table for each type.

- DeviceID, Foreign Key to Device.ID
- ConfigurationJSON String
- ConfigurationApplied Bool

### Events

Options:

1. 1 log table
   - 1 enum defining event types
   - dump to file and truncate table periodically
2. many log tables
   - 1 Model per EventType
   - dump or just truncate noisy tables on a case by case basis

- DeviceID uint
- Type
- Message

### Layout

- Name String
- LayoutConfig String

