<?php
// GENERATED CODE -- DO NOT EDIT!

namespace Upload;

/**
 */
class BaseUploadClient extends \Grpc\BaseStub {

    /**
     * @param string $hostname hostname
     * @param array $opts channel options
     * @param \Grpc\Channel $channel (optional) re-use channel object
     */
    public function __construct($hostname, $opts, $channel = null) {
        parent::__construct($hostname, $opts, $channel);
    }

    /**
     * @param array $metadata metadata
     * @param array $options call options
     * @return \Grpc\ClientStreamingCall
     */
    public function Upload($metadata = [], $options = []) {
        return $this->_clientStreamRequest('/upload.BaseUpload/Upload',
        ['\Upload\UploadResponse','decode'],
        $metadata, $options);
    }

}
