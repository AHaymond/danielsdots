Vim�UnDo� F��Au�lK���d��.O��@���3G��,"                                     V�	(    _�                        
    ����                                                                                                                                                                                                                                                                                                                               
                 v       V��     �                type IssueChangelogItem struct {5�_�                           ����                                                                                                                                                                                                                                                                                                                               
                 v       V��     �      	         	Field       string5�_�                           ����                                                                                                                                                                                                                                                                                                                               
                 v       V��     �      	         	IssueID       string5�_�                            ����                                                                                                                                                                                                                                                                                                                                                V       V��     �                import (   	"database/sql"   )5�_�                            ����                                                                                                                                                                                                                                                                                                                                                V       V��     �                 5�_�                            ����                                                                                                                                                                                                                                                                                                                                                  V        V��     �                	FieldID     sql.NullString   	From        sql.NullString    	ChangelogID int8 `db:"history"`   	to          sql.NullString5�_�                            ����                                                                                                                                                                                                                                                                                                                                                  V        V��     �             �             5�_�      	                     ����                                                                                                                                                                                                                                                                                                                                                  V        V��     �               	IssueID     int85�_�      
           	           ����                                                                                                                                                                                                                                                                                                                                                  V        V��    �                  package models       type IssueFixVersion struct {   	IssueID     int8   	Version     int8   }5�_�   	              
          ����                                                                                                                                                                                                                                                                                                                                                  V        V�	     �               	IssueID int85�_�   
                        ����                                                                                                                                                                                                                                                                                                                                                  V        V�		    �                  package models       type IssueFixVersion struct {   	IssueID int8 `db:"issue"`   	Version int8   }5�_�                           ����                                                                                                                                                                                                                                                                                                                                                  V        V�	     �               	Version int85�_�                           ����                                                                                                                                                                                                                                                                                                                                                  V        V�	     �               	VersionID int85�_�                           ����                                                                                                                                                                                                                                                                                                                                                  V        V�	&     �               	VersionID int8 `db:"version"5�_�                             ����                                                                                                                                                                                                                                                                                                                                                  V        V�	'    �                  package models       type IssueFixVersion struct {   	IssueID int8 `db:"issue"`   	VersionID int8 `db:"version"`   }5��