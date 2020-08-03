const sass = require('node-sass');

module.exports = function (grunt) {
    grunt.initConfig({
        copy: {
            build: {
                files: [
                    {
                        expand: true,
                        cwd: './src',
                        src: ['**', '!**/*.sass', '!**/*.ts', '!**/*.go', '!**/go.mod'],
                        dest: './dist'
                    },
                    {
                        expand: true,
                        cwd: ".",
                        src: ["./res/**"],
                        dest: './dist'
                    }
                ]
            }
        },
        ts: {
            website: {
                files: [
                    {
                        src: ['src/**/*.ts'],
                        dest: './dist'
                    }
                ],
                options: {
                    target: 'es6',
                    sourceMap: false,
                    rootDir: 'src'
                }
            }
        },
        sass: {
            options: {
                implementation: sass,
                sourceMap: true
            },
            dist: {
                files: [
                    {
                        expand: true,
                        cwd: 'src',
                        src: ['**/*.sass'],
                        dest: './dist',
                        ext: '.css'
                    }
                ]
            }
        },
        watch: {
            ts: {
                files: ['src/**/*.ts'],
                tasks: ['ts:website']
            },
            sass: {
                files: ['src/**/*.sass'],
                tasks: ['sass']
            },
            images: {
                files: ['res/**'],
                tasks: ['copy']
            },
            views: {
                files: ['src/**/*.html'],
                tasks: ['copy']
            }
        }
    });

    grunt.loadNpmTasks('grunt-contrib-copy');
    grunt.loadNpmTasks('grunt-contrib-watch');
    grunt.loadNpmTasks('grunt-ts');
    grunt.loadNpmTasks('grunt-sass');

    grunt.registerTask('build', ['copy', 'ts', 'sass']);
}