package wolfram_test

import (
    wolfram "github.com/KeizoBookman/go-wolfram-alpha"
    "testing"
)


func TestParse(t *testing.T) {
xmlValue =`

<?xml version='1.0' encoding='UTF-8'?>
<queryresult success='true'
    error='false'
    numpods='2'
    datatypes='AdministrativeDivision,City,Country,Internet,University'
    timedout='Data,Character'
    timedoutpods=''
    timing='1.179'
    parsetiming='0.232'
    parsetimedout='false'
    recalculate='http://www4b.wolframalpha.com/api/v2/recalc.jsp?id=MSPa72881d3fce4ibfii72fg000037141g16e3acbcci&amp;s=38'
    id='MSPa72891d3fce4ibfii72fg0000663g72e68d5a9gf3'
    host='http://www4b.wolframalpha.com'
    server='38'
    related='http://www4b.wolframalpha.com/api/v2/relatedQueries.jsp?id=MSPa72901d3fce4ibfii72fg00001g866f27412e3478&amp;s=38'
    version='2.6'>
 <pod title='Input interpretation'
     scanner='Identity'
     id='Input'
     position='100'
     error='false'
     numsubpods='1'>
  <subpod title=''>
   <plaintext>Karlsruher Institut für Technologie</plaintext>
   <img src='http://www4b.wolframalpha.com/Calculate/MSP/MSP72911d3fce4ibfii72fg00002g44a1h3g41ch4f9?MSPStoreType=image/gif&amp;s=38'
       alt='Karlsruher Institut für Technologie'
       title='Karlsruher Institut für Technologie'
       width='228'
       height='18' />
  </subpod>
 </pod>
 <pod title='Basic information'
     scanner='Data'
     id='BasicInformation:UniversityData'
     position='200'
     error='false'
     numsubpods='1'>
  <subpod title=''>
   <plaintext>location | Karlsruhe, Baden-Wurttemberg, Germany  (population: 286327)
website | www.kit.edu
year founded | 1825  (189 years ago)</plaintext>
   <img src='http://www4b.wolframalpha.com/Calculate/MSP/MSP72921d3fce4ibfii72fg000060cg5igah4eca9c6?MSPStoreType=image/gif&amp;s=38'
       alt='location | Karlsruhe, Baden-Wurttemberg, Germany  (population: 286327)
website | www.kit.edu
year founded | 1825  (189 years ago)'
       title='location | Karlsruhe, Baden-Wurttemberg, Germany  (population: 286327)
website | www.kit.edu
year founded | 1825  (189 years ago)'
       width='496'
       height='117' />
  </subpod>
  <states count='1'>
   <state name='More'
       input='BasicInformation:UniversityData__More' />
  </states>
 </pod>
 <assumptions count='1'>
  <assumption type='Clash'
      word='KIT'
      template='Assuming &quot;${word}&quot; is ${desc1}. Use as ${desc2} instead'
      count='9'>
   <value name='University'
       desc='a university'
       input='*C.KIT-_*University-' />
   <value name='Financial'
       desc='a financial entity'
       input='*C.KIT-_*Financial-' />
   <value name='Airport'
       desc='an airport'
       input='*C.KIT-_*Airport-' />
   <value name='Protein'
       desc='a protein'
       input='*C.KIT-_*Protein-' />
   <value name='GivenName'
       desc='a given name'
       input='*C.KIT-_*GivenName-' />
   <value name='Person'
       desc='a person'
       input='*C.KIT-_*Person-' />
   <value name='AcronymClass'
       desc='an acronym'
       input='*C.KIT-_*AcronymClass-' />
   <value name='FileFormat'
       desc='a file format'
       input='*C.KIT-_*FileFormat-' />
   <value name='Word'
       desc='a word'
       input='*C.KIT-_*Word-' />
  </assumption>
 </assumptions>
 <sources count='2'>
  <source url='http://www.wolframalpha.com/sources/CityDataSourceInformationNotes.html'
      text='City data' />
  <source url='http://www.wolframalpha.com/sources/UniversityDataSourceInformationNotes.html'
      text='University data' />
 </sources>
</queryresult>
`

}

