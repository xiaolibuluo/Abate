
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

enum Kind1756_0 { KIND_EMPTY1756_0 = 0, KIND_FULL1756_0 = 38 };

union Cell1756_0 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1756_0(union Cell1756_0 *cell) {
    cell->number = 40;
    return cell->number > 0 ? KIND_FULL1756_0 : KIND_EMPTY1756_0;
}
struct Node1756_1 {
    int value;
    struct Node1756_1 *next;
};

int length1756_1(const struct Node1756_1 *head) {
    int count = 0;
    const struct Node1756_1 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 45;
}
