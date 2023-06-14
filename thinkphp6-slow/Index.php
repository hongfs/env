<?php
namespace app\controller;

use think\Db;
use app\BaseController;

class Index extends BaseController
{
    public function index()
    {
        return json($this->getOrder());
    }

    public function getOrder(?int $id = null)
    {
        if(is_null($id)) {
            $id = random_int(10100, 10400);
        }

        $data = Db::table('orders')
            ->where('orderNumber', $id)
            ->find();

        if(!$data) {
            return [];
        }

        $data['goods'] = Db::table('orderdetails')
            ->where('orderNumber', $data['id'])
            ->select();

        return $data;
    }

    public function hello($name = 'ThinkPHP6')
    {
        return 'hello,' . $name;
    }
}
