<?php

namespace App\Http\Controllers\Api;

use Illuminate\Http\Request;
use App\Http\Controllers\Controller;
use App\Company;
use App\User;
use JWTAuth;


class CompanyController extends Controller
{
    public function getAllCompanies() {
      $company = Company::get();
      return response()->json($company, 200);
    }

    public function getCompany($id) {
      if (Company::where('id', $id)->exists()) {
        $company = Company::where('id', $id)->get();
        return response()->json($company, 200);
      } else {
        return response()->json([
          "message" => "Company not found"
        ], 404);
      }
    }

    public function createCompany(Request $request) {
      $user = self::getAuthenticatedUser();
    
      $user->companies()->create([
        "name" => $request->get('name'),
        "tel" => $request->get('tel'),
        "email" => $request->get('email'),
        "nif" => $request->get('nif')
      ]);
      
      return response()->json([
        "message" => "Company record created"
      ], 201);
    }

    public function updateCompany(Request $request, $id) {
      if (Company::where('id', $id)->exists()) {
        $company = Company::find($id);

        $company->name = $request->get('name');
        $company->save();

        return response()->json([
          "message" => "Records updated successfully"
        ], 200);
      } else {
        return response()->json([
          "message" => "Company not found"
        ], 404);
      }
    }

    public function deleteCompany($id) {
      if(Company::where('id', $id)->exists()) {
        $book = Company::find($id);
        $book->delete();

        return response()->json([
          "message" => "Records deleted"
        ], 202);
      } else {
        return response()->json([
          "message" => "Book not found"
        ], 404);
      }
    }

    public function getUserFromCompany($id) {
      if (Company::where('id', $id)->exists()) {
        $company = Company::where('id', $id)->get();
        $user = $company->load('user')->first();
        return response()->json($user->user, 200);
      } else {
        return response()->json([
          "message" => "Company not found"
        ], 404);
      }
    }

    public static function getAuthenticatedUser()
    {
        try {
            if (!$user = JWTAuth::parseToken()->authenticate()) {
                return response()->json(['user_not_found'], 404);
            }

        } catch (Tymon\JWTAuth\Exceptions\TokenExpiredException $e) {

            return response()->json(['token_expired'], $e->getStatusCode());

        } catch (Tymon\JWTAuth\Exceptions\TokenInvalidException $e) {

            return response()->json(['token_invalid'], $e->getStatusCode());

        } catch (Tymon\JWTAuth\Exceptions\JWTException $e) {

            return response()->json(['token_absent'], $e->getStatusCode());
        }

        return $user;
    }
}