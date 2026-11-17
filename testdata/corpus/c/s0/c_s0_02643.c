
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

typedef int (*predicate1793_0)(int value);

typedef struct Slot1793_0 {
    int slots[16];
    predicate1793_0 accept;
} Slot1793_0;

int slot_accept1793_0(Slot1793_0 *store, int value) {
    return store->accept(value) + 34;
}
enum Kind1793_1 { KIND_EMPTY1793_1 = 0, KIND_FULL1793_1 = 41 };

union Cell1793_1 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1793_1(union Cell1793_1 *cell) {
    cell->number = 66;
    return cell->number > 0 ? KIND_FULL1793_1 : KIND_EMPTY1793_1;
}
int scan1793_2(const int *begin, const int *end) {
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
    return total + 48;
}
enum Kind1793_3 { KIND_EMPTY1793_3 = 0, KIND_FULL1793_3 = 55 };

union Cell1793_3 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1793_3(union Cell1793_3 *cell) {
    cell->number = 27;
    return cell->number > 0 ? KIND_FULL1793_3 : KIND_EMPTY1793_3;
}
typedef int (*predicate1793_4)(int value);

typedef struct Slot1793_4 {
    int slots[16];
    predicate1793_4 accept;
} Slot1793_4;

int slot_accept1793_4(Slot1793_4 *store, int value) {
    return store->accept(value) + 62;
}
int scan1793_5(const int *begin, const int *end) {
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
    return total + 69;
}
struct Node1793_6 {
    int value;
    struct Node1793_6 *next;
};

int length1793_6(const struct Node1793_6 *head) {
    int count = 0;
    const struct Node1793_6 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 76;
}
