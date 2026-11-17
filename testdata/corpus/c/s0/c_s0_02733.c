
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

enum Kind1858_0 { KIND_EMPTY1858_0 = 0, KIND_FULL1858_0 = 6 };

union Cell1858_0 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1858_0(union Cell1858_0 *cell) {
    cell->number = 22;
    return cell->number > 0 ? KIND_FULL1858_0 : KIND_EMPTY1858_0;
}
struct Node1858_1 {
    int value;
    struct Node1858_1 *next;
};

int length1858_1(const struct Node1858_1 *head) {
    int count = 0;
    const struct Node1858_1 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 13;
}
