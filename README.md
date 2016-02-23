# MySQLbeat
MySQLbeat - Elastic Beat for MySQL Monitoring

MySQLbeat is the <a
href="https://www.elastic.co/products/beats">beat</a> used for MySQL
monitoring.  It is a lightweight agent that installs on MySQL servers,
reads periodically from MySQL statistics, and indexes them in
Elasticsearch.

# Elasticsearch template

To apply MySQLbeat template for MySQL status:

    curl -XPUT 'http://localhost:9200/_template/mysqlbeat -d@etc/mysqlbeat.template.json

# Building, Testing, and Running

    # Build
    export GO15VENDOREXPERIMENT=1
    make

    # Test
    make unit-tests
    make testsuite

    # Run
    ./mysqlbeat -c mysqlbeat.yml

# Exported Fields

MySQLbeat exports a variety of sections:

* type: status - Output of 'SHOW GLOBAL STATUS'
* type: variables - Output of 'SHOW GLOBAL VARIABLES'
* type: sysschema - 
* type: repl - 

...


