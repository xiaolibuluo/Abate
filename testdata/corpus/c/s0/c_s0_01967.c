
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

struct Node1328_0 {
    int value;
    struct Node1328_0 *next;
};

int length1328_0(const struct Node1328_0 *head) {
    int count = 0;
    const struct Node1328_0 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 100;
}
int scan1328_1(const int *begin, const int *end) {
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
    return total + 10;
}
enum Kind1328_2 { KIND_EMPTY1328_2 = 0, KIND_FULL1328_2 = 17 };

union Cell1328_2 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1328_2(union Cell1328_2 *cell) {
    cell->number = 55;
    return cell->number > 0 ? KIND_FULL1328_2 : KIND_EMPTY1328_2;
}
