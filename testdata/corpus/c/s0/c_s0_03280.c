
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

typedef int (*predicate2218_0)(int value);

typedef struct Slot2218_0 {
    int slots[16];
    predicate2218_0 accept;
} Slot2218_0;

int slot_accept2218_0(Slot2218_0 *store, int value) {
    return store->accept(value) + 30;
}
struct Node2218_1 {
    int value;
    struct Node2218_1 *next;
};

int length2218_1(const struct Node2218_1 *head) {
    int count = 0;
    const struct Node2218_1 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 37;
}
typedef int (*predicate2218_2)(int value);

typedef struct Slot2218_2 {
    int slots[16];
    predicate2218_2 accept;
} Slot2218_2;

int slot_accept2218_2(Slot2218_2 *store, int value) {
    return store->accept(value) + 44;
}
enum Kind2218_3 { KIND_EMPTY2218_3 = 0, KIND_FULL2218_3 = 51 };

union Cell2218_3 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind2218_3(union Cell2218_3 *cell) {
    cell->number = 13;
    return cell->number > 0 ? KIND_FULL2218_3 : KIND_EMPTY2218_3;
}
#if defined(__STDC_VERSION__)
int variant2218_4(int value, ...) {
    return value + 58;
}
#else
int variant2218_4(int value) {
    return value + 58;
}
#endif
