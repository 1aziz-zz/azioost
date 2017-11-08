FROM java:8

RUN wget -q https://services.gradle.org/distributions/gradle-3.3-bin.zip \
    && unzip gradle-3.3-bin.zip -d /opt \
    && rm gradle-3.3-bin.zip

ENV GRADLE_HOME /opt/gradle-3.3
ENV PATH $PATH:/opt/gradle-3.3/bin
COPY src /home/src
COPY settings.gradle /home
COPY build.gradle /home
WORKDIR /home
RUN gradle build
CMD ["java", "-jar", "build/libs/personservice-0.1.jar"]
