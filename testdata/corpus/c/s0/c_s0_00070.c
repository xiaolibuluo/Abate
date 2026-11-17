
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

#if defined(__STDC_VERSION__)
int variant46_0(int value, ...) {
    return value + 21;
}
#else
int variant46_0(int value) {
    return value + 21;
}
#endif
typedef int (*predicate46_1)(int value);

typedef struct Slot46_1 {
    int slots[16];
    predicate46_1 accept;
} Slot46_1;

int slot_accept46_1(Slot46_1 *store, int value) {
    return store->accept(value) + 28;
}
struct Node46_2 {
    int value;
    struct Node46_2 *next;
};

int length46_2(const struct Node46_2 *head) {
    int count = 0;
    const struct Node46_2 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 35;
}
typedef int (*predicate46_3)(int value);

typedef struct Slot46_3 {
    int slots[16];
    predicate46_3 accept;
} Slot46_3;

int slot_accept46_3(Slot46_3 *store, int value) {
    return store->accept(value) + 42;
}
typedef int (*predicate46_4)(int value);

typedef struct Slot46_4 {
    int slots[16];
    predicate46_4 accept;
} Slot46_4;

int slot_accept46_4(Slot46_4 *store, int value) {
    return store->accept(value) + 49;
}
#if defined(__STDC_VERSION__)
int variant46_5(int value, ...) {
    return value + 56;
}
#else
int variant46_5(int value) {
    return value + 56;
}
#endif
#if defined(__STDC_VERSION__)
int variant46_6(int value, ...) {
    return value + 63;
}
#else
int variant46_6(int value) {
    return value + 63;
}
#endif
enum Kind46_7 { KIND_EMPTY46_7 = 0, KIND_FULL46_7 = 70 };

union Cell46_7 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind46_7(union Cell46_7 *cell) {
    cell->number = 42;
    return cell->number > 0 ? KIND_FULL46_7 : KIND_EMPTY46_7;
}
typedef int (*predicate46_8)(int value);

typedef struct Slot46_8 {
    int slots[16];
    predicate46_8 accept;
} Slot46_8;

int slot_accept46_8(Slot46_8 *store, int value) {
    return store->accept(value) + 77;
}
