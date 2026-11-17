
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

enum Kind1216_0 { KIND_EMPTY1216_0 = 0, KIND_FULL1216_0 = 99 };

union Cell1216_0 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1216_0(union Cell1216_0 *cell) {
    cell->number = 42;
    return cell->number > 0 ? KIND_FULL1216_0 : KIND_EMPTY1216_0;
}
typedef int (*predicate1216_1)(int value);

typedef struct Slot1216_1 {
    int slots[16];
    predicate1216_1 accept;
} Slot1216_1;

int slot_accept1216_1(Slot1216_1 *store, int value) {
    return store->accept(value) + 9;
}
enum Kind1216_2 { KIND_EMPTY1216_2 = 0, KIND_FULL1216_2 = 16 };

union Cell1216_2 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1216_2(union Cell1216_2 *cell) {
    cell->number = 64;
    return cell->number > 0 ? KIND_FULL1216_2 : KIND_EMPTY1216_2;
}
typedef int (*predicate1216_3)(int value);

typedef struct Slot1216_3 {
    int slots[16];
    predicate1216_3 accept;
} Slot1216_3;

int slot_accept1216_3(Slot1216_3 *store, int value) {
    return store->accept(value) + 23;
}
int scan1216_4(const int *begin, const int *end) {
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
    return total + 30;
}
#if defined(__STDC_VERSION__)
int variant1216_5(int value, ...) {
    return value + 37;
}
#else
int variant1216_5(int value) {
    return value + 37;
}
#endif
struct Node1216_6 {
    int value;
    struct Node1216_6 *next;
};

int length1216_6(const struct Node1216_6 *head) {
    int count = 0;
    const struct Node1216_6 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 44;
}
enum Kind1216_7 { KIND_EMPTY1216_7 = 0, KIND_FULL1216_7 = 51 };

union Cell1216_7 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1216_7(union Cell1216_7 *cell) {
    cell->number = 58;
    return cell->number > 0 ? KIND_FULL1216_7 : KIND_EMPTY1216_7;
}
