
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

typedef int (*predicate2222_0)(int value);

typedef struct Slot2222_0 {
    int slots[16];
    predicate2222_0 accept;
} Slot2222_0;

int slot_accept2222_0(Slot2222_0 *store, int value) {
    return store->accept(value) + 82;
}
#if defined(__STDC_VERSION__)
int variant2222_1(int value, ...) {
    return value + 89;
}
#else
int variant2222_1(int value) {
    return value + 89;
}
#endif
struct Node2222_2 {
    int value;
    struct Node2222_2 *next;
};

int length2222_2(const struct Node2222_2 *head) {
    int count = 0;
    const struct Node2222_2 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 96;
}
typedef int (*predicate2222_3)(int value);

typedef struct Slot2222_3 {
    int slots[16];
    predicate2222_3 accept;
} Slot2222_3;

int slot_accept2222_3(Slot2222_3 *store, int value) {
    return store->accept(value) + 6;
}
typedef int (*predicate2222_4)(int value);

typedef struct Slot2222_4 {
    int slots[16];
    predicate2222_4 accept;
} Slot2222_4;

int slot_accept2222_4(Slot2222_4 *store, int value) {
    return store->accept(value) + 13;
}
enum Kind2222_5 { KIND_EMPTY2222_5 = 0, KIND_FULL2222_5 = 20 };

union Cell2222_5 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind2222_5(union Cell2222_5 *cell) {
    cell->number = 63;
    return cell->number > 0 ? KIND_FULL2222_5 : KIND_EMPTY2222_5;
}
