import json
import time
import torch
import os
import requests
from PIL import Image
import timm
from sklearn.preprocessing import normalize
from timm.data import resolve_data_config
from timm.data.transforms_factory import create_transform
from urllib.parse import urlparse, parse_qs
from http.server import BaseHTTPRequestHandler, HTTPServer

class FeatureExtractor:
    def __init__(self, modelname):
        # Load the pre-trained model
        self.model = timm.create_model(
            modelname, pretrained=True, num_classes=0, global_pool="avg"
        )
        self.model.eval()

        # Get the input size required by the model
        self.input_size = self.model.default_cfg["input_size"]

        config = resolve_data_config({}, model=modelname)
        # Get the preprocessing function provided by TIMM for the model
        self.preprocess = create_transform(**config)

    def __call__(self, imagepath):
        # Preprocess the input image
        input_image = Image.open(imagepath).convert("RGB")  # Convert to RGB if needed
        input_image = self.preprocess(input_image)

        # Convert the image to a PyTorch tensor and add a batch dimension
        input_tensor = input_image.unsqueeze(0)

        # Perform inference
        with torch.no_grad():
            output = self.model(input_tensor)

        # Extract the feature vector
        feature_vector = output.squeeze().numpy()

        return normalize(feature_vector.reshape(1, -1), norm="l2").flatten()

class MyRequestHandler(BaseHTTPRequestHandler):
    def do_GET(self):
        # 当收到GET请求时调用
        if self.path.startswith('/hello'):
            url_components = urlparse(self.path)
            query_string = parse_qs(url_components.query)

            if query_string.get('image') == "":
                self.send_error(400, "Missing image parameter")
                return

            image_response = requests.get(query_string.get('image')[0])

            if image_response.status_code != 200:
                self.send_error(500, "Failed to fetch image")
                return

            extractor = FeatureExtractor("resnet34")

            filepath = os.path.join("./img/", "%d" % time.time()) + os.path.splitext(query_string.get('image')[0])[1]

            with open(filepath, "wb") as f:
                f.write(image_response.content)

            image_embedding = extractor(filepath)

            os.remove(filepath)

            self.send_response(200)  # 发送200响应代码表示成功
            self.end_headers()
            self.wfile.write(json.dumps(image_embedding.tolist()).encode('utf-8'))  # 响应体
        else:
            # 若路径不是'/hello'，返回404未找到
            self.send_error(404, "Page not Found")

def run(server_class=HTTPServer, handler_class=MyRequestHandler, port=8080):
    server_address = ('0.0.0.0', port)
    httpd = server_class(server_address, handler_class)
    print(f'Starting httpd on port {port}...')
    httpd.serve_forever()

if __name__ == "__main__":
    run()
