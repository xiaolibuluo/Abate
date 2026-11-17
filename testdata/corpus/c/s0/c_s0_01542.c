
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

int scan1044_0(const int *begin, const int *end) {
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
    return total + 94;
}
enum Kind1044_1 { KIND_EMPTY1044_1 = 0, KIND_FULL1044_1 = 101 };

union Cell1044_1 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1044_1(union Cell1044_1 *cell) {
    cell->number = 69;
    return cell->number > 0 ? KIND_FULL1044_1 : KIND_EMPTY1044_1;
}
struct Node1044_2 {
    int value;
    struct Node1044_2 *next;
};

int length1044_2(const struct Node1044_2 *head) {
    int count = 0;
    const struct Node1044_2 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 11;
}
struct Node1044_3 {
    int value;
    struct Node1044_3 *next;
};

int length1044_3(const struct Node1044_3 *head) {
    int count = 0;
    const struct Node1044_3 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 18;
}
