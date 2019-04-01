# Version can be changed
FROM node:10
WORKDIR /code

RUN git clone {{GIT_REPO_PATH}}

RUN cd runtime-wrapper \
  && yarn install --frozen-lockfile --no-cache --production \
  && yarn global add forever

EXPOSE 55668

CMD ["forever", "runtime-wraper/src/app.js"]