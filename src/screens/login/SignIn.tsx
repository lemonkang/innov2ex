import { View,StatusBar, Image,StyleSheet, TextInput, Button, Text } from "react-native"
import { Colors } from "react-native/Libraries/NewAppScreen";

const SignInScreen=({navigation}:any)=>{
return <View style={styles.mainContainer}>
 <StatusBar backgroundColor="black" barStyle='light-content' />
    <Image style={styles.logo} source={require('../../assets/images/icons/app/arrow/light.png')}/>
     <View style={styles.inputbtn}>
     <TextInput  style={{backgroundColor:'white',height:38}} placeholder='输入授权码'/>
     <Button  title="使用Token授权登录" />
     </View>
     <View style={styles.bottomText}>
     <Text style={{color:'white'}}>登录既表示你同意</Text>  
     <Text style={{color:'blue'}}>服务条款</Text>  
   
     <Text style={{color:'white'}}>和</Text>
     <Text style={{color:'blue'}}>隐私政策</Text>  

     </View>
   
</View>
}
const styles = StyleSheet.create({
    mainContainer:{
        flex:1,
        backgroundColor:'black',
        position:'relative'
    },
    logo:{
       width:150*0.9,
       height:150*0.9,
       position:'absolute',
       top:'50%',
       left:'50%',
       marginLeft:-1/2*150*0.9,
       marginTop:-1/2*150*0.9
    },
    inputbtn:{
        height:100,
        width:'80%',
        justifyContent:'space-between',
        position:'absolute',
        bottom:'15%',
        margin:'auto',
        alignSelf:'center'

    },
    bottomText:{
        display:"flex",
        flexDirection:'row',
        fontSize:12,
        color:'white',
        alignItems:'center',
        justifyContent:'center',
        position:'absolute',
        bottom:0,
        
        marginBottom:4,
        alignSelf:'center'
        
        
       
    }

  });
  
export default SignInScreen