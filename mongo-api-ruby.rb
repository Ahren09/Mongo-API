#!/usr/bin/ruby

require 'mongo'
client = Mongo::Client.new([ '0.0.0.0:27017' ], :database => 'test')
db = client.database

# puts db.client.database
# puts db.collection_names

collection = client[:newMgoTest]
