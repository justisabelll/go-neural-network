package main

import "fmt"

// neuron receives a set of inputs and a bias
// each input has a corresponding weight
type Neuron struct {
    Weights []float32
    Bias    float32
}

type Layer struct {
    Neurons []Neuron
    Inputs  []float32
}

func (l *Layer) addNeuron(weights []float32, bias float32) {
    l.Neurons = append(l.Neurons, Neuron{Weights: weights, Bias: bias})
}

// multiply each input by its weight,
// sum all products, then add the bias
func calcOutput(n *Neuron, inputs []float32) float32 {
    output := float32(0)

    for i := 0; i < len(inputs); i++{
        output += inputs[i] * n.Weights[i]
    }

    return output + n.Bias
}


// example usage 
func main() {
    firstNeuron := Neuron{
        Weights: []float32{.2, .8, -.5},
        Bias:    2,
    }

    secondNeuron := Neuron{
        Weights: []float32{.2, .8, -.5, 1.0},
        Bias:    2,
    }

    fmt.Println("first neuron output:", calcOutput(&firstNeuron, []float32{1, 2, 3}))
    fmt.Println("second neuron output:", calcOutput(&secondNeuron, []float32{1, 2, 3, 2.5}))

    firstLayer := Layer{
        Inputs: []float32{1, 2, 3, 2.5},
    }

    firstLayer.addNeuron([]float32{.2, .8, -.5, 1}, 2)
    firstLayer.addNeuron([]float32{.5, -.91, .26, -.5}, 3)
    firstLayer.addNeuron([]float32{-.26, -.27, .17, .87}, .5)

    for i, neuron := range firstLayer.Neurons {
        fmt.Printf("first layer: neuron %d output: %f\n", i, calcOutput(&neuron, firstLayer.Inputs))
    }

}


