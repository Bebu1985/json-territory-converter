package main

import "github.com/Bebu1985/jsonTerritoryConverter/area"

func main() {
	Path := "D:\\Bebu\\Documents\\Versammlung\\Gebiete\\Fieldservice\\ServiceAreaAdministration\\Data\\"
	area.CreateOrUpdateAreaDatabase(Path+"areas.json", Path+"result.db")

}
