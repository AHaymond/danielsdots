Vim�UnDo� >��c��	��E��i$U]�X�DkǠ��      *  docker.mstry.io/etl/ios_app_imports bash                             V7�7   
 _�                             ����                                                                                                                                                                                                                                                                                                                                                  V        V7�     �                $d=$(date --date="-1 month" +"%Y_%m")   Vgsutil_auth_token="1/VZRBj5x3DOC2pRzqSRFPRkEMZB7-7U4s_JPYGz88i0TBactUREZofsF9C7PrpE-j"   1gs_bucket="pubsite_prod_rev_14866784520087192582"5�_�                       "    ����                                                                                                                                                                                                                                                                                                                                                  V        V7�'    �               .aws_s3_path="mc-metrics/app_store/android/$d/"5�_�                       &    ����                                                                                                                                                                                                                                                                                                                                                  V        V7�.     �               *aws_s3_path="mc-metrics/app_store/ios/$d/"5�_�                       &    ����                                                                                                                                                                                                                                                                                                                                                  V        V7�.     �               )aws_s3_path="mc-metrics/app_store/ios/d/"5�_�                       &    ����                                                                                                                                                                                                                                                                                                                                                  V        V7�/    �               (aws_s3_path="mc-metrics/app_store/ios//"5�_�                    
       ����                                                                                                                                                                                                                                                                                                                                                  V        V7�5    �   	              -e "HOME=/app" \5�_�                           ����                                                                                                                                                                                                                                                                                                                                                  V        V7�9    �   
              -v $HOME/.aws:/app/.aws \5�_�      	                     ����                                                                                                                                                                                                                                                                                                                                                  V        V7�C    �                 ,  docker.mstry.io/etl/android_app_stats bash5�_�      
           	           ����                                                                                                                                                                                                                                                                                                                                                  V        V7�    �                  -e "GS_BUCKET=$gs_bucket" \   )  -e "GS_AUTH_TOKEN=$gsutil_auth_token" \5�_�   	              
           ����                                                                                                                                                                                                                                                                                                                                                             V7��     �                 �         
    5�_�   
                        ����                                                                                                                                                                                                                                                                                                                                                             V7��     �                 -e "IOS_ACCOUNT5�_�                            ����                                                                                                                                                                                                                                                                                                                                                             V7��    �                �             5�_�                           ����                                                                                                                                                                                                                                                                                                                                                             V7�     �                 �             5�_�                           ����                                                                                                                                                                                                                                                                                                                                                             V7�     �                 -e "IOS_ACCOUNT=85200725"5�_�                           ����                                                                                                                                                                                                                                                                                                                                                             V7�     �                 -e "IOS_PASSWORD=teachme123"5�_�                       !    ����                                                                                                                                                                                                                                                                                                                                                             V7�     �               !  -e "IOS_ACCOUNT=masteryconnect"5�_�                       
    ����                                                                                                                                                                                                                                                                                                                                                             V7�   	 �               #  -e "IOS_ACCOUNT=masteryconnect" \5�_�                            ����                                                                                                                                                                                                                                                                                                                                                             V7�6   
 �                 *  docker.mstry.io/etl/ios_app_imports bash5��