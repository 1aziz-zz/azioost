FROM maven:alpine
COPY src /home/src
COPY pom.xml /home
WORKDIR /home
RUN mvn clean install
CMD ["java","-jar","target/eureka-service-0.1.jar"]
