package models

// PrintFileMetadata - Print file metadata parsed from G-code or SL1, all data are optional
type PrintFileMetadata struct {

	// Degrees Celsius
	BedTemperature int32 `json:"bed_temperature,omitempty"`

	BedTemperaturePerTool []int32 `json:"bed_temperature per tool,omitempty"`

	// Nozzle temperature, Degrees Celsius
	Temperature int32 `json:"temperature,omitempty"`

	TemperaturePerTool []int32 `json:"temperature per tool,omitempty"`

	// Millimeters
	BrimWidth int32 `json:"brim_width,omitempty"`

	EstimatedPrintingTimeNormalMode string `json:"estimated printing time (normal mode),omitempty"`

	// Seconds
	EstimatedPrintTime int32 `json:"estimated_print_time,omitempty"`

	FadedLayers int32 `json:"faded_layers,omitempty"`

	FilamentCost float32 `json:"filament cost,omitempty"`

	FilamentCostPerTool []float32 `json:"filament cost per tool,omitempty"`

	FilamentUsedCm3 float32 `json:"filament used [cm3],omitempty"`

	FilamentUsedCm3PerTool []float32 `json:"filament used [cm3] per tool,omitempty"`

	FilamentUsedG float32 `json:"filament used [g],omitempty"`

	FilamentUsedGPerTool []float32 `json:"filament used [g] per tool,omitempty"`

	FilamentUsedMm float32 `json:"filament used [mm],omitempty"`

	FilamentUsedMmPerTool []float32 `json:"filament used [mm] per tool,omitempty"`

	FilamentType string `json:"filament_type,omitempty"`

	FilamentTypePerTool []string `json:"filament_type per tool,omitempty"`

	// Percents
	FillDensity string `json:"fill_density,omitempty"`

	// Seconds
	InitialExposureTime int32 `json:"initial_exposure_time,omitempty"`

	// Millimeters
	LayerHeight float32 `json:"layer_height,omitempty"`

	MaterialName string `json:"material_name,omitempty"`

	// Seconds
	ExposureTime int32 `json:"exposure_time,omitempty"`

	// Seconds
	MaxExposureTime int32 `json:"max_exposure_time,omitempty"`

	// Seconds
	MaxInitialExposureTime int32 `json:"max_initial_exposure_time,omitempty"`

	// Seconds
	MinExposureTime int32 `json:"min_exposure_time,omitempty"`

	// Seconds
	MinInitialExposureTime int32 `json:"min_initial_exposure_time,omitempty"`

	// Millimeters
	NozzleDiameter float32 `json:"nozzle_diameter,omitempty"`

	NozzleDiameterPerTool []float32 `json:"nozzle_diameter per tool,omitempty"`

	NormalPercentPresent bool `json:"normal_percent_present,omitempty"`

	NormalLeftPresent bool `json:"normal_left_present,omitempty"`

	QuietPercentPresent bool `json:"quiet_percent_present,omitempty"`

	QuietLeftPresent bool `json:"quiet_left_present,omitempty"`

	LayerInfoPresent bool `json:"layer_info_present,omitempty"`

	MaxLayerZ float32 `json:"max_layer_z,omitempty"`

	// Seconds
	PrintTime int32 `json:"print_time,omitempty"`

	PrinterModel string `json:"printer_model,omitempty"`

	SupportMaterial string `json:"support_material,omitempty"`

	Ironing int32 `json:"ironing,omitempty"`

	RequiredResinMl float32 `json:"required_resin_ml,omitempty"`

	Profile string `json:"profile,omitempty"`
}
