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

const customers = [
    {
        "id": "2b6d51d8-0486-4c0d-a38b-07349f0f2f57",
        "name": "João Silva",
        "nif": "500012345",
        "phoneNumber": "+244900000001",
        "email": "joao1@exemplo.com"
    },
    {
        "id": "2bc5bf19-0a06-4571-b59c-2f6856981138",
        "name": "Maria Oliveira",
        "nif": "500022345",
        "phoneNumber": "+244900000002",
        "email": "maria2@exemplo.com"
    },
    {
        "id": "ef0b7a28-734f-4e82-b158-5b10c8b19198",
        "name": "Carlos Ferreira",
        "nif": "500032345",
        "phoneNumber": "+244900000003",
        "email": "carlos3@exemplo.com"
    },
    {
        "id": "f3f641e0-b138-4338-b929-9127daf77e3b",
        "name": "Ana Costa",
        "nif": "500042345",
        "phoneNumber": "+244900000004",
        "email": "ana4@exemplo.com"
    },
    {
        "id": "854cc9df-e857-4325-96a4-8b0a16d6ac01",
        "name": "Pedro Santos",
        "nif": "500052345",
        "phoneNumber": "+244900000005",
        "email": "pedro5@exemplo.com"
    },
    {
        "id": "e1fc29da-2826-4413-a7a3-48318d0a13f0",
        "name": "Luísa Lima",
        "nif": "500062345",
        "phoneNumber": "+244900000006",
        "email": "lu\u00edsa6@exemplo.com"
    },
    {
        "id": "24e047fd-d47a-4902-9dad-15b81bc2ab61",
        "name": "José Pereira",
        "nif": "500072345",
        "phoneNumber": "+244900000007",
        "email": "jos\u00e97@exemplo.com"
    },
    {
        "id": "67111fb3-4356-4d63-a152-709ae520eaa5",
        "name": "Rita Almeida",
        "nif": "500082345",
        "phoneNumber": "+244900000008",
        "email": "rita8@exemplo.com"
    },
    {
        "id": "772dcd33-eccf-44f4-ac75-87e0240aaa36",
        "name": "Manuel Carvalho",
        "nif": "500092345",
        "phoneNumber": "+244900000009",
        "email": "manuel9@exemplo.com"
    },
    {
        "id": "20cdf3b3-0251-4584-85d1-ccbb3900d1cb",
        "name": "Sofia Fernandes",
        "nif": "5000102345",
        "phoneNumber": "+2449000000010",
        "email": "sofia10@exemplo.com"
    }
]

function toogleCustomers() {
    const el  = document.querySelector("#menu-customer")
    if (el.classList.contains("hidden")) {
        el.classList.remove("hidden");
        el.classList.add("block");
        return
    }

    el.classList.add("hidden");
    el.classList.remove("block");
}

function searchCustomers(el) {
    const container = document.querySelector("#customers-list");

    if (el.value.length === 0) {
        container.innerHTML = `<li class="text-sm mt-2 border-b p-2 cursor-pointer" onclick="toogleClientCreateForm()">Criar Novo Cliente</li>`
        return;
    }

    if (el.value.length < 3) {
        return;
    }

    let filteredCustomers = []

    for (const customer of customers) {
        if (customer.name.toLowerCase().includes(el.value.toLowerCase()) ||
            customer.nif.includes(el.value) ||
            customer.phoneNumber.includes(el.value)) {
            filteredCustomers.push(customer)
        }
    }

    container.innerHTML = "";

    if (filteredCustomers.length === 0) {
        return
    }

    for (const customer of filteredCustomers) {
        const li = document.createElement("li")
        li.className = "text-sm mt-3 border-b p-2 cursor-pointer"
        li.innerText = customer.name
        li.addEventListener("click", () => selectCustomer(customer.id, customer.name))
        container.appendChild(li)
    }

}

function selectCustomer(id, name) {
    const customerNameEl = document.querySelector("#customer-name");
    const customerIdEl = document.querySelector("#customer-id")

    customerNameEl.innerHTML = name;
    customerIdEl.value = id
    toogleCustomers();
}

function getPosition(event) {
    const calendar = event.target;
    const col = calendar.clientWidth / 6;
    const deslocX = Math.trunc((event.layerX) / col);
    return deslocX + 1;
}
