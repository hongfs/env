<?php

require_once __DIR__ . '/vendor/autoload.php';

use PhpOffice\PhpSpreadsheet\IOFactory;

// 加载我们的数据文件
$spreadsheet = IOFactory::load('/data/data.xlsx');

$sheet = $spreadsheet->getActiveSheet();

// 获取表格数据
$sheet_data = $sheet->toArray(null, true, true, true);

// 默认KEY所在列
$key_index = 'C';

// 获取第3行的语言key
foreach ($sheet_data[3] as $key => $lang) {
    if (!$lang || $key === 0) continue;

    $data = [];

    // 循环表格所有数据
    foreach ($sheet_data as $index => $item) {
        if ($index <= 3) continue;

        // 如果有固定 KEY 就使用或者采用默认 KEY
        if (!is_null($item['A'])) {
            $data[$item['A']] = trim($item[$key], ' ');
        } else if (!is_null($item[$key_index])) {
            $data[$item[$key_index]] = trim($item[$key], ' ');
        }
    }

    if(!count($data)) {
        continue;
    }

    // 文件生成
    file_put_contents('/data/' . time() . '_' . $lang . '.json', json_encode($data, JSON_UNESCAPED_UNICODE));
}
