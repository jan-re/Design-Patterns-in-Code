from iterable import Collection

# The simple example above does not really illustrate the potential of the iterator.
# Various kinds of iterators could iterate over complex collections in various ways.
# Like in a list of lists of strings.
# An iterator could give you all items in the first list, then all the items in the second list, etc.
# A different iterator could give you the first position in all lists, then second, etc.

if __name__ == "__main__":
    items = ["1st", "2nd", "3rd", "4th"]

    collection = Collection(items)

    # Iterate over items
    for item in collection:
        print(item)

    print("\nNow iterating in reverse order:\n")

    # Iteratore in reverse order
    for item in collection.get_reverse_iterator():
        print(item)