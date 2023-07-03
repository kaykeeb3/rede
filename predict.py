import json
import numpy as np
from neural_network import NeuralNetwork

model = NeuralNetwork()
model.model.load_weights('model.h5')

def predict(request):
    X = np.array(request['x'])
    y = model.predict(X)
    return {'y': y.tolist()}

if __name__ == '__main__':
    request = json.loads(input())
    response = predict(request)
    print(json.dumps(response))
  
