import * as React from 'react';
import { Button, View, Text } from 'react-native';
import { NavigationContainer } from '@react-navigation/native';
import { createNativeStackNavigator } from '@react-navigation/native-stack';
import { ROUTES } from './navigation/routes';
import SignInScreen from './screens/login/SignIn';
const StackNavigator = createNativeStackNavigator();


function App() {
  return (
    <NavigationContainer>
      <StackNavigator.Navigator>
        <StackNavigator.Screen
          name={ROUTES.SignIn}
          component={SignInScreen}
        />
    
      </StackNavigator.Navigator>
    </NavigationContainer>
  );
}

export default App;