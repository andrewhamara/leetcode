type SeatManager struct {
    seats []int
}

func Constructor(n int) SeatManager {
    sm := SeatManager{seats: make([]int, n)}
    return sm
}

func (this *SeatManager) Reserve() int {
    for i,v := range this.seats {
        if v == 0 {
            this.seats[i] = 1
            return i + 1
        }
    }
    return -1
}

func (this *SeatManager) Unreserve(seatNumber int) {
    this.seats[seatNumber-1] = 0
}
