import java.util.Scanner;

class Solution{

    public static int[] reverseArr(int[] arr){

        for (int i = 0; i < arr.length / 2; i++){
            int temp = arr[i];
            arr[i] = arr[arr.length - 1 - i];
            arr[arr.length - 1 - i] = temp;

        }


        return arr;
    }


    public static void main(String[] args){

        Scanner sc = new Scanner(System.in);

        int arrSize = sc.nextInt();
        int[] arr = new int[arrSize];

        for (int i = 0; i < arrSize; i++){
            arr[i] = sc.nextInt();
        }

        int[] reversedArr = reverseArr(arr);

        System.out.print("Reversed Array: ");
        for (int i = 0; i < reversedArr.length; i++) {
            System.out.print(reversedArr[i] + " ");
        }
        System.out.println();


    }
}