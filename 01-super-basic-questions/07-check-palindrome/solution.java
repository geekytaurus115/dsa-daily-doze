import java.util.Scanner;

class Main{
    public static void main(String[] args){
        Scanner myObj = new Scanner(System.in);

        System.out.println("Enter a number: ");
        int num = myObj.nextInt();

        int original = num;
        int reverse = 0;
        

        while (num != 0){
            int digit = num % 10;
            reverse = reverse * 10 + digit;
            num /= 10;           
        }

        if (original == reverse){
            System.out.println("Palindrom");
        }else{
            System.out.println("Not Palindrom");
        }

    }
}