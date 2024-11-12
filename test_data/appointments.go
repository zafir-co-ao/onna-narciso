package testdata

import "github.com/zafir-co-ao/onna-narciso/internal/scheduling"

var Appointments = []scheduling.Appointment{

	{
		ID:               "1",
		CustomerID:       "1",
		CustomerName:     "Paola Oliveira",
		ProfessionalID:   Professionals[0].ID,
		ProfessionalName: Professionals[0].Name,
		ServiceID:        Services[0].ID,
		ServiceName:      Services[0].Name,
		Date:             "2024-11-12",
		Status:           scheduling.StatusScheduled,
		Hour:             "08:00",
		Duration:         180,
	},
	{
		ID:               "2",
		CustomerID:       "2",
		CustomerName:     "Juliana Paes",
		ProfessionalID:   Professionals[1].ID,
		ProfessionalName: Professionals[1].Name,
		ServiceID:        Services[2].ID,
		ServiceName:      Services[2].Name,
		Status:           scheduling.StatusScheduled,
		Date:             "2024-11-12",
		Hour:             "10:30",
		Duration:         90,
	},
	{
		ID:               "3",
		CustomerID:       "3",
		CustomerName:     "Gisele Bündchen",
		ProfessionalID:   Professionals[2].ID,
		ProfessionalName: Professionals[2].Name,
		ServiceID:        Services[3].ID,
		ServiceName:      Services[3].Name,
		Status:           scheduling.StatusScheduled,
		Date:             "2024-11-12",
		Hour:             "12:00",
		Duration:         60,
	},
	{
		ID:               "4",
		CustomerID:       "2",
		CustomerName:     "Juliana Paes",
		ProfessionalID:   Professionals[2].ID,
		ProfessionalName: Professionals[2].Name,
		ServiceID:        Services[2].ID,
		ServiceName:      Services[2].Name,
		Status:           scheduling.StatusClosed,
		Date:             "2024-11-12",
		Hour:             "10:30",
		Duration:         90,
	},
}
