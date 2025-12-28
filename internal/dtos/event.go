package dtos


type Event struct {
    T uint8  // event type (1 = click)
    O uint8  // object type (1 = course, 2 = product, ...)
    I uint32 // object id (numeric or hash)
    U uint32 // user id (optional, hashed)
    Ts int64 // unix timestamp (seconds)
}
