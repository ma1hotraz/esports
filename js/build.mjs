import logger from 'loglevel';
import * as esbuild from 'esbuild';

const result = await esbuild.build({
  entryPoints: ['src/app.jsx'],
  bundle: true,
  write: true,
  ignoreAnnotations: true,
  minify: true,
  jsx: 'automatic',
  outdir: '../public/js',
});

if (result.errors.length) {
  result.errors.forEach((err) => {
    logger.error(err);
  });
  process.exit(1);
}
