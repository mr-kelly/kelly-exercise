var gulp = require('gulp');

default_task = function() {

	gulp.src('../spider_project/data.db')
		.pipe(gulp.dest('../star_web/'));

	gulp.src('bower_components/bootstrap/dist/js/bootstrap.min.js')
		.pipe(gulp.dest('../star_web/static/js'));

	gulp.src('bower_components/jquery/dist/jquery.min.js')
		.pipe(gulp.dest('../star_web/static/js'));

	gulp.src('bower_components/jquery/dist/jquery.min.map')
		.pipe(gulp.dest('../star_web/static/js'));

    gulp.src('bower_components/bootstrap/dist/css/bootstrap.min.css')
		.pipe(gulp.dest('../star_web/static/css'));
	gulp.src('bower_components/bootstrap/dist/css/bootstrap-theme.min.css')
		.pipe(gulp.dest('../star_web/static/css'));
	gulp.src('bower_components/bootstrap/dist/css/bootstrap.css.map')
		.pipe(gulp.dest('../star_web/static/css'));
	gulp.src('bower_components/bootstrap/dist/css/bootstrap-theme.css.map')
		.pipe(gulp.dest('../star_web/static/css'));

	gulp.src('bower_components/bootstrap/dist/fonts/*.*')
		.pipe(gulp.dest('../star_web/static/fonts'));
    	
}

gulp.task('init', default_task);
gulp.task('default', default_task);