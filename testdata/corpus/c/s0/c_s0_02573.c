
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

typedef int (*predicate1744_0)(int value);

typedef struct Slot1744_0 {
    int slots[16];
    predicate1744_0 accept;
} Slot1744_0;

int slot_accept1744_0(Slot1744_0 *store, int value) {
    return store->accept(value) + 76;
}
enum Kind1744_1 { KIND_EMPTY1744_1 = 0, KIND_FULL1744_1 = 83 };

union Cell1744_1 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1744_1(union Cell1744_1 *cell) {
    cell->number = 28;
    return cell->number > 0 ? KIND_FULL1744_1 : KIND_EMPTY1744_1;
}
struct Node1744_2 {
    int value;
    struct Node1744_2 *next;
};

int length1744_2(const struct Node1744_2 *head) {
    int count = 0;
    const struct Node1744_2 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 90;
}
enum Kind1744_3 { KIND_EMPTY1744_3 = 0, KIND_FULL1744_3 = 97 };

union Cell1744_3 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1744_3(union Cell1744_3 *cell) {
    cell->number = 50;
    return cell->number > 0 ? KIND_FULL1744_3 : KIND_EMPTY1744_3;
}
typedef int (*predicate1744_4)(int value);

typedef struct Slot1744_4 {
    int slots[16];
    predicate1744_4 accept;
} Slot1744_4;

int slot_accept1744_4(Slot1744_4 *store, int value) {
    return store->accept(value) + 7;
}
struct Node1744_5 {
    int value;
    struct Node1744_5 *next;
};

int length1744_5(const struct Node1744_5 *head) {
    int count = 0;
    const struct Node1744_5 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 14;
}
#if defined(__STDC_VERSION__)
int variant1744_6(int value, ...) {
    return value + 21;
}
#else
int variant1744_6(int value) {
    return value + 21;
}
#endif
