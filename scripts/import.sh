home="$PWD"
ogr2ogr_image=registry.docker.ir/osgeo/gdal:alpine-small-latest
data_path=$home/data/districts.geojson
config_file=$home/config/.env
export $(grep -v '^#' $config_file | xargs)
db_connection="host=127.0.0.1 dbname=$POSTGRES_DB user=$POSTGRES_USER password=$POSTGRES_PASSWORD"

docker run --rm -v=$data_path:/data.geojson --network=host $ogr2ogr_image ogr2ogr -f "PostgreSQL" PG:"$db_connection" /data.geojson -nln districts

if [ $? -eq 0 ]
then
echo "data imported successfully"
else
echo "data not imported"
fi
