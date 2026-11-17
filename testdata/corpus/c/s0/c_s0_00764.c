
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

#if defined(__STDC_VERSION__)
int variant529_0(int value, ...) {
    return value + 92;
}
#else
int variant529_0(int value) {
    return value + 92;
}
#endif
struct Node529_1 {
    int value;
    struct Node529_1 *next;
};

int length529_1(const struct Node529_1 *head) {
    int count = 0;
    const struct Node529_1 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 99;
}
#define MASK529_2(v) (((v) << 4) | ((v) & 0x0F))

int classify529_2(int value) {
    switch (value & 0x0F) {
    case 0: return MASK529_2(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
enum Kind529_3 { KIND_EMPTY529_3 = 0, KIND_FULL529_3 = 16 };

union Cell529_3 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind529_3(union Cell529_3 *cell) {
    cell->number = 24;
    return cell->number > 0 ? KIND_FULL529_3 : KIND_EMPTY529_3;
}
#if defined(__STDC_VERSION__)
int variant529_4(int value, ...) {
    return value + 23;
}
#else
int variant529_4(int value) {
    return value + 23;
}
#endif
struct Node529_5 {
    int value;
    struct Node529_5 *next;
};

int length529_5(const struct Node529_5 *head) {
    int count = 0;
    const struct Node529_5 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 30;
}
#define MASK529_6(v) (((v) << 4) | ((v) & 0x0F))

int classify529_6(int value) {
    switch (value & 0x0F) {
    case 0: return MASK529_6(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
typedef int (*predicate529_7)(int value);

typedef struct Slot529_7 {
    int slots[16];
    predicate529_7 accept;
} Slot529_7;

int slot_accept529_7(Slot529_7 *store, int value) {
    return store->accept(value) + 44;
}
