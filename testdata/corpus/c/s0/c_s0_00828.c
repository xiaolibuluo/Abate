
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

enum Kind571_0 { KIND_EMPTY571_0 = 0, KIND_FULL571_0 = 56 };

union Cell571_0 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind571_0(union Cell571_0 *cell) {
    cell->number = 41;
    return cell->number > 0 ? KIND_FULL571_0 : KIND_EMPTY571_0;
}
#if defined(__STDC_VERSION__)
int variant571_1(int value, ...) {
    return value + 63;
}
#else
int variant571_1(int value) {
    return value + 63;
}
#endif
#define MASK571_2(v) (((v) << 4) | ((v) & 0x0F))

int classify571_2(int value) {
    switch (value & 0x0F) {
    case 0: return MASK571_2(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
struct Node571_3 {
    int value;
    struct Node571_3 *next;
};

int length571_3(const struct Node571_3 *head) {
    int count = 0;
    const struct Node571_3 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 77;
}
#define MASK571_4(v) (((v) << 4) | ((v) & 0x0F))

int classify571_4(int value) {
    switch (value & 0x0F) {
    case 0: return MASK571_4(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
typedef int (*predicate571_5)(int value);

typedef struct Slot571_5 {
    int slots[16];
    predicate571_5 accept;
} Slot571_5;

int slot_accept571_5(Slot571_5 *store, int value) {
    return store->accept(value) + 91;
}
#if defined(__STDC_VERSION__)
int variant571_6(int value, ...) {
    return value + 98;
}
#else
int variant571_6(int value) {
    return value + 98;
}
#endif
