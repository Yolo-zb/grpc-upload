<?php

namespace console;

require dirname(__FILE__) . '/../vendor/autoload.php';

$b = new \Client\UploadClient();

$b->send(__DIR__ . '/../storage/ABC.jpeg');
