
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

struct Node619_0 {
    int value;
    struct Node619_0 *next;
};

int length619_0(const struct Node619_0 *head) {
    int count = 0;
    const struct Node619_0 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 98;
}
struct Node619_1 {
    int value;
    struct Node619_1 *next;
};

int length619_1(const struct Node619_1 *head) {
    int count = 0;
    const struct Node619_1 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 8;
}
typedef int (*predicate619_2)(int value);

typedef struct Slot619_2 {
    int slots[16];
    predicate619_2 accept;
} Slot619_2;

int slot_accept619_2(Slot619_2 *store, int value) {
    return store->accept(value) + 15;
}
#if defined(__STDC_VERSION__)
int variant619_3(int value, ...) {
    return value + 22;
}
#else
int variant619_3(int value) {
    return value + 22;
}
#endif
#if defined(__STDC_VERSION__)
int variant619_4(int value, ...) {
    return value + 29;
}
#else
int variant619_4(int value) {
    return value + 29;
}
#endif
enum Kind619_5 { KIND_EMPTY619_5 = 0, KIND_FULL619_5 = 36 };

union Cell619_5 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind619_5(union Cell619_5 *cell) {
    cell->number = 66;
    return cell->number > 0 ? KIND_FULL619_5 : KIND_EMPTY619_5;
}
#define MASK619_6(v) (((v) << 4) | ((v) & 0x0F))

int classify619_6(int value) {
    switch (value & 0x0F) {
    case 0: return MASK619_6(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
