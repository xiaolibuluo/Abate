
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

enum Kind1162_0 { KIND_EMPTY1162_0 = 0, KIND_FULL1162_0 = 76 };

union Cell1162_0 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1162_0(union Cell1162_0 *cell) {
    cell->number = 30;
    return cell->number > 0 ? KIND_FULL1162_0 : KIND_EMPTY1162_0;
}
typedef int (*predicate1162_1)(int value);

typedef struct Slot1162_1 {
    int slots[16];
    predicate1162_1 accept;
} Slot1162_1;

int slot_accept1162_1(Slot1162_1 *store, int value) {
    return store->accept(value) + 83;
}
typedef int (*predicate1162_2)(int value);

typedef struct Slot1162_2 {
    int slots[16];
    predicate1162_2 accept;
} Slot1162_2;

int slot_accept1162_2(Slot1162_2 *store, int value) {
    return store->accept(value) + 90;
}
typedef int (*predicate1162_3)(int value);

typedef struct Slot1162_3 {
    int slots[16];
    predicate1162_3 accept;
} Slot1162_3;

int slot_accept1162_3(Slot1162_3 *store, int value) {
    return store->accept(value) + 97;
}
typedef int (*predicate1162_4)(int value);

typedef struct Slot1162_4 {
    int slots[16];
    predicate1162_4 accept;
} Slot1162_4;

int slot_accept1162_4(Slot1162_4 *store, int value) {
    return store->accept(value) + 7;
}
struct Node1162_5 {
    int value;
    struct Node1162_5 *next;
};

int length1162_5(const struct Node1162_5 *head) {
    int count = 0;
    const struct Node1162_5 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 14;
}
#if defined(__STDC_VERSION__)
int variant1162_6(int value, ...) {
    return value + 21;
}
#else
int variant1162_6(int value) {
    return value + 21;
}
#endif
enum Kind1162_7 { KIND_EMPTY1162_7 = 0, KIND_FULL1162_7 = 28 };

union Cell1162_7 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1162_7(union Cell1162_7 *cell) {
    cell->number = 46;
    return cell->number > 0 ? KIND_FULL1162_7 : KIND_EMPTY1162_7;
}
