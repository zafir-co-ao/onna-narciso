function openEventDialog(event) {
    const week = document.querySelectorAll(".week-day-cell")

    const day = week[getPosition(event) - 1]
    const date = day.getAttribute("data-week-day")

    document.querySelector("#event-date").setAttribute("value", date)

    const dialogEl = document.querySelector("#add-event-dialog")
    dialogEl.showModal()
}

function closeEventDialog() {
    const dialogEl = document.querySelector("#add-event-dialog")
    dialogEl.close()
}

function toogleCustomers() {}

function getPosition(event) {
    const calendar = event.target;
    const col = calendar.clientWidth / 6;
    const deslocX = Math.trunc((event.layerX) / col);
    return deslocX + 1;
}
