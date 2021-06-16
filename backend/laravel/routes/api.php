<?php

use Illuminate\Http\Request;
use Illuminate\Support\Facades\Route;
/*
|--------------------------------------------------------------------------
| API Routes
|--------------------------------------------------------------------------
|
| Here is where you can register API routes for your application. These
| routes are loaded by the RouteServiceProvider within a group which
| is assigned the "api" middleware group. Enjoy building your API!
|
*/

Route::group(['namespace' => 'Api'], function () {
    Route::group(['middleware' => ['jwt.verify']], function() {
        Route::get('user', 'UserController@getAuthenticatedUser');
        Route::post('company', 'CompanyController@createCompany');
        Route::get('reports', 'ReportsController@getReports');
        Route::get('users/companies', 'UserController@getCompaniesFromUser');
    });

    Route::get('companies', 'CompanyController@getAllCompanies');
    Route::get('company/{id}', 'CompanyController@getCompany');
    Route::get('company/user/{id}', 'CompanyController@getUserFromCompany');
    Route::post('users/register', 'UserController@register');
    Route::post('users/login', 'UserController@login');
    Route::put('company/{id}', 'CompanyController@updateCompany');
    Route::delete('company/{id}','CompanyController@deleteCompany');
});


// // DISCOTECAS
//     // POST
//     Route::post('discoteca', 'Api\DiscotecaController@createDsicoteca');

// // EVENTS
//     // POST
//     Route::post('event', 'Api\EventController@createEvent');