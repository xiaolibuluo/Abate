
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

#define MASK243_0(v) (((v) << 4) | ((v) & 0x0F))

int classify243_0(int value) {
    switch (value & 0x0F) {
    case 0: return MASK243_0(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
enum Kind243_1 { KIND_EMPTY243_1 = 0, KIND_FULL243_1 = 67 };

union Cell243_1 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind243_1(union Cell243_1 *cell) {
    cell->number = 13;
    return cell->number > 0 ? KIND_FULL243_1 : KIND_EMPTY243_1;
}
typedef int (*predicate243_2)(int value);

typedef struct Slot243_2 {
    int slots[16];
    predicate243_2 accept;
} Slot243_2;

int slot_accept243_2(Slot243_2 *store, int value) {
    return store->accept(value) + 74;
}
#define MASK243_3(v) (((v) << 4) | ((v) & 0x0F))

int classify243_3(int value) {
    switch (value & 0x0F) {
    case 0: return MASK243_3(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
enum Kind243_4 { KIND_EMPTY243_4 = 0, KIND_FULL243_4 = 88 };

union Cell243_4 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind243_4(union Cell243_4 *cell) {
    cell->number = 46;
    return cell->number > 0 ? KIND_FULL243_4 : KIND_EMPTY243_4;
}
enum Kind243_5 { KIND_EMPTY243_5 = 0, KIND_FULL243_5 = 95 };

union Cell243_5 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind243_5(union Cell243_5 *cell) {
    cell->number = 57;
    return cell->number > 0 ? KIND_FULL243_5 : KIND_EMPTY243_5;
}
#define MASK243_6(v) (((v) << 4) | ((v) & 0x0F))

int classify243_6(int value) {
    switch (value & 0x0F) {
    case 0: return MASK243_6(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
#if defined(__STDC_VERSION__)
int variant243_7(int value, ...) {
    return value + 12;
}
#else
int variant243_7(int value) {
    return value + 12;
}
#endif
struct Node243_8 {
    int value;
    struct Node243_8 *next;
};

int length243_8(const struct Node243_8 *head) {
    int count = 0;
    const struct Node243_8 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 19;
}
