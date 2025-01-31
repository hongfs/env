FROM docker.io/library/python:3.13.1-slim

COPY . .

RUN pip install --quiet -r requirements.txt

CMD ["python", "main.py"]
