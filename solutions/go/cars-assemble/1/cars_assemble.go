package cars


// CalculateWorkingCarsPerHour calculates how many working cars are
// produced by the assembly line every hour.
func CalculateWorkingCarsPerHour(productionRate int, successRate float64) float64 {
	
	var floated = successRate / 100
    
    return float64(productionRate) * floated 
    
}

// CalculateWorkingCarsPerMinute calculates how many working cars are
// produced by the assembly line every minute.
func CalculateWorkingCarsPerMinute(productionRate int, successRate float64) int {
	
    var avg = successRate / 100
    var ratePerMinute int
    
    ratePerMinute = int((float64(productionRate) * avg) / 60)
    
    return ratePerMinute
    
}

// CalculateCost works out the cost of producing the given number of cars.
func CalculateCost(carsCount int) uint {
	
    chunk := (carsCount / 10) * 95000

    rest := (carsCount % 10) * 10000

    total := chunk + rest
    
    return uint(total)

}