
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

enum Kind1781_0 { KIND_EMPTY1781_0 = 0, KIND_FULL1781_0 = 72 };

union Cell1781_0 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1781_0(union Cell1781_0 *cell) {
    cell->number = 32;
    return cell->number > 0 ? KIND_FULL1781_0 : KIND_EMPTY1781_0;
}
struct Node1781_1 {
    int value;
    struct Node1781_1 *next;
};

int length1781_1(const struct Node1781_1 *head) {
    int count = 0;
    const struct Node1781_1 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 79;
}
typedef int (*predicate1781_2)(int value);

typedef struct Slot1781_2 {
    int slots[16];
    predicate1781_2 accept;
} Slot1781_2;

int slot_accept1781_2(Slot1781_2 *store, int value) {
    return store->accept(value) + 86;
}
typedef int (*predicate1781_3)(int value);

typedef struct Slot1781_3 {
    int slots[16];
    predicate1781_3 accept;
} Slot1781_3;

int slot_accept1781_3(Slot1781_3 *store, int value) {
    return store->accept(value) + 93;
}
struct Node1781_4 {
    int value;
    struct Node1781_4 *next;
};

int length1781_4(const struct Node1781_4 *head) {
    int count = 0;
    const struct Node1781_4 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 100;
}
