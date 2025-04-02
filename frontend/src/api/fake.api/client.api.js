const rations = {
    small: { _id: "0", name: "1200", color: "blue-300"},
    medium: { _id: "1", name: "1500", color: "indigo-500"},
    big: { _id: "2", name: "1800", color: "pink-500"},
    large: { _id: "3", name: "2200", color: "orange-500"}
}

const times = {
    early: {_id: "0", name: "18-20", color: "success"},
    late: {_id: "1", name: "20-22", color: "info"}
}

const clients = [
    {
        _id: "0",
        name: "Петров Петр Петрович",
        rations: rations.big,
        address: "Кирова 1",
        instrument: "yes",
        time: times.late
    },
    {
        _id: "1",
        name: "Иванов Иван Иванович",
        rations: rations.medium,
        address: "Гагарина 2",
        instrument: "no",
        time: times.early
    }
]

export function fetchAll() {
    return clients
}