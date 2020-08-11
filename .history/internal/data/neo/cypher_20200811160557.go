package neo

const seedDB = `
	MERGE (sysAdm :Role {id: 1})
	ON CREATE SET sysAdm.name = 'System Administrator', sysAdm.parent_id = null
	MERGE (locMan :Role {id: 2})
	ON CREATE SET locMan.name = 'Location Manager', locMan.parent_id = 1
	MERGE (sup :Role {id: 3})
	ON CREATE SET sup.name = 'Supervisor', sup.parent_id = 2
	MERGE (empl :Role {id: 4})
	ON CREATE SET empl.name = 'Employee', empl.parent_id = 3
	MERGE (train :Role {id: 5})
	ON CREATE SET train.name = 'Trainer', train.parent_id = 3
	MERGE (adam :User {id: 1})
	ON CREATE SET adam.name = 'Adam Admin', adam.role_id = 1
	MERGE (em :User {id: 2})
	ON CREATE SET em.name = 'Emily Employee', em.role_id = 4
	MERGE (sam :User {id: 3})
	ON CREATE SET sam.name = 'Sam Supervisor', sam.role_id = 3
	MERGE (mary :User {id: 4})
	ON CREATE SET mary.name = 'Mary Manager', mary.role_id = 2
	MERGE (steve :User {id: 5})
	ON CREATE SET steve.name = 'Steve Trainer', steve.role_id = 5
	MERGE (sysAdm)<-[:REPORTS_TO]-(locMan)
	MERGE (locMan)<-[:REPORTS_TO]-(sup)
	MERGE (sup)<-[:REPORTS_TO]-(empl)
	MERGE (sup)<-[:REPORTS_TO]-(train)
	MERGE (sysAdm)<-[:HAS_ROLE_OF]-(adam)
	MERGE (empl)<-[:HAS_ROLE_OF]-(em)
	MERGE (sup)<-[:HAS_ROLE_OF]-(sam)
	MERGE (locMan)<-[:HAS_ROLE_OF]-(mary)
	MERGE (train)<-[:HAS_ROLE_OF]-(steve)
`