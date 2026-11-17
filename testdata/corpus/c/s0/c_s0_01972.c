
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

enum Kind1332_0 { KIND_EMPTY1332_0 = 0, KIND_FULL1332_0 = 55 };

union Cell1332_0 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1332_0(union Cell1332_0 *cell) {
    cell->number = 61;
    return cell->number > 0 ? KIND_FULL1332_0 : KIND_EMPTY1332_0;
}
int scan1332_1(const int *begin, const int *end) {
    int total = 0;
    const int *cursor = begin;
    for (; cursor != end; ++cursor) {
        if (*cursor >= 0) {
            total += *cursor;
        }
    }
    do {
        --cursor;
    } while (cursor != begin && total > 0);
    return total + 62;
}
struct Node1332_2 {
    int value;
    struct Node1332_2 *next;
};

int length1332_2(const struct Node1332_2 *head) {
    int count = 0;
    const struct Node1332_2 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 69;
}
