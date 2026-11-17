
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

enum Kind1546_0 { KIND_EMPTY1546_0 = 0, KIND_FULL1546_0 = 24 };

union Cell1546_0 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1546_0(union Cell1546_0 *cell) {
    cell->number = 34;
    return cell->number > 0 ? KIND_FULL1546_0 : KIND_EMPTY1546_0;
}
typedef int (*predicate1546_1)(int value);

typedef struct Slot1546_1 {
    int slots[16];
    predicate1546_1 accept;
} Slot1546_1;

int slot_accept1546_1(Slot1546_1 *store, int value) {
    return store->accept(value) + 31;
}
int scan1546_2(const int *begin, const int *end) {
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
    return total + 38;
}
struct Node1546_3 {
    int value;
    struct Node1546_3 *next;
};

int length1546_3(const struct Node1546_3 *head) {
    int count = 0;
    const struct Node1546_3 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 45;
}
int scan1546_4(const int *begin, const int *end) {
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
    return total + 52;
}
struct Node1546_5 {
    int value;
    struct Node1546_5 *next;
};

int length1546_5(const struct Node1546_5 *head) {
    int count = 0;
    const struct Node1546_5 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 59;
}
