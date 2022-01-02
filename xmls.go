package main

var loginXml = `<soapenv:Envelope xmlns:soapenv="http://schemas.xmlsoap.org/soap/envelope/" xmlns:ser="http://www.ca.com/UnicenterServicePlus/ServiceDesk">
    <soapenv:Header/>
    <soapenv:Body> 
        <ser:login>
             <!-- login ldap do usuario -->
            <username>{{.Username}}</username>
             <!-- senha ldap do usuario -->
            <password>{{.Password}}</password> 
        </ser:login>
    </soapenv:Body>
</soapenv:Envelope>`

var getHandleForUseridXml = `<soapenv:Envelope xmlns:soapenv="http://schemas.xmlsoap.org/soap/envelope/" xmlns:ser="http://www.ca.com/UnicenterServicePlus/ServiceDesk">
<soapenv:Header/>
<soapenv:Body>
   <ser:getHandleForUserid>
      <!-- sid retornado pelo loginXml -->
      <sid>{{.Handle}}</sid>
      <!-- login ldap do usuario -->
      <userID>{{.Username}}</userID>
   </ser:getHandleForUserid>
</soapenv:Body>
</soapenv:Envelope>`

var createRequestXml = `<soapenv:Envelope xmlns:soapenv="http://schemas.xmlsoap.org/soap/envelope/" xmlns:ser="http://www.ca.com/UnicenterServicePlus/ServiceDesk">
<soapenv:Header/>
<soapenv:Body>
   <ser:createRequest>
      <!-- sid retornado pelo loginXml -->
      <sid>{{.Handle}}</sid>
      <!-- handle retornado pelo getHandleForUseridXml -->
      <creatorHandle>{{.CreatorHandle}}</creatorHandle>
      <attrVals>
        <string>description</string> 
        <string>{{.Description}}</string> 
        <string>category</string> 
        <!-- area da solicitacao -->
        <!-- Servicos Administrativos.PRODABEL.Recursos Humanos.Registro de Ponto -->
        <string>pcat:424040</string> 
        <string>log_agent</string> 
        <string>{{.CreatorHandle}}</string> 
        <string>requested_by</string> 
        <string>{{.CreatorHandle}}</string> 
        <string>summary</string> 
        <string>{{.Summary}}</string> 
        <string>status</string> 
        <string>{{.Status}}</string> 
        <string>customer</string> 
        <string>{{.CreatorHandle}}</string> 
        <string>type</string> 
        <string>R</string> 
        <string>priority</string> 
        <string>0</string> 
        <string>assignee</string> 
        <string>{{.CreatorHandle}}</string> 
      </attrVals>
      <propertyValues>
      </propertyValues>
      <template></template>
      <attributes>
      </attributes>
      <newChangeHandle></newChangeHandle>
      <newChangeNumber></newChangeNumber>
      <newRequestHandle></newRequestHandle>
      <newRequestNumber></newRequestNumber>
   </ser:createRequest>
</soapenv:Body>
</soapenv:Envelope>`

var logoutXml = `<soapenv:Envelope xmlns:soapenv="http://schemas.xmlsoap.org/soap/envelope/" xmlns:ser="http://www.ca.com/UnicenterServicePlus/ServiceDesk">
   <soapenv:Header/>
   <soapenv:Body>
      <ser:logout>
         <sid>{{.Handle}}</sid>
      </ser:logout>
   </soapenv:Body>
</soapenv:Envelope>`

var doSelectIC = `<?xml version='1.0' encoding='UTF-8'?>
<soapenv:Envelope xmlns:soapenv="http://schemas.xmlsoap.org/soap/envelope/" xmlns:ser="http://www.ca.com/UnicenterServicePlus/ServiceDesk">
   <soapenv:Header/>
   <soapenv:Body>
      <ser:doSelect>
         <sid>{{.Handle}}</sid>
         <objectType>nr</objectType>
         <whereClause>name='{{.NomeIC}}'</whereClause>
         <maxRows>999</maxRows>
         <attributes>
            <!--1 or more repetitions:-->
            <string>id</string>
         </attributes>
      </ser:doSelect>
   </soapenv:Body>
</soapenv:Envelope>`

var doSelectStatus = `<?xml version='1.0' encoding='UTF-8'?>
<soapenv:Envelope xmlns:soapenv="http://schemas.xmlsoap.org/soap/envelope/" xmlns:ser="http://www.ca.com/UnicenterServicePlus/ServiceDesk">
   <soapenv:Header/>
   <soapenv:Body>
      <ser:doSelect>
         <sid>{{.Handle}}</sid>
         <objectType>chgstat</objectType>
         <whereClause>sym='{{.NomeStatus}}'</whereClause>
         <maxRows>1</maxRows>
         <attributes>
            <!--1 or more repetitions:-->
            <string>id</string>
         </attributes>
      </ser:doSelect>
   </soapenv:Body>
</soapenv:Envelope>`

var changeStatusXml = `<soapenv:Envelope xmlns:soapenv="http://schemas.xmlsoap.org/soap/envelope/" xmlns:ser="http://www.ca.com/UnicenterServicePlus/ServiceDesk">
   <soapenv:Header/>
   <soapenv:Body>
      <ser:changeStatus>
         <sid>{{.Handle}}</sid>
         <creator>{{.Handle}}</creator>
         <objectHandle>{{.ObjHandle}}</objectHandle>
         <description>{{.Desc}}</description>
         <newStatusHandle>{{.StatusHandle}}</newStatusHandle>
      </ser:changeStatus>
   </soapenv:Body>
</soapenv:Envelope>`
