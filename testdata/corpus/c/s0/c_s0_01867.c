
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

struct Node1254_0 {
    int value;
    struct Node1254_0 *next;
};

int length1254_0(const struct Node1254_0 *head) {
    int count = 0;
    const struct Node1254_0 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 11;
}
int scan1254_1(const int *begin, const int *end) {
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
    return total + 18;
}
int scan1254_2(const int *begin, const int *end) {
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
    return total + 25;
}
enum Kind1254_3 { KIND_EMPTY1254_3 = 0, KIND_FULL1254_3 = 32 };

union Cell1254_3 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1254_3(union Cell1254_3 *cell) {
    cell->number = 36;
    return cell->number > 0 ? KIND_FULL1254_3 : KIND_EMPTY1254_3;
}
enum Kind1254_4 { KIND_EMPTY1254_4 = 0, KIND_FULL1254_4 = 39 };

union Cell1254_4 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1254_4(union Cell1254_4 *cell) {
    cell->number = 47;
    return cell->number > 0 ? KIND_FULL1254_4 : KIND_EMPTY1254_4;
}
typedef int (*predicate1254_5)(int value);

typedef struct Slot1254_5 {
    int slots[16];
    predicate1254_5 accept;
} Slot1254_5;

int slot_accept1254_5(Slot1254_5 *store, int value) {
    return store->accept(value) + 46;
}
