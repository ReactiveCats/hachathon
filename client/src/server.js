const fastify = require('fastify')({
  logger: { level: 'error' },
  pluginTimeout: 0,
});
const makeNext = require('next');

const port = parseInt(process.env.PORT, 10) || 3000;
const dev = process.env.NODE_ENV !== 'production';

fastify.register((fastify, options, next) => {
  const app = makeNext({ dev });
  const handle = app.getRequestHandler();
  app
    .prepare()
    .then(() => {
      if (dev) {
        fastify.get('/_next/*', (req, reply) => {
          return handle(req.raw, reply.raw).then(() => {
            reply.sent = true;
          });
        });
      }

      fastify.all('/*', (req, reply) => {
        return handle(req.raw, reply.raw).then(() => {
          reply.sent = true;
        });
      });

      fastify.setNotFoundHandler((request, reply) => {
        return app.render404(request.raw, reply.raw).then(() => {
          reply.sent = true;
        });
      });

      next();
    })
    .catch((err) => next(err));
});

fastify.listen(port, (err) => {
  if (err) {
    throw err;
  }
  console.log(`Server listening on http://localhost:${port}`);
});
