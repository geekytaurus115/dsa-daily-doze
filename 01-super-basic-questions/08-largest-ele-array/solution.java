import java.util.Scanner;

class Main{
    public static void main(String[] args){
        Scanner sc = new Scanner(System.in);

        System.out.print("Enter the size of the array: ");
        int num = sc.nextInt();

        int[] arr = new int[num];

        for (int i = 0; i < num; i++){
            arr[i] = sc.nextInt();
        }

        int maxEle = arr[0];

        for (int i = 0; i < num; i++){
            if (arr[i]> maxEle){
                maxEle = arr[i];
            }
        }

        System.out.println(maxEle);

    }
}