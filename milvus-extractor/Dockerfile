FROM docker.io/library/python:3.13.1-slim

WORKDIR /code

COPY . .

RUN pip install --no-cache-dir -r requirements.txt

CMD ["python", "main.py"]
