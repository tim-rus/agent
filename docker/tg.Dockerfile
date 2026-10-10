FROM python:3.14.7-alpine
WORKDIR /app
COPY ./tg/requirements.txt ./requirements.txt
RUN pip install -r requirements.txt
COPY ./tg/ .
CMD ["python", "main.py"]
