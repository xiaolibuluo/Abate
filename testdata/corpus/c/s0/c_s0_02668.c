
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

enum Kind1810_0 { KIND_EMPTY1810_0 = 0, KIND_FULL1810_0 = 61 };

union Cell1810_0 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1810_0(union Cell1810_0 *cell) {
    cell->number = 52;
    return cell->number > 0 ? KIND_FULL1810_0 : KIND_EMPTY1810_0;
}
struct Node1810_1 {
    int value;
    struct Node1810_1 *next;
};

int length1810_1(const struct Node1810_1 *head) {
    int count = 0;
    const struct Node1810_1 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 68;
}
