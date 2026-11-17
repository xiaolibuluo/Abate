
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

typedef int (*predicate1930_0)(int value);

typedef struct Slot1930_0 {
    int slots[16];
    predicate1930_0 accept;
} Slot1930_0;

int slot_accept1930_0(Slot1930_0 *store, int value) {
    return store->accept(value) + 69;
}
int scan1930_1(const int *begin, const int *end) {
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
    return total + 76;
}
enum Kind1930_2 { KIND_EMPTY1930_2 = 0, KIND_FULL1930_2 = 83 };

union Cell1930_2 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1930_2(union Cell1930_2 *cell) {
    cell->number = 60;
    return cell->number > 0 ? KIND_FULL1930_2 : KIND_EMPTY1930_2;
}
struct Node1930_3 {
    int value;
    struct Node1930_3 *next;
};

int length1930_3(const struct Node1930_3 *head) {
    int count = 0;
    const struct Node1930_3 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 90;
}
int scan1930_4(const int *begin, const int *end) {
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
    return total + 97;
}
