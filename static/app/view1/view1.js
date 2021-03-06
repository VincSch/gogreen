'use strict';

angular.module('myApp.view1', ['ngRoute', 'chart.js'])

    .config(['$routeProvider', function ($routeProvider) {
        $routeProvider.when('/view1', {
            templateUrl: 'view1/view1.html',
            controller: 'View1Ctrl'
        });
    }])

    .controller('View1Ctrl', function ($scope, $http) {

        $http({method: 'GET', url: 'http://localhost:12345/measurements'}).success(function (data) {
            var measureTimes = new Array();
            var temperature = new Array();
            var ph = new Array();

            function formatChart(date) {
                var d = new Date(date);
                var dateString = d.getDate() + "." + d.getMonth() + "." + d.getFullYear()
                    + " @ " + d.getHours() + ":" + d.getMinutes() + "." + d.getSeconds()
                return dateString
            }

            angular.forEach(data, function (item) {
                measureTimes.push(formatChart(item.time));
            });
            $scope.measuremnts = measureTimes

            angular.forEach(data, function (item) {
                temperature.push(item.temperature.value);
            });
            $scope.temp = temperature

            angular.forEach(data, function (item) {
                ph.push(item.ph.value);
            });
            $scope.ph = ph

            $scope.labels = measureTimes
            $scope.series = ['Temp', 'PH'];
            $scope.data = [temperature, ph]
            $scope.onClick = function (points, evt) {
                console.log(points, evt);
            };
            $scope.datasetOverride = [{yAxisID: 'y-axis-1'}, {yAxisID: 'y-axis-2'}];
            $scope.options = {
                scales: {
                    yAxes: [
                        {
                            id: 'y-axis-1',
                            type: 'linear',
                            display: true,
                            position: 'left'
                        },
                        {
                            id: 'y-axis-2',
                            type: 'linear',
                            display: true,
                            position: 'right'
                        }
                    ]
                }
            };

        });
    });


/*.controller("LineCtrl", function ($scope) {

    $scope.labels = ["January", "February", "March", "April", "May", "June", "July"];
    $scope.series = ['Series A', 'Series B'];
    $scope.data = [
        [65, 59, 80, 81, 56, 55, 40],
        [28, 48, 40, 19, 86, 27, 90]
    ];
    $scope.onClick = function (points, evt) {
        console.log(points, evt);
    };
    $scope.datasetOverride = [{ yAxisID: 'y-axis-1' }, { yAxisID: 'y-axis-2' }];
    $scope.options = {
        scales: {
            yAxes: [
                {
                    id: 'y-axis-1',
                    type: 'linear',
                    display: true,
                    position: 'left'
                },
                {
                    id: 'y-axis-2',
                    type: 'linear',
                    display: true,
                    position: 'right'
                }
            ]
        }
    };
});;*/
