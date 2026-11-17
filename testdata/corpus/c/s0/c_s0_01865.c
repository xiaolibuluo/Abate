
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

int scan1252_0(const int *begin, const int *end) {
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
    return total + 82;
}
typedef int (*predicate1252_1)(int value);

typedef struct Slot1252_1 {
    int slots[16];
    predicate1252_1 accept;
} Slot1252_1;

int slot_accept1252_1(Slot1252_1 *store, int value) {
    return store->accept(value) + 89;
}
int scan1252_2(const int *begin, const int *end) {
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
    return total + 96;
}
enum Kind1252_3 { KIND_EMPTY1252_3 = 0, KIND_FULL1252_3 = 6 };

union Cell1252_3 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1252_3(union Cell1252_3 *cell) {
    cell->number = 22;
    return cell->number > 0 ? KIND_FULL1252_3 : KIND_EMPTY1252_3;
}
typedef int (*predicate1252_4)(int value);

typedef struct Slot1252_4 {
    int slots[16];
    predicate1252_4 accept;
} Slot1252_4;

int slot_accept1252_4(Slot1252_4 *store, int value) {
    return store->accept(value) + 13;
}
struct Node1252_5 {
    int value;
    struct Node1252_5 *next;
};

int length1252_5(const struct Node1252_5 *head) {
    int count = 0;
    const struct Node1252_5 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 20;
}
