const path = require('path')
const execFile = require('child_process').execFile
const gulp = require('gulp')
const stylus = require('gulp-stylus')
const pug = require('gulp-pug')
const errorHandler = require('gulp-error-handle')
const minifyCssNames = require('gulp-minify-css-names')
const csso = require('gulp-csso')
const kill = require('tree-kill')

process.on('uncaughtException', (err) => {
  console.log('Exception:', err)
  process.exit(2)
})

// changes to these paths must be correlated with changes in httpserver/httpserver.go
const buildDir = '.build'
const buildPublicDir = path.join(buildDir, 'public')
const buildTemplatesDir = path.join(buildDir, 'templates')

const frontendDir = 'frontend'
const publicDir = path.join(frontendDir, 'public')
const templatesDir = path.join(frontendDir, 'templates')
const stylesDir = path.join(frontendDir, 'styles')

function startServer () {
  const serverProcess = execFile('go', ['run', 'main.go'])
  serverProcess.stdout.pipe(process.stdout)
  serverProcess.stderr.pipe(process.stderr)
  return serverProcess
}

gulp.task('assets', function () {
  return gulp.src(path.join(publicDir, '**/*'))
    .pipe(gulp.dest(buildPublicDir))
})

gulp.task('css', function () {
  return gulp.src(path.join(stylesDir, '*.styl'))
    .pipe(errorHandler())
    .pipe(stylus())
    .pipe(csso())
    .pipe(gulp.dest(buildPublicDir))
})

gulp.task('html', function () {
  return gulp.src(path.join(templatesDir, '*.pug'))
    .pipe(errorHandler())
    .pipe(pug())
    .pipe(gulp.dest(buildTemplatesDir))
})

gulp.task('go-server', function () {
  var serverProcess = startServer()

  gulp.watch([
    '**/*.go',
    buildDir,
  ], function () {
    console.log('Restarting go server...')
    kill(serverProcess.pid, 'SIGTERM')
    serverProcess = startServer()
  })
})

gulp.task('build', ['css', 'html'], function() {
  return gulp.src([
    path.join(buildPublicDir, '*.css'),
    path.join(buildTemplatesDir, '*.html')
  ], {base: '.build/'})
      .pipe(minifyCssNames({
        postfix: '',
        prefix: 'ots-'
      }))
      .pipe(gulp.dest('.build/'))
})

gulp.task('default', ['css', 'html', 'assets', 'go-server'], function () {
  gulp.watch(path.join(stylesDir, '**/*.styl'), ['css'])
  gulp.watch(path.join(templatesDir, '*.pug'), ['html'])
  gulp.watch(path.join(publicDir, '**/*'), ['assets'])
})

