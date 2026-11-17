
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

typedef int (*predicate854_0)(int value);

typedef struct Slot854_0 {
    int slots[16];
    predicate854_0 accept;
} Slot854_0;

int slot_accept854_0(Slot854_0 *store, int value) {
    return store->accept(value) + 49;
}
enum Kind854_1 { KIND_EMPTY854_1 = 0, KIND_FULL854_1 = 56 };

union Cell854_1 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind854_1(union Cell854_1 *cell) {
    cell->number = 20;
    return cell->number > 0 ? KIND_FULL854_1 : KIND_EMPTY854_1;
}
struct Node854_2 {
    int value;
    struct Node854_2 *next;
};

int length854_2(const struct Node854_2 *head) {
    int count = 0;
    const struct Node854_2 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 63;
}
#if defined(__STDC_VERSION__)
int variant854_3(int value, ...) {
    return value + 70;
}
#else
int variant854_3(int value) {
    return value + 70;
}
#endif
typedef int (*predicate854_4)(int value);

typedef struct Slot854_4 {
    int slots[16];
    predicate854_4 accept;
} Slot854_4;

int slot_accept854_4(Slot854_4 *store, int value) {
    return store->accept(value) + 77;
}
