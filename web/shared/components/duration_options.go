package components

var AppointmentDurations = []InputOption{
	{"15 Minutos", "15"},
	{"30 Minutos", "30"},
	{"45 Minutos", "45"},
	{"1 Hora", "60"},
	{"1 Hora e 30 Minutos", "90"},
	{"2 Horas", "120"},
	{"2 Horas e 30 Minutos", "150"},
	{"3 Horas", "180"},
}

func FindDuration(v string) InputOption {
	for _, d := range AppointmentDurations {
		if d[1] == v {
			return d
		}
	}

	return AppointmentDurations[4]
}
