function getPosition(event) {
    const calendar = event.target;
    const col = calendar.clientWidth / 5;
    const deslocX = Math.trunc((event.layerX) / col);
    return deslocX + 1;
}
