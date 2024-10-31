function openScheduleForm(event, rows) {
    // hidden("#error-messages")
    // hidden("#customer-form")

    const week = document.querySelectorAll(".week-day-cell")
    const day = week[getPosition(event) - 1]
    const date = day.getAttribute("data-week-day")
    const hour = getHour(event, rows)

    document.querySelector("#event-date").setAttribute("value", date)
    document.querySelector("#event-hour").setAttribute("value", hour)
    document.querySelector("#event-service").setAttribute("value", 4)
    document.querySelector("#event-professional").setAttribute("value", 1)

    document.querySelector("#professional").innerHTML = "Sara Gomes"
    document.querySelector("#service").innerHTML = "Manicure + Pedicure"
    document.querySelector("#date").innerHTML = date
    document.querySelector("#hour").innerHTML = hour

    // openDialog("#schedule-dialog")
}

function openRescheduleForm(event, id) {
    event.stopPropagation()
    openDialog("#reschedule-dialog")
    htmx.ajax("GET", `/appointments/${id}`, {target: "#reschedule-form", swap: "innerHTML"})
}

function cancelEvent() {
    const id = document.querySelector("#appointment-id").getAttribute("value")

    if (!id) {
        console.error("Erro ao cancelar evento")
        return
    }

    htmx.ajax("POST", `/appointments/${id}/cancel`, { target: `#${id}`, swap: "outerHTML"})


    closeDialog("#cancel-dialog")
    closeDialog("#reschedule-dialog")
}

function closeScheduleForm() {
    show("#select-customer")
    closeDialog('#schedule-dialog')
}
