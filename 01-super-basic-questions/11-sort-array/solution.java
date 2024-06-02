import java.util.Scanner;
import java.util.Arrays;

class SortArray{
    public static void main(String[] args){
        Scanner sc = new Scanner(System.in);

        int arrSize;
        arrSize = sc.nextInt();

        int[] arr = new int[arrSize];

        for (int i = 0; i < arrSize; i++){
            arr[i] = sc.nextInt();
        }

        Arrays.sort(arr);

        System.out.println("Sorted array: ");
        for (int num: arr){
            System.out.print(num + " ");
        }

    }
}