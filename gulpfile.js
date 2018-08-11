const execFile = require('child_process').execFile
const gulp = require('gulp')
const stylus = require('gulp-stylus')
const pug = require('gulp-pug')
const errorHandler = require('gulp-error-handle')

gulp.task('css', function () {
  return gulp.src('styles/*.styl')
    .pipe(errorHandler())
    .pipe(stylus())
    .pipe(gulp.dest('.build/assets'))
})

gulp.task('html', function () {
  return gulp.src('templates/*.pug')
    .pipe(errorHandler())
    .pipe(pug())
    .pipe(gulp.dest('.build/templates'))
})

gulp.task('go-server', function () {
  var server
  setTimeout(function () {
    console.log('killing server')
    server.kill('SIGTERM')
  }, 4000)
  function startServer () {
    var goProcess = execFile('go', ['run', 'main.go'])
    goProcess.stdout.on('data', function (data) {
      console.log('stdout: ' + data)
    })
    goProcess.stderr.on('data', function (data) {
      // console.log('stderr: ' + data)
      console.log(data)
    })
    goProcess.on('close', function (code) {
      console.log('closing code: ' + code)
    })
    goProcess.on('error', function (err) {
      console.log('error: ' + err)
    })

    function onExit() {
      goProcess.kill('SIGINT')
      process.exit(0)
    }
    process.on('SIGINT', onExit)
    process.on('exit', onExit)

    return goProcess
  }

  gulp.watch('**/*.go', function () {
    console.log('killing process')
    server.kill('SIGINT') // ici le kill ne fonctionne pas vraiment, et le process se trouvé sous init...
    setTimeout(function () {
      console.log('creating new process')
      server = startServer()
      console.log('created new process')
    }, 2000)
  })
  server = startServer()
})

gulp.task('build', ['css', 'html'])

gulp.task('default', ['css', 'html'], function () {
  gulp.watch('styles/**/*.styl', ['css'])
  gulp.watch('templates/**/*.pug', ['html'])
})

