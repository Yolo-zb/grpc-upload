<?php

namespace Client;

use Upload\UploadRequest;

require dirname(__FILE__) . '/../vendor/autoload.php';

class UploadClient
{
    public function send($filePath)
    {
        $fileExt = pathinfo($filePath, PATHINFO_EXTENSION);
        $client = new \Upload\BaseUploadClient('127.0.0.1:8887', [
            'credentials' => \Grpc\ChannelCredentials::createInsecure(),
        ]);
        $stream = $client->Upload();
        $request = new UploadRequest();
        $fileInfo = new \Upload\FileInfo();
        $fileInfo->setFilePath(dirname(__FILE__) . '/../storage/');
        $fileInfo->setFileExt('.' . $fileExt);
        $request->setFileInfo($fileInfo);
        $stream->write($request);
        $i = 0;                 //分割的块编号
        $fp = fopen($filePath, "r"); //要分割的文件
        while (!feof($fp)) {
            $chunkData = fread($fp, 2048);//切割的块大小 5m
            $request->setFile($chunkData);
            $stream->write($request);
            $i++;
        }
        fclose($fp);
        list($reply, $status) = $stream->wait();
        var_dump($reply->getSavePath());
        var_dump($status);
    }
}