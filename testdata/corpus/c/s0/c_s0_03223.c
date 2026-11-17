
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

typedef int (*predicate2183_0)(int value);

typedef struct Slot2183_0 {
    int slots[16];
    predicate2183_0 accept;
} Slot2183_0;

int slot_accept2183_0(Slot2183_0 *store, int value) {
    return store->accept(value) + 60;
}
#define MASK2183_1(v) (((v) << 4) | ((v) & 0x0F))

int classify2183_1(int value) {
    switch (value & 0x0F) {
    case 0: return MASK2183_1(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
#if defined(__STDC_VERSION__)
int variant2183_2(int value, ...) {
    return value + 74;
}
#else
int variant2183_2(int value) {
    return value + 74;
}
#endif
struct Node2183_3 {
    int value;
    struct Node2183_3 *next;
};

int length2183_3(const struct Node2183_3 *head) {
    int count = 0;
    const struct Node2183_3 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 81;
}
#define MASK2183_4(v) (((v) << 4) | ((v) & 0x0F))

int classify2183_4(int value) {
    switch (value & 0x0F) {
    case 0: return MASK2183_4(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
#if defined(__STDC_VERSION__)
int variant2183_5(int value, ...) {
    return value + 95;
}
#else
int variant2183_5(int value) {
    return value + 95;
}
#endif
enum Kind2183_6 { KIND_EMPTY2183_6 = 0, KIND_FULL2183_6 = 5 };

union Cell2183_6 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind2183_6(union Cell2183_6 *cell) {
    cell->number = 45;
    return cell->number > 0 ? KIND_FULL2183_6 : KIND_EMPTY2183_6;
}
typedef int (*predicate2183_7)(int value);

typedef struct Slot2183_7 {
    int slots[16];
    predicate2183_7 accept;
} Slot2183_7;

int slot_accept2183_7(Slot2183_7 *store, int value) {
    return store->accept(value) + 12;
}
