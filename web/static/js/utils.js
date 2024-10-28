
function openDialog(target) {
    document.querySelector(target).showModal()
}

function closeDialog(target) {
    document.querySelector(target).close()
}

function getHour(event, rows) {
    const h = calculateHour(event, rows)
    const m = calculateMinutes(event, rows)
    if (m.toString().length === 1 ) {
        return `${h}:0${m}`
    }

    return `${h}:${m}`
}

function calculateHour(event, rows) {
    const calendar = event.target;
    const row = calendar.clientHeight / parseInt(rows);
    const deslocY = Math.trunc((event.layerY - (3 * row)) / row);
    const h = Math.trunc(deslocY / 4 + 8);
    return h
}

function calculateMinutes(event, rows) {
    const calendar = event.target;
    const row = calendar.clientHeight / parseInt(rows);
    const deslocY = Math.trunc((event.layerY - (3 * row)) / row);
    const m = Math.trunc((deslocY % 4) * 15);
    return m
}

function getPosition(event) {
    const calendar = event.target;
    const col = calendar.clientWidth / 6;
    const deslocX = Math.trunc((event.layerX) / col);
    return deslocX + 1;
}

function toogle(target) {
    const el = document.querySelector(target)
    if (el.classList.contains("hidden")) {
        el.classList.remove("hidden");
        el.classList.add("block");
        return
    }

    el.classList.add("hidden");
    el.classList.remove("block");
}


function show(target) {
    const el = document.querySelector(target)
    if (el.classList.contains("block")) return

    el.classList.remove("hidden");
    el.classList.add("block");
}

function hidden(target) {
    const el = document.querySelector(target)
    if (el.classList.contains("hidden")) return

    el.classList.add("hidden");
    el.classList.remove("block");
}
