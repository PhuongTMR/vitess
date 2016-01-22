<?php
// DO NOT EDIT! Generated by Protobuf-PHP protoc plugin 1.0
// Source: tabletmanagerdata.proto
//   Date: 2016-01-22 01:34:35

namespace Vitess\Proto\Tabletmanagerdata {

  class ExecuteFetchAsDbaResponse extends \DrSlump\Protobuf\Message {

    /**  @var \Vitess\Proto\Query\QueryResult */
    public $result = null;
    

    /** @var \Closure[] */
    protected static $__extensions = array();

    public static function descriptor()
    {
      $descriptor = new \DrSlump\Protobuf\Descriptor(__CLASS__, 'tabletmanagerdata.ExecuteFetchAsDbaResponse');

      // OPTIONAL MESSAGE result = 1
      $f = new \DrSlump\Protobuf\Field();
      $f->number    = 1;
      $f->name      = "result";
      $f->type      = \DrSlump\Protobuf::TYPE_MESSAGE;
      $f->rule      = \DrSlump\Protobuf::RULE_OPTIONAL;
      $f->reference = '\Vitess\Proto\Query\QueryResult';
      $descriptor->addField($f);

      foreach (self::$__extensions as $cb) {
        $descriptor->addField($cb(), true);
      }

      return $descriptor;
    }

    /**
     * Check if <result> has a value
     *
     * @return boolean
     */
    public function hasResult(){
      return $this->_has(1);
    }
    
    /**
     * Clear <result> value
     *
     * @return \Vitess\Proto\Tabletmanagerdata\ExecuteFetchAsDbaResponse
     */
    public function clearResult(){
      return $this->_clear(1);
    }
    
    /**
     * Get <result> value
     *
     * @return \Vitess\Proto\Query\QueryResult
     */
    public function getResult(){
      return $this->_get(1);
    }
    
    /**
     * Set <result> value
     *
     * @param \Vitess\Proto\Query\QueryResult $value
     * @return \Vitess\Proto\Tabletmanagerdata\ExecuteFetchAsDbaResponse
     */
    public function setResult(\Vitess\Proto\Query\QueryResult $value){
      return $this->_set(1, $value);
    }
  }
}

