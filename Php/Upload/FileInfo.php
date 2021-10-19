<?php
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: proto/upload.proto

namespace Upload;

use Google\Protobuf\Internal\GPBType;
use Google\Protobuf\Internal\RepeatedField;
use Google\Protobuf\Internal\GPBUtil;

/**
 * Generated from protobuf message <code>upload.FileInfo</code>
 */
class FileInfo extends \Google\Protobuf\Internal\Message
{
    /**
     * Generated from protobuf field <code>string filePath = 1;</code>
     */
    protected $filePath = '';
    /**
     * Generated from protobuf field <code>string fileExt = 2;</code>
     */
    protected $fileExt = '';

    /**
     * Constructor.
     *
     * @param array $data {
     *     Optional. Data for populating the Message object.
     *
     *     @type string $filePath
     *     @type string $fileExt
     * }
     */
    public function __construct($data = NULL) {
        \GPBMetadata\Proto\Upload::initOnce();
        parent::__construct($data);
    }

    /**
     * Generated from protobuf field <code>string filePath = 1;</code>
     * @return string
     */
    public function getFilePath()
    {
        return $this->filePath;
    }

    /**
     * Generated from protobuf field <code>string filePath = 1;</code>
     * @param string $var
     * @return $this
     */
    public function setFilePath($var)
    {
        GPBUtil::checkString($var, True);
        $this->filePath = $var;

        return $this;
    }

    /**
     * Generated from protobuf field <code>string fileExt = 2;</code>
     * @return string
     */
    public function getFileExt()
    {
        return $this->fileExt;
    }

    /**
     * Generated from protobuf field <code>string fileExt = 2;</code>
     * @param string $var
     * @return $this
     */
    public function setFileExt($var)
    {
        GPBUtil::checkString($var, True);
        $this->fileExt = $var;

        return $this;
    }

}

