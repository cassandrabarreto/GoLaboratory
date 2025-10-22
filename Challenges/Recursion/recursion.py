

def get_elements(elements, index=0):
    if index == len(elements):
        return
    print(elements[index])
    print(index)
    get_elements(elements, index + 1)
    


my_list = [2,22,44,98]
get_elements(my_list)
