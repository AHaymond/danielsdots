Vim�UnDo� G 8�̌�*��*+6�����g�>L��(U                                       W4�   
 _�                             ����                                                                                                                                                                                                                                                                                                                                                             W4��     �                   5�_�                           ����                                                                                                                                                                                                                                                                                                                                                             W4��    �                  package worker5�_�                           ����                                                                                                                                                                                                                                                                                                                                                             W4�4     �               �               5�_�                            ����                                                                                                                                                                                                                                                                                                                                                             W4�5     �             5�_�                            ����                                                                                                                                                                                                                                                                                                                                                             W4�6    �                  package worker       %var WorkerQueue chan chan WorkRequest5�_�                            ����                                                                                                                                                                                                                                                                                                                                                             W4�C     �               �               5�_�                            ����                                                                                                                                                                                                                                                                                                                                                             W4�C     �             5�_�                            ����                                                                                                                                                                                                                                                                                                                                                             W4�D    �                  package worker       %var WorkerQueue chan chan WorkRequest       $func StartDispatcher(nworkers int) {   V	// First, initialize the channel we are going to put the workers' work channels into.   5	WorkerQueue := make(chan chan WorkRequest, nworkers)       #	// Now, create all of our workers.    	for i := 0; i < nworkers; i++ {   %		fmt.Println("Starting worker", i+1)   '		worker := NewWorker(i+1, WorkerQueue)   		worker.Start()   	}       	go func() {   		for {   			select {   			case work := <-WorkQueue:   (				fmt.Println("Received work request")   				go func() {   					worker := <-WorkerQueue       ,					fmt.Println("Dispatching work request")   					worker <- work   				}()   			}   		}   	}()   }5�_�         	                  ����                                                                                                                                                                                                                                                                                                                                                             W4�     �             �             5�_�                            ����                                                                                                                                                                                                                                                                                                                                                             W4Ҁ     �             5�_�                            ����                                                                                                                                                                                                                                                                                                                                                             W4҂    �                   package worker       import "github.com/apex/log"       %var WorkerQueue chan chan WorkRequest       $func StartDispatcher(nworkers int) {   V	// First, initialize the channel we are going to put the workers' work channels into.   5	WorkerQueue := make(chan chan WorkRequest, nworkers)       #	// Now, create all of our workers.    	for i := 0; i < nworkers; i++ {   %		fmt.Println("Starting worker", i+1)   '		worker := NewWorker(i+1, WorkerQueue)   		worker.Start()   	}       	go func() {   		for {   			select {   			case work := <-WorkQueue:   (				fmt.Println("Received work request")   				go func() {   					worker := <-WorkerQueue       ,					fmt.Println("Dispatching work request")   					worker <- work   				}()   			}   		}   	}()   }5�_�                          ����                                                                                                                                                                                                                                                                                                                                                             W4ҋ     �                (				fmt.Println("Received work request")5�_�                           ����                                                                                                                                                                                                                                                                                                                                                             W4Ҏ    �                   package worker       import "github.com/apex/log"       %var WorkerQueue chan chan WorkRequest       $func StartDispatcher(nworkers int) {   V	// First, initialize the channel we are going to put the workers' work channels into.   5	WorkerQueue := make(chan chan WorkRequest, nworkers)       #	// Now, create all of our workers.    	for i := 0; i < nworkers; i++ {   %		fmt.Println("Starting worker", i+1)   '		worker := NewWorker(i+1, WorkerQueue)   		worker.Start()   	}       	go func() {   		for {   			select {   			case work := <-WorkQueue:   %				log.Info("Received work request")   				go func() {   					worker := <-WorkerQueue       ,					fmt.Println("Dispatching work request")   					worker <- work   				}()   			}   		}   	}()   }5�_�                           ����                                                                                                                                                                                                                                                                                                                                                             W4Ғ     �                ,					fmt.Println("Dispatching work request")5�_�                           ����                                                                                                                                                                                                                                                                                                                                                             W4Җ    �                   package worker       import "github.com/apex/log"       %var WorkerQueue chan chan WorkRequest       $func StartDispatcher(nworkers int) {   V	// First, initialize the channel we are going to put the workers' work channels into.   5	WorkerQueue := make(chan chan WorkRequest, nworkers)       #	// Now, create all of our workers.    	for i := 0; i < nworkers; i++ {   %		fmt.Println("Starting worker", i+1)   '		worker := NewWorker(i+1, WorkerQueue)   		worker.Start()   	}       	go func() {   		for {   			select {   			case work := <-WorkQueue:   %				log.Info("Received work request")   				go func() {   					worker := <-WorkerQueue       )					log.Info("Dispatching work request")   					worker <- work   				}()   			}   		}   	}()   }5�_�                           ����                                                                                                                                                                                                                                                                                                                                                             W4Ҝ     �                %		fmt.Println("Starting worker", i+1)5�_�                       
    ����                                                                                                                                                                                                                                                                                                                                                             W4ҟ    �                   package worker       import "github.com/apex/log"       %var WorkerQueue chan chan WorkRequest       $func StartDispatcher(nworkers int) {   V	// First, initialize the channel we are going to put the workers' work channels into.   5	WorkerQueue := make(chan chan WorkRequest, nworkers)       #	// Now, create all of our workers.    	for i := 0; i < nworkers; i++ {   #		log.Infof("Starting worker", i+1)   '		worker := NewWorker(i+1, WorkerQueue)   		worker.Start()   	}       	go func() {   		for {   			select {   			case work := <-WorkQueue:   %				log.Info("Received work request")   				go func() {   					worker := <-WorkerQueue       )					log.Info("Dispatching work request")   					worker <- work   				}()   			}   		}   	}()   }5�_�                           ����                                                                                                                                                                                                                                                                                                                                                             W4�     �                #		log.Infof("Starting worker", i+1)5�_�                           ����                                                                                                                                                                                                                                                                                                                                                             W4�    �                   package worker       import "github.com/apex/log"       %var WorkerQueue chan chan WorkRequest       $func StartDispatcher(nworkers int) {   V	// First, initialize the channel we are going to put the workers' work channels into.   5	WorkerQueue := make(chan chan WorkRequest, nworkers)       #	// Now, create all of our workers.    	for i := 0; i < nworkers; i++ {   %		log.Infof("Starting worker\n", i+1)   '		worker := NewWorker(i+1, WorkerQueue)   		worker.Start()   	}       	go func() {   		for {   			select {   			case work := <-WorkQueue:   %				log.Info("Received work request")   				go func() {   					worker := <-WorkerQueue       )					log.Info("Dispatching work request")   					worker <- work   				}()   			}   		}   	}()   }5�_�                          ����                                                                                                                                                                                                                                                                                                                                                             W4�k   	 �                   package worker       import "github.com/apex/log"       %var WorkerQueue chan chan WorkRequest       $func StartDispatcher(nworkers int) {   V	// First, initialize the channel we are going to put the workers' work channels into.   5	WorkerQueue := make(chan chan WorkRequest, nworkers)       #	// Now, create all of our workers.    	for i := 0; i < nworkers; i++ {   %		log.Infof("Starting worker\n", i+1)   '		worker := NewWorker(i+1, WorkerQueue)   		worker.Start()   	}       	go func() {   		for {   			select {   			case work := <-WorkQueue:   %				log.Info("Received work request")   				go func() {   					worker := <-WorkerQueue       )					log.Info("Dispatching work request")   					worker <- work   				}()   			}   		}   	}()   }5�_�                           ����                                                                                                                                                                                                                                                                                                                                                             W4�     �                 package worker5�_�                            ����                                                                                                                                                                                                                                                                                                                                                             W4�   
 �                   package skrape       import "github.com/apex/log"       %var WorkerQueue chan chan WorkRequest       $func StartDispatcher(nworkers int) {   V	// First, initialize the channel we are going to put the workers' work channels into.   5	WorkerQueue := make(chan chan WorkRequest, nworkers)       #	// Now, create all of our workers.    	for i := 0; i < nworkers; i++ {   %		log.Infof("Starting worker\n", i+1)   '		worker := NewWorker(i+1, WorkerQueue)   		worker.Start()   	}       	go func() {   		for {   			select {   			case work := <-WorkQueue:   %				log.Info("Received work request")   				go func() {   					worker := <-WorkerQueue       )					log.Info("Dispatching work request")   					worker <- work   				}()   			}   		}   	}()   }5�_�                           ����                                                                                                                                                                                                                                                                                                                                                             W4�     �                					�         !      					if IsComplete() {   						worker.5�_�                           ����                                                                                                                                                                                                                                                                                                                                                             W4҇     �                				c5�_�                	           ����                                                                                                                                                                                                                                                                                                                                                             W4�n     �                �                    5�_�   	       
                  ����                                                                                                                                                                                                                                                                                                                                                             W4�w     �                i5�_�   	              
           ����                                                                                                                                                                                                                                                                                                                                                             W4�u     �                 �         !           5��