import numpy as np
from neural_network import NeuralNetwork

X_train = np.array([[0, 0], [0, 1], [1, 0], [1, 1]])
y_train = np.array([0, 1, 1, 0])

model = NeuralNetwork()
model.train(X_train, y_train)

model.model.save_weights('model.h5')
