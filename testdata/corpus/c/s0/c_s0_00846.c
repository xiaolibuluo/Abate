
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

struct Node582_0 {
    int value;
    struct Node582_0 *next;
};

int length582_0(const struct Node582_0 *head) {
    int count = 0;
    const struct Node582_0 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 5;
}
#if defined(__STDC_VERSION__)
int variant582_1(int value, ...) {
    return value + 12;
}
#else
int variant582_1(int value) {
    return value + 12;
}
#endif
typedef int (*predicate582_2)(int value);

typedef struct Slot582_2 {
    int slots[16];
    predicate582_2 accept;
} Slot582_2;

int slot_accept582_2(Slot582_2 *store, int value) {
    return store->accept(value) + 19;
}
#define MASK582_3(v) (((v) << 4) | ((v) & 0x0F))

int classify582_3(int value) {
    switch (value & 0x0F) {
    case 0: return MASK582_3(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
enum Kind582_4 { KIND_EMPTY582_4 = 0, KIND_FULL582_4 = 33 };

union Cell582_4 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind582_4(union Cell582_4 *cell) {
    cell->number = 40;
    return cell->number > 0 ? KIND_FULL582_4 : KIND_EMPTY582_4;
}
#define MASK582_5(v) (((v) << 4) | ((v) & 0x0F))

int classify582_5(int value) {
    switch (value & 0x0F) {
    case 0: return MASK582_5(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
enum Kind582_6 { KIND_EMPTY582_6 = 0, KIND_FULL582_6 = 47 };

union Cell582_6 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind582_6(union Cell582_6 *cell) {
    cell->number = 62;
    return cell->number > 0 ? KIND_FULL582_6 : KIND_EMPTY582_6;
}
typedef int (*predicate582_7)(int value);

typedef struct Slot582_7 {
    int slots[16];
    predicate582_7 accept;
} Slot582_7;

int slot_accept582_7(Slot582_7 *store, int value) {
    return store->accept(value) + 54;
}
