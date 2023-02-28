FROM python:3.11-bullseye
COPY . /app
EXPOSE 5001
WORKDIR /app
RUN pip3 install -r requirements.txt
CMD [ "python3", "-m" , "flask", "run", "--host=0.0.0.0", "--port=5001"]