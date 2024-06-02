import java.util.Scanner;

class Solution {
    public static void main(String[] args) {
        Scanner sc = new Scanner(System.in);

        int arrSize = sc.nextInt();  

        int[] arr = new int[arrSize]; 

        for (int i = 0; i < arrSize; i++) {
            arr[i] = sc.nextInt();
        }

        int sum = 0;
        for (int i = 0; i < arrSize; i++) {
            sum += arr[i];
        }

        System.out.print("Sum: "+sum);

    }
}
