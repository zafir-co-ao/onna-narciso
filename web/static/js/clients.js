const customers = [
    {
        "id": "1",
        "name": "João Silva",
        "nif": "500012345",
        "phoneNumber": "+244900000001",
        "email": "joao1@exemplo.com"
    },
    {
        "id": "2",
        "name": "Maria Oliveira",
        "nif": "500022345",
        "phoneNumber": "+244900000002",
        "email": "maria2@exemplo.com"
    },
    {
        "id": "3",
        "name": "Carlos Ferreira",
        "nif": "500032345",
        "phoneNumber": "+244900000003",
        "email": "carlos3@exemplo.com"
    }
]

function searchCustomers(el) {
    const container = document.querySelector("#customers-list");

    if (el.value.length === 0) {
        container.innerHTML = `<li class="text-sm mt-2 border-b p-2 cursor-pointer" onclick="toogleCustomerForm()">Criar Novo Cliente</li>`
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
    toogle("#menu-customer");
}


function toogleCustomerForm() {
    toogle("#customer-form")
    toogle("#menu-customer")
    toogle("#select-customer")

    clearCustomerForm()
}

function clearCustomerForm() {
    const nameEl = document.querySelector("#new-customer-name")
    const phoneEl = document.querySelector("#new-customer-phone")

    if (!nameEl || !phoneEl) return

    if (nameEl.value.length === 0 && phoneEl.value.length === 0) return

    nameEl.value = ""
    phoneEl.value = ""
}
