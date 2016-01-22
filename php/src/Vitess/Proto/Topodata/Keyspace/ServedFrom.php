<?php
// DO NOT EDIT! Generated by Protobuf-PHP protoc plugin 1.0
// Source: topodata.proto
//   Date: 2016-01-22 01:34:42

namespace Vitess\Proto\Topodata\Keyspace {

  class ServedFrom extends \DrSlump\Protobuf\Message {

    /**  @var int - \Vitess\Proto\Topodata\TabletType */
    public $tablet_type = null;
    
    /**  @var string[]  */
    public $cells = array();
    
    /**  @var string */
    public $keyspace = null;
    

    /** @var \Closure[] */
    protected static $__extensions = array();

    public static function descriptor()
    {
      $descriptor = new \DrSlump\Protobuf\Descriptor(__CLASS__, 'topodata.Keyspace.ServedFrom');

      // OPTIONAL ENUM tablet_type = 1
      $f = new \DrSlump\Protobuf\Field();
      $f->number    = 1;
      $f->name      = "tablet_type";
      $f->type      = \DrSlump\Protobuf::TYPE_ENUM;
      $f->rule      = \DrSlump\Protobuf::RULE_OPTIONAL;
      $f->reference = '\Vitess\Proto\Topodata\TabletType';
      $descriptor->addField($f);

      // REPEATED STRING cells = 2
      $f = new \DrSlump\Protobuf\Field();
      $f->number    = 2;
      $f->name      = "cells";
      $f->type      = \DrSlump\Protobuf::TYPE_STRING;
      $f->rule      = \DrSlump\Protobuf::RULE_REPEATED;
      $descriptor->addField($f);

      // OPTIONAL STRING keyspace = 3
      $f = new \DrSlump\Protobuf\Field();
      $f->number    = 3;
      $f->name      = "keyspace";
      $f->type      = \DrSlump\Protobuf::TYPE_STRING;
      $f->rule      = \DrSlump\Protobuf::RULE_OPTIONAL;
      $descriptor->addField($f);

      foreach (self::$__extensions as $cb) {
        $descriptor->addField($cb(), true);
      }

      return $descriptor;
    }

    /**
     * Check if <tablet_type> has a value
     *
     * @return boolean
     */
    public function hasTabletType(){
      return $this->_has(1);
    }
    
    /**
     * Clear <tablet_type> value
     *
     * @return \Vitess\Proto\Topodata\Keyspace\ServedFrom
     */
    public function clearTabletType(){
      return $this->_clear(1);
    }
    
    /**
     * Get <tablet_type> value
     *
     * @return int - \Vitess\Proto\Topodata\TabletType
     */
    public function getTabletType(){
      return $this->_get(1);
    }
    
    /**
     * Set <tablet_type> value
     *
     * @param int - \Vitess\Proto\Topodata\TabletType $value
     * @return \Vitess\Proto\Topodata\Keyspace\ServedFrom
     */
    public function setTabletType( $value){
      return $this->_set(1, $value);
    }
    
    /**
     * Check if <cells> has a value
     *
     * @return boolean
     */
    public function hasCells(){
      return $this->_has(2);
    }
    
    /**
     * Clear <cells> value
     *
     * @return \Vitess\Proto\Topodata\Keyspace\ServedFrom
     */
    public function clearCells(){
      return $this->_clear(2);
    }
    
    /**
     * Get <cells> value
     *
     * @param int $idx
     * @return string
     */
    public function getCells($idx = NULL){
      return $this->_get(2, $idx);
    }
    
    /**
     * Set <cells> value
     *
     * @param string $value
     * @return \Vitess\Proto\Topodata\Keyspace\ServedFrom
     */
    public function setCells( $value, $idx = NULL){
      return $this->_set(2, $value, $idx);
    }
    
    /**
     * Get all elements of <cells>
     *
     * @return string[]
     */
    public function getCellsList(){
     return $this->_get(2);
    }
    
    /**
     * Add a new element to <cells>
     *
     * @param string $value
     * @return \Vitess\Proto\Topodata\Keyspace\ServedFrom
     */
    public function addCells( $value){
     return $this->_add(2, $value);
    }
    
    /**
     * Check if <keyspace> has a value
     *
     * @return boolean
     */
    public function hasKeyspace(){
      return $this->_has(3);
    }
    
    /**
     * Clear <keyspace> value
     *
     * @return \Vitess\Proto\Topodata\Keyspace\ServedFrom
     */
    public function clearKeyspace(){
      return $this->_clear(3);
    }
    
    /**
     * Get <keyspace> value
     *
     * @return string
     */
    public function getKeyspace(){
      return $this->_get(3);
    }
    
    /**
     * Set <keyspace> value
     *
     * @param string $value
     * @return \Vitess\Proto\Topodata\Keyspace\ServedFrom
     */
    public function setKeyspace( $value){
      return $this->_set(3, $value);
    }
  }
}

