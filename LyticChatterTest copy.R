
  


library(httr)
library(jsonlite)

fun <- function(){
  Age <- readline("How old are you?")
  
  #Age <- as.String(unlist(strsplit(x, ",")))
  
  Age <<- as.numeric(Age)
  
  
  Gender <- readline("Are you male or female?")
  
  Gender <<- as.character(Gender)
  
  
  Symptom <- readline("Describe the symptoms you are experiencing.")
  
  Symptom <<- as.character(Symptom)
  
  #return(list(x))
  
}

Age <- Age * 365



getting <- paste('https://sandbox.healthnavigatorapis.com/3.0/FindCccByPopulation?freetextchiefcomplaints=',Symptom,'&ageindays=',as.character(Age),'&gender=B&languageid=1'
                 ,sep="")




calling <-GET(getting,authenticate("9c86fe28-6c85-47ee-8b75-94149e8cdebb", "2485e0ae-7674-43ef-888c-643942bb4ebe", type="basic"))



calist <- content(calling,as='text')

sympID<- fromJSON((calist))

sympID <- sympID[2,2]











############
a  <- GET("https://sandbox.healthnavigatorapis.com/3.0/FindCccByPopulation?freetextchiefcomplaints=Symptom&ageindays=Age&gender=B&languageid=1",authenticate("9c86fe28-6c85-47ee-8b75-94149e8cdebb", "2485e0ae-7674-43ef-888c-643942bb4ebe", type="basic"))

b <- content(a,as='text')


c <- 






 
 
 