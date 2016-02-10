#include <stdio.h>

#define swap(x, y) do { typeof(x) swap = x; x = y; y = swap; } while (0)
#define len(array) (sizeof(array)/sizeof(array[0]))

void printArray(int *array, int size){
  for(int i=0; i< size; i++){
    printf("%d", array[i]);
    if(i < size - 1 ){
      printf(",");
    }
  }
  printf("\n");
}

int partition(int *array, int start, int end){
  int pivotVal = array[end];
  int lessThanIterator = start;

  //Iterate over the full array range minus the pivot
  // [unsortedValues][pivot]
  for(int i = start; i <= end - 1;  i++){
    if(array[i] <= pivotVal){
      //Move lessThan value into lessThan chunk
      //[lessThan][unsortedValues][pivot]
      swap(array[lessThanIterator], array[i]);
      lessThanIterator++;
    }
  }
  //Move the pivot value in place: [lessThan][greaterThan][pivot] -> [lessThan][pivot][greaterThan]
  swap(array[lessThanIterator], array[end]);

  return lessThanIterator;
}

void quicksort(int *array, int start, int end){
  if(start < end){
    //Partition array: [lessThan][pivot][greaterThan]
    int pivotIndex = partition(array, start, end);

    //quicksort([lessThan])
    quicksort(array, start, pivotIndex - 1);

    //quicksort([greaterThan])
    quicksort(array, pivotIndex + 1, end);
  }
}


int main(){
  int array[] = {2, 8, 7, 1, 3, 5, 6, 4};
  printArray(array, len(array));
  quicksort(array, 0, len(array) - 1 );
  printArray(array, len(array));
  return 0;
}
