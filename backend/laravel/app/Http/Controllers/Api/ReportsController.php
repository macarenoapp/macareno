<?php

namespace App\Http\Controllers\Api;

use App\Http\Controllers\Controller;
use Illuminate\Support\Facades\Redis;
use JWTAuth;
use Illuminate\Support\Facades\DB;

class ReportsController extends Controller
{

    public function getReports() {
        $user = self::getAuthenticatedUser();
        $user = $user->original["user"]->id;
        error_log(json_encode($user));
        
        $discotecas = DB::select("SELECT id,name FROM discotecas WHERE user = $user");

        error_log(json_encode($discotecas));

        foreach ($discotecas as $key => $discoteca) {
            $redis = Redis::get($discoteca->id);
            error_log(json_encode($redis));
            if (substr($redis, 0, 1) === ',') { 
                $redis = substr($redis, 1);
            }

            $total = "(".$redis.")";
            error_log(json_encode($total));

            if ($redis){
                $users = DB::select("SELECT username,email FROM users WHERE id in $total");
                $discotecas[$key]->users = $users;
            } else {
                $discotecas[$key]->users = null;
            }
        }

        return response()->json($discotecas, 200);
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

        return response()->json(compact('user'));
    }

}
