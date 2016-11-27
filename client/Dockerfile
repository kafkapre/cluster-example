FROM node





RUN mkdir -p /opt/AngularApp/src/service
WORKDIR /opt/AngularApp

COPY package.json /opt/AngularApp/
RUN npm install

ADD src/ /opt/AngularApp/src
ADD src/service /opt/AngularApp/src/service
COPY index.html /opt/AngularApp/
COPY typings.json /opt/AngularApp/
COPY tsconfig.json /opt/AngularApp/




RUN ls
RUN echo -----------------
RUN ls src
RUN echo -----------------
RUN ls src/service


#RUN npm install http-server
RUN find / -name npm

EXPOSE 3000

CMD ["/usr/local/bin/npm", "start"]
