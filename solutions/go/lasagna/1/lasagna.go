package lasagna

//defined constant
const OvenTime = 40

// RemainingOvenTime returns the remaining minutes based on the `actual` minutes already in the oven.
func RemainingOvenTime(actualMinutesInOven int) int {

    var realTime int

    realTime = OvenTime - actualMinutesInOven
    
    return realTime
}

// PreparationTime calculates the time needed to prepare the lasagna based on the amount of layers.
func PreparationTime(numberOfLayers int) int {

    var timeSpent int

    timeSpent = numberOfLayers * 2

    return timeSpent
      
}

// ElapsedTime calculates the time elapsed cooking the lasagna. This time includes the preparation time and the time the lasagna is baking in the oven.
func ElapsedTime(numberOfLayers, actualMinutesInOven int) int {

	var time int
   
    time = ((numberOfLayers * 2) + actualMinutesInOven)

    return time
    
}
